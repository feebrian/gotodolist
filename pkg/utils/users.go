package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/go-layer-architecture/internal/domain"
)

type UserBundle struct {
	u domain.CreateUser
}

func HashPassword(p string) string {
	h := sha256.New()
	h.Write([]byte(p))
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}
