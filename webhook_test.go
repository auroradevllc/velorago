package velora

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestVerifySignature(t *testing.T) {
	// Note: This implementation works, however it seems Velora's server side implementation does not
	body := `{
  "id": "evt_abc123",
  "type": "stream.online",
  "timestamp": "2026-01-18T15:30:00Z",
  "data": {
    "streamId": "stream_xyz",
    "userId": "user_123",
    "username": "coolstreamer",
    "title": "Playing Velora Games!",
    "startedAt": "2026-01-18T15:30:00Z"
  }
}`

	timestamp := "2026-01-18T15:30:00Z"

	secret := "testing123"

	expected := "sha256="

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(timestamp + "." + body))

	calc := h.Sum(nil)

	expected += hex.EncodeToString(calc)

	err := VerifySignature(timestamp+"."+body, expected, secret)

	if err != nil {
		t.Fatal(err)
	}
}
