package spotilocal

import (
	rand "math/rand"
	http "net/http"
	logger "github.com/op/go-logging"
	time "time"
	"encoding/json"
)

const ORIGIN_URL string = "https://open.spotify.com"
const DEFAULT_PORT int = 4370
const MIN_PORT int = DEFAULT_PORT
const MAX_PORT int = 4380;
const RANDOM_STRING_LEN int = 10
const LOCAL_URL_PREFIX string = ".spotilocal.com"

var log = logger.MustGetLogger("Spotilocal")

type Spotilocal struct {
	Host string
	Port string
}

//generates a random string as subdomain
func getSubDomaiin() string {
	const chars = "abcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, RANDOM_STRING_LEN);
	for i := 0; i < RANDOM_STRING_LEN; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result) + LOCAL_URL_PREFIX;
}

// Detect connectable ports by trying port ranges 4370 to 4380
func Connect () {
	randomString := getSubDomaiin()
	log.Debug("Connecting to " + randomString)
	decodedToken := GetToken()
	log.Debug(decodedToken.T)
}

// Gets token from remote
func GetToken () Token {
	var token Token
	response, resp_err := http.Get(ORIGIN_URL+ "/token")

	if resp_err != nil {
		panic(resp_err.Error())
	}
	defer response.Body.Close()

	decode_error := json.NewDecoder(response.Body).Decode(&token)
	if decode_error != nil {
		panic(decode_error.Error())
	}
	return token
}


type Token struct {
	T string
}

// start local server
func (s *Spotilocal) Start(address string){}
