package utils

import (
	rnd "math/rand"
	"time"
)

const (
	KIBANA_CONFIG_FILE = "./config/kibana.json"
	LOG_PATH           = "/logs/go.log"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateRandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rnd.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateRandomInt() int {
	r := rnd.New(rnd.NewSource(time.Now().UnixNano()))
	return r.Int()
}

func GenerateRandomExternalPlayerID() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rnd.Intn(len(letterRunes))]
	}
	return string(b)
}
