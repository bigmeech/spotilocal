package spotilocal

import (
	rand "math/rand"
	http "net/http"
	logger "github.com/op/go-logging"
	time "time"
)

const ORIGIN_URL string = "https://open.spotiy.com"
const DEFAULT_PORT int = 4370
const MIN_PORT int = DEFAULT_PORT
const MAX_PORT int = 4380;
const RANDOM_STRING_LEN int = 10

var log = logger.MustGetLogger("Spotilocal")

type Spotilocal struct {
	Host string
	Port string
}

//generates a random string as
func getSubDomaiin() string {
	const chars = "abcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, RANDOM_STRING_LEN);
	for i := 0; i < RANDOM_STRING_LEN; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result);
}

// Detect connectable ports by trying port ranges 4370 to 4380
func Connect () {
	randomString := getSubDomaiin()
	log.Info("Connecting to " + randomString)
	http.Get(randomString)
}

// Gets token from remote
func (s *Spotilocal) GetToken (){}

// start local server
func (s *Spotilocal) Start(address string){}
