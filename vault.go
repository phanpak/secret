package secret

import (
	"fmt"

	"github.com/phanpak/secret/encrypt"
)

type Vault struct {
	EncKey  string
	secrets map[string]string
}

func File(key string) *Vault {
	return &Vault{
		EncKey:  key,
		secrets: make(map[string]string),
	}
}

func (v *Vault) Set(key, value string) error {
	_, ok := v.secrets[key]
	if ok {
		return fmt.Errorf("the key \"%s\" is already in the vault", key)

	}
	encryptedValue, err := encrypt.Encrypt(v.EncKey, value)
	if err != nil {
		return err
	}

	v.secrets[key] = encryptedValue
	return nil
}

func (v *Vault) Get(key string) (string, error) {
	cipherHex, ok := v.secrets[key]

	if !ok {
		return "", fmt.Errorf("There is no key \"%s\" in the vault", key)
	}

	// value, err := encrypt.Decrypt(v.EncKey, cipherHex)
	value, err := encrypt.Decrypt(v.EncKey, cipherHex)
	if err != nil {
		return "", nil
	}

	return value, nil
}
