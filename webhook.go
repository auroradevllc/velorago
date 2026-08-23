package velora

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"strings"
)

var (
	ErrInvalidHeader    = errors.New("invalid expected header, should match sha256=SIGNATURE")
	ErrInvalidHasher    = errors.New("unsupported hashing algorithm")
	ErrInvalidSignature = errors.New("invalid signature")
)

func VerifySignature(body, expected, secret string) error {
	splitIdx := strings.Index(expected, "=")

	if splitIdx == -1 {
		return ErrInvalidHeader
	}

	hasher := expected[0:splitIdx]

	expected = expected[splitIdx+1:]

	// An unknown algorithm yields a nil constructor, which would panic in hmac.New.
	newFn := NewHasher(hasher)

	if newFn == nil {
		return fmt.Errorf("%v: %s", ErrInvalidHasher, hasher)
	}

	provided, err := hex.DecodeString(expected)

	if err != nil {
		return err
	}

	mac := hmac.New(newFn, []byte(secret))
	mac.Write([]byte(body))

	if !hmac.Equal(provided, mac.Sum(nil)) {
		return ErrInvalidSignature
	}

	return nil
}

func NewHasher(hasher string) func() hash.Hash {
	switch hasher {
	case "sha1":
		return sha1.New
	case "sha256":
		return sha256.New
	case "sha384":
		return sha512.New384
	case "sha512":
		return sha512.New
	}

	return nil
}
