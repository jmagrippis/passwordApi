package main

import (
	"github.com/jmagrippis/password"
	"github.com/julienschmidt/httprouter"

	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Message is a simple struct used for json responses.
// It has a body and it might have an error message and the recommended endpoints the user should try next.
type Message struct {
	Body      string   `json:"message"`
	Endpoints []string `json:"endpoints,omitempty"`
	Error     string   `json:"error,omitempty"`
}

// Welcome controller introduces clients to the API.
func Welcome(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response, err := json.Marshal(&Message{
		Body:      "Welcome to the Memorable Password Generator!",
		Endpoints: []string{"generator/:amount", "generator/:amount/safe"},
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Generate is the main controller of the API. It generates an :amount of passwords
// formatted according to the optional get parameters of the request.
func Generate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	amount, err := strconv.Atoi(params.ByName("amount"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if amount > 255 || amount < 1 {
		http.Error(w, "Please request a number of passwords between 1 and 255, inclusive", http.StatusBadRequest)
		return
	}

	generator := getGenerator(r.URL.Query())

	passwords := make([]string, amount)

	for i := 0; i < amount; i++ {
		passwords[i] = generator.Generate()
	}

	response, err := json.Marshal(passwords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

// GenerateSafe is a helper endpoint, generating a password that will be accepted on most web forms.
func GenerateSafe(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	amount, err := strconv.Atoi(params.ByName("amount"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if amount > 255 || amount < 1 {
		http.Error(w, "Please request a number of passwords between 1 and 255, inclusive", http.StatusBadRequest)
		return
	}

	properties := url.Values{}

	delimiters := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	properties.Add("delimiter", delimiters[rand.Intn(len(delimiters))])
	suffixes := []string{"!", "#", "$", "^", "&", "(", ")", "="}
	properties.Add("suffix", suffixes[rand.Intn(len(suffixes))])

	generator := getGenerator(properties)
	generator.SetTitleCase(true)

	passwords := make([]string, amount)

	for i := 0; i < amount; i++ {
		passwords[i] = generator.Generate()
	}

	response, err := json.Marshal(passwords)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

// getGenerator defines the dictionary and returns a new generator that will be seeded according to the current time.
// Also sets any options defined from the query parameters.
func getGenerator(parameters url.Values) *password.Generator {
	dictionary := &password.Dictionary{
		Adverbs:    []string{"cuddling", "slapping", "shouting", "jumping"},
		Subjects:   []string{"mermaids", "unicorns", "lions", "piranhas"},
		Verbs:      []string{"love", "fancy", "eat", "bring", "fear", "aggravate"},
		Adjectives: []string{"beautiful", "homely", "magical", "posh", "excellent"},
		Objects:    []string{"teddy-bears", "diamonds", "buckets", "boxes"},
	}

	generator := password.NewGenerator(dictionary, time.Now().UnixNano())

	delimiter := parameters.Get("delimiter")
	if delimiter != "" {
		generator.SetDelimiter(delimiter)
	}

	prefix := parameters.Get("prefix")
	if prefix != "" {
		generator.SetPrefix(prefix)
	}

	suffix := parameters.Get("suffix")
	if suffix != "" {
		generator.SetSuffix(suffix)
	}
	return generator
}
