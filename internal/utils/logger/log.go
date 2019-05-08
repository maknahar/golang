package logger

import (
	"math/rand"
	"time"

	"os"

	"github.com/maknahar/go-web-skeleton/internal/utils/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

func GetLogger() *log.Entry {

	l := log.WithFields(log.Fields{
		"trackId": RandomString(10),
	})

	// Only log the set level severity or above.
	l.Level = log.Level(config.LogLevel)

	return l
}

func RandomString(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
