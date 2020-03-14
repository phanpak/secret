package secret

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/phanpak/secret/encrypt"
)

var filepath = "secrets.json"

type Vault struct {
	EncKey   string
	secrets  map[string]string
	filepath string
	mutex    sync.Mutex
}

func File(key string) *Vault {
	return &Vault{
		EncKey:   key,
		secrets:  make(map[string]string),
		filepath: filepath,
	}
}

func (v *Vault) load() error {
	file, err := os.Open(v.filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	var encryptedSecrets map[string]string
	err = json.NewDecoder(file).Decode(&encryptedSecrets)
	if err != nil {
		return nil
	}
	secrets := make(map[string]string, len(encryptedSecrets))
	for key, value := range encryptedSecrets {
		decValue, err := encrypt.Decrypt(v.EncKey, value)
		if err != nil {
			return err
		}
		secrets[key] = decValue
	}
	v.secrets = secrets
	return nil
}

func (v *Vault) save() error {
	encryptedSecrets := make(map[string]string, len(v.secrets))

	for key, value := range v.secrets {
		encValue, err := encrypt.Encrypt(v.EncKey, value)
		encryptedSecrets[key] = encValue
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	w := json.NewEncoder(file)
	err = w.Encode(&encryptedSecrets)
	if err != nil {
		return err
	}
	return nil
}

func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return err
	}
	_, ok := v.secrets[key]
	if ok {
		return fmt.Errorf("the key \"%s\" is already in the vault", key)

	}
	encryptedValue, err := encrypt.Encrypt(v.EncKey, value)
	if err != nil {
		return err
	}

	v.secrets[key] = encryptedValue
	err = v.save()
	if err != nil {
		return err
	}
	return nil
}

func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return "", err
	}
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
