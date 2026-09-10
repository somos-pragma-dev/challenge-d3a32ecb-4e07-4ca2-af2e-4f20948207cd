package idempotency

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type IdempotencyKey string

func GenerateIdempotencyKey(ctx context.Context, requestBody []byte) IdempotencyKey {
	hash := sha256.New()
	hash.Write(requestBody)
	return IdempotencyKey(hex.EncodeToString(hash.Sum(nil)))
}