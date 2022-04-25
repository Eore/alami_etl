package uid

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateUID() string {
	b := make([]byte, 32)
	rand.Read(b)

	return base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
}
