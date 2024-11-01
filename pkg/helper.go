package pkg

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func UuidString() string {
	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	return uuid
}

func IntRand(length int) float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	var letters = []rune("1234567890")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	f, _ := strconv.ParseFloat(string(b), 64)
	return f
}
