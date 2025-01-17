package app

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gilcrest/go-api-basic/domain/errs"
	"github.com/gilcrest/go-api-basic/domain/secure"
)

// APIKeyStringGenerator creates a random, 128 API key string
type APIKeyStringGenerator interface {
	RandomString(n int) (string, error)
}

// APIKey is an API key for interacting with the system
type APIKey struct {
	// key: the unencrypted API key string
	key string
	// ciphertext: the encrypted API key as []byte
	ciphertext []byte
	// deactivationDate: the date the API key is no longer usable
	deactivationDate time.Time
}

// NewAPIKey initializes an APIKey. It generates both a 128-bit (16 byte)
// random string as an API key and its corresponding ciphertext bytes
func NewAPIKey(g APIKeyStringGenerator, ek *[32]byte) (APIKey, error) {
	k, err := g.RandomString(18)
	if err != nil {
		return APIKey{}, err
	}

	ct, err := secure.Encrypt([]byte(k), ek)
	if err != nil {
		return APIKey{}, err
	}

	return APIKey{key: k, ciphertext: ct}, nil
}

// NewAPIKeyFromCipher initializes an APIKey
func NewAPIKeyFromCipher(ciphertext string, ek *[32]byte) (APIKey, error) {
	eak, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return APIKey{}, errs.E(errs.Internal, err)
	}

	apiKey, err := secure.Decrypt(eak, ek)
	if err != nil {
		return APIKey{}, err
	}

	return APIKey{key: string(apiKey), ciphertext: eak}, nil
}

// Key returns the key for the API key
func (a APIKey) Key() string {
	return a.key
}

// Ciphertext returns the hex encoded text of the encrypted cipher bytes for the API key
func (a APIKey) Ciphertext() string {
	return hex.EncodeToString(a.ciphertext)
}

// DeactivationDate returns the Deactivation Date for the API key
func (a APIKey) DeactivationDate() time.Time {
	return a.deactivationDate
}

// SetDeactivationDate sets the deactivation date value to AppAPIkey
// TODO - try SetDeactivationDate as a candidate for generics with 1.18
func (a *APIKey) SetDeactivationDate(t time.Time) {
	a.deactivationDate = t
}

// SetStringAsDeactivationDate sets the deactivation date value to
// AppAPIkey given a string in RFC3339 format
func (a *APIKey) SetStringAsDeactivationDate(s string) error {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return errs.E(errs.Validation, err)
	}
	a.deactivationDate = t

	return nil
}

// isValid validates the API Key
func (a APIKey) isValid(realm string) error {
	now := time.Now()
	if a.deactivationDate.Before(now) {
		return errs.E(errs.Unauthenticated, errs.Realm(realm), fmt.Sprintf("Key Deactivation Date %s is before current time %s", a.deactivationDate.String(), now.String()))
	}
	return nil
}
