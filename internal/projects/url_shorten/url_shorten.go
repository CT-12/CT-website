package urlshorten

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

type UrlShortenRequest struct {
	LongUrl string `json:"long_url" binding:"required,url"`
}

func GenerateShortURL(long_url string) string {
	input := fmt.Sprintf(
		"%s-%d-%d", long_url, time.Now().UnixNano(), rand.Int(), 
	)
	hash := sha256.Sum256([]byte(input))

	return fmt.Sprintf("%x", hash)[:8] // 取前8位作為短網址
}