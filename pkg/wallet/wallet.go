package wallet

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/pbkdf2"
	"lukechampine.com/blake3"
)

const (
	SecretKeyLength           = 32
	PBKDF2NbRound             = 10000
	FileModeUserReadWriteOnly = 0o600
	Base58Version             = 0x00
	SaltSize                  = 12
	NonceSize                 = 12
)

// KeyPair structure contains all the information necessary to save a key pair securely.
type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte
	Salt       []byte
	Nonce      []byte
}

// Wallet structure allows to link a nickname, an address and a version to one or more key pairs.
type Wallet struct {
	Version  uint8
	Nickname string
	Address  string
	KeyPair  KeyPair
}

// aead returns a authenticated encryption with associated data.
func aead(password []byte, salt []byte) (cipher.AEAD, error) {
	secretKey := pbkdf2.Key([]byte(password), salt, PBKDF2NbRound, SecretKeyLength, sha256.New)

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, fmt.Errorf("intializing block ciphering: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("intializing the AES block cipher wrapped in a Gallois Counter Mode ciphering: %w", err)
	}

	return aesGCM, nil
}

// Protect encrypts the private key using the given password.
// The encryption algorithm used to protect the private key is AES-GCM and
// the secret key is derived from the given password using the PBKDF2 algorithm.
func (w *Wallet) Protect(password string) error {

	aead, err := aead([]byte(password), w.KeyPair.Salt[:])
	if err != nil {
		return fmt.Errorf("while protecting wallet: %w", err)
	}

	w.KeyPair.PrivateKey = aead.Seal(
		nil,
		w.KeyPair.Nonce[:],
		w.KeyPair.PrivateKey,
		nil)

	return nil
}

// Unprotect decrypts the private key using the given password.
// The encryption algorithm used to unprotect the private key is AES-GCM and
// the secret key is derived from the given password using the PBKDF2 algorithm.
func (w *Wallet) Unprotect(password string) error {
	aead, err := aead([]byte(password), w.KeyPair.Salt[:])
	if err != nil {
		return fmt.Errorf("while unprotecting wallet: %w", err)
	}

	pk, err := aead.Open(nil, w.KeyPair.Nonce[:], w.KeyPair.PrivateKey, nil)
	if err != nil {
		return fmt.Errorf("opening the private key seal: %w", err)
	}

	w.KeyPair.PrivateKey = pk

	return nil
}

// Persist stores the wallet on the file system.
// Note: the wallet is stored in JSON format and in the working directory.
func (w *Wallet) Persist() error {
	jsonified, err := json.Marshal(w)
	if err != nil {
		return fmt.Errorf("marshalling wallet: %w", err)
	}

	err = os.WriteFile(Filename(w.Nickname), jsonified, FileModeUserReadWriteOnly)
	if err != nil {
		return fmt.Errorf("writing wallet to '%s: %w", Filename(w.Nickname), err)
	}

	return nil
}

// LoadAll loads all the wallets in the working directory.
// Note: a wallet must have: `wallet_` prefix and a `.json` extension.
func LoadAll() ([]Wallet, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("returning working directory: %w", err)
	}

	files, err := os.ReadDir(workingDir)
	if err != nil {
		return nil, fmt.Errorf("reading working directory '%s': %w", workingDir, err)
	}

	wallets := []Wallet{}
	for _, f := range files {
		fileName := f.Name()

		if strings.HasPrefix(fileName, "wallet_") && strings.HasSuffix(fileName, ".json") {
			content, err := os.ReadFile(fileName)
			if err != nil {
				return nil, fmt.Errorf("reading file '%s': %w", fileName, err)
			}

			wallet := Wallet{} //nolint:exhaustruct

			err = json.Unmarshal(content, &wallet)
			if err != nil {
				return nil, fmt.Errorf("json unmarshaling file '%s': %w", fileName, err)
			}

			wallets = append(wallets, wallet)
		}
	}

	return wallets, nil
}

// Load loads the wallet that match the given name in the working directory
// Note: `wallet_` prefix and a `.json` extension are automatically added.
func Load(nickname string) (*Wallet, error) {
	walletName := Filename(nickname)
	content, err := os.ReadFile(walletName)
	if err != nil {
		return nil, fmt.Errorf("reading file '%s': %w", walletName, err)
	}

	wallet := Wallet{} //nolint:exhaustruct

	err = json.Unmarshal(content, &wallet)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling file '%s': %w", walletName, err)
	}

	return &wallet, nil
}

// Generate instantiates a new wallet, protects its private key and persists it.
// Everything is dynamically generated except from the nickname.
func Generate(nickname string, password string) (*Wallet, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, fmt.Errorf("generating ed25519 keypair: %w", err)
	}

	addr := blake3.Sum256(publicKey)

	salt := make([]byte, SaltSize)

	_, err = rand.Read(salt[:])
	if err != nil {
		return nil, fmt.Errorf("generating random salt: %w", err)
	}

	nonce := make([]byte, NonceSize)

	_, err = rand.Read(nonce[:])
	if err != nil {
		return nil, fmt.Errorf("generating random nonce: %w", err)
	}

	wallet, err := New(nickname, addr, privateKey, publicKey, salt, nonce)
	if err != nil {
		return nil, fmt.Errorf("instantiating a new wallet: %w", err)
	}

	err = wallet.Protect(password)
	if err != nil {
		return nil, fmt.Errorf("protecting the new wallet: %w", err)
	}

	err = wallet.Persist()
	if err != nil {
		return nil, fmt.Errorf("persisting the new wallet: %w", err)
	}

	return wallet, nil

}

// New instantiates a new wallet.
func New(nickname string, addr [32]byte, privateKey []byte, publicKey []byte, salt []byte, nonce []byte) (*Wallet, error) {
	wallet := Wallet{
		Version:  0,
		Nickname: nickname,
		Address:  "A" + base58.CheckEncode(addr[:], Base58Version),
		KeyPair: KeyPair{
			PrivateKey: privateKey,
			PublicKey:  publicKey,
			Salt:       salt,
			Nonce:      nonce,
		},
	}

	return &wallet, nil
}

// Delete removes wallet from file system
func Delete(nickname string) (err error) {
	err = os.Remove(Filename(nickname))
	if err != nil {
		return fmt.Errorf("deleting wallet '%s': %w", Filename(nickname), err)
	}

	return nil
}

// Filename returns the wallet Filename based on the given nickname.
func Filename(nickname string) string {
	return fmt.Sprintf("wallet_%s.json", nickname)
}
