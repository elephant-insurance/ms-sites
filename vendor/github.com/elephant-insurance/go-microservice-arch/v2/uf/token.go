package uf

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"text/template"
	"time"
)

// NewAzureServiceBusTokenFactory creates a token factory for Azure Service Bus messages
func NewAzureServiceBusTokenFactory(secret, keyName, resourceName, queueOrTopic string) (TokenFactory, error) {
	resourceURI := fmt.Sprintf(asbURIPattern, resourceName, queueOrTopic)
	escapedURI := template.URLQueryEscaper(resourceURI)
	return NewTokenFactory(secret, keyName, escapedURI)
}

// NewTokenFactory creates a token factory for the submitted params
func NewTokenFactory(secret, keyName, resourceURI string) (TokenFactory, error) {
	/*key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return nil, err
	}*/

	return &hasher{
		key:         []byte(secret),
		keyName:     keyName,
		resourceURI: resourceURI,
	}, nil
}

type TokenFactory interface {
	GetSum(clearText *string) ([]byte, error)
	GetSASToken(expiry *time.Time) (string, error)
}

type hasher struct {
	key         []byte
	keyName     string
	resourceURI string
}

func (h *hasher) GetSum(clearText *string) ([]byte, error) {
	if h == nil || len(h.key) == 0 {
		return nil, errors.New(`hasher not initialized`)
	}

	if clearText == nil {
		return nil, errors.New(`nil cleartext`)
	}

	thehash := hmac.New(sha256.New, h.key)
	_, _ = thehash.Write([]byte(*clearText))

	return thehash.Sum(nil), nil
}

// GetSASToken returns a formatted SAS token using the factory's set secret, uri, and key name.
// Expiry defaults to one week from now (UTC) if not sent.
func (h *hasher) GetSASToken(expiry *time.Time) (string, error) {
	if h == nil || len(h.key) == 0 || h.keyName == `` || h.resourceURI == `` {
		return ``, errors.New(`TokenFactory not properly initialized`)
	}

	// if expiry is nil, set it to a week from now
	if expiry == nil {
		wfn := time.Now().UTC().AddDate(0, 0, 7)
		expiry = &wfn
	}

	expSeconds := expiry.Unix()
	clearText := fmt.Sprintf("%v\n%v", h.resourceURI, expSeconds)
	hashBytes, serr := h.GetSum(&clearText)
	if serr != nil {
		return ``, serr
	}

	hashString := template.URLQueryEscaper(base64.StdEncoding.EncodeToString(hashBytes))

	return fmt.Sprintf(`SharedAccessSignature sr=%v&sig=%v&se=%v&skn=%v`, h.resourceURI, hashString, expSeconds, h.keyName), nil
}

const (
	asbURIPattern string = `https://%v.servicebus.windows.net/%v/messages`
)
