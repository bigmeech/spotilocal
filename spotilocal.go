package spotilocal

import (
	rand "math/rand"
	http "net/http"
	logger "github.com/op/go-logging"
	time "time"
	json "encoding/json"
	ioutil "io/ioutil"
)

const ORIGIN_URL string = "https://open.spotify.com"
const LOCAL_URL string = "https://open.spotify.com"
const DEFAULT_PORT int = 4370
const MIN_PORT int = DEFAULT_PORT
const MAX_PORT int = 4380;
const RANDOM_STRING_LEN int = 10
const LOCAL_URL_PREFIX string = ".spotilocal.com"

const TOKEN_PATH string = "/token"
const CSRF_TOKEN_PATH string = "/simplecsrf/token.json"

var log = logger.MustGetLogger("Spotilocal")

type Spotilocal struct {
	Host string
	Port string
	T Token
	CToken CSRFToken
}

//generates a random string as subdomain
func getSubDomain(str_len int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz123456789"
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, str_len);
	for i := 0; i < str_len; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result) + LOCAL_URL_PREFIX;
}

// Detect connectable ports by trying port ranges 4370 to 4380
func Connect () Spotilocal {
	var token Token
	var csrf_token CSRFToken

	randomString := getSubDomain(RANDOM_STRING_LEN)
	log.Debug("Connecting to " + randomString)

	//TYPES ASSERTIONS YO!!!!!!! FROM SUMTING TO ANYTING BACK TO SOMETHING.
	decodedToken := getJSON(ORIGIN_URL + TOKEN_PATH, token).(Token)
	decodedCSRFToken := getJSON(ORIGIN_URL + CSRF_TOKEN_PATH, csrf_token).(CSRFToken);

	log.Debug(decodedToken)

	return Spotilocal{ T: decodedToken, CToken: decodedCSRFToken }
}

// Gets token from remote
func getJSON (url string, target interface{}) interface{} {
	response, resp_err := http.Get(url)

	if resp_err != nil {
		panic(resp_err.Error())
	}
	defer response.Body.Close()

	raw_body, io_error := ioutil.ReadAll(response.Body)
	if io_error != nil {
		panic(io_error)
	}
	json_data := []byte(raw_body)

	decode_error := json.Unmarshal(json_data, &target)
	if decode_error != nil {
		panic(decode_error.Error())
	}
	return target
}


type CSRFToken struct {
	Token string `json: "token"`
}

type Token struct {
	T string `json: "t"`
}

// start local server
func (s *Spotilocal) Start(address string){}
