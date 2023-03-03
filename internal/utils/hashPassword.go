package utils

import (
	"crypto/sha1"
	"fmt"

	"github.com/Sakagam1/DBMS_TASK/internal/config"
)

func GeneratePasswordHash(password string) string {
	conf := config.GetConfig()
	salt := conf.Salt

	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
