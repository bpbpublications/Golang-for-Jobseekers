package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewMemoryStore() MemoryStore {
	return MemoryStore{items: make(map[string]string)}
}

type MemoryStore struct {
	items map[string]string
}

func (m *MemoryStore) Add(shortendURL, longURL string) error {
	if m.items[shortendURL] != "" {
		return fmt.Errorf("value already exists here")
	}
	m.items[shortendURL] = longURL
	log.Println(m.items)
	return nil
}

func (m *MemoryStore) Remove(shortenedURL string) error {
	delete(m.items, shortenedURL)
	return nil
}

func (m *MemoryStore) Get(shortendURL string) (string, error) {
	longURL := m.items[shortendURL]
	if longURL == "" {
		return "", fmt.Errorf("no mapped url available here")
	}
	return longURL, nil
}

type Store interface {
	Add(shortenedURL, longURL string) error
	Remove(shortenedURL string) error
	Get(shortendURL string) (string, error)
}

type AddPath struct {
	domain string
	store  Store
}

func (a *AddPath) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type addPathRequest struct {
		URL string `json:"url"`
	}

	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unexpected error"))
		return
	}
	var parsed addPathRequest
	json.Unmarshal(raw, &parsed)

	h := sha1.New()
	h.Write([]byte(parsed.URL))
	sum := h.Sum(nil)
	hash := hex.EncodeToString(sum)[:10]

	err = a.store.Add(hash, parsed.URL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("unexpected error :: %v", err)))
		return
	}

	type addPathResponse struct {
		ShortenedURL string `json:"shortened_url"`
		LongURL      string `json:"long_url"`
	}

	pathResp := addPathResponse{ShortenedURL: fmt.Sprintf("%v/%v", a.domain, hash), LongURL: parsed.URL}
	rawResp, err := json.Marshal(pathResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unexpected error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(rawResp)
}

type DeletePath struct {
	store Store
}

func (p *DeletePath) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty hash"))
		return
	}

	err := p.store.Remove(hash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("unexpected error :: %v", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted"))
}

type RedirectPath struct {
	store Store
}

func (p *RedirectPath) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty hash"))
		return
	}

	longURL, err := p.store.Get(hash)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}

	http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
}

type HandleViaStruct struct{}

func (*HandleViaStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	defer log.Print("End hello world request")
	fmt.Fprintf(w, "Hello World via Struct")
}

func main() {
	log.Print("Hello world sample started.")
	r := mux.NewRouter()
	redirectPath := "http://localhost:8080/r"
	mem := NewMemoryStore()
	r.Handle("/", &HandleViaStruct{}).Methods("GET")
	r.Handle("/add", &AddPath{domain: redirectPath, store: &mem}).Methods("POST")
	r.Handle("/r/{hash}", &DeletePath{store: &mem}).Methods("DELETE")
	r.Handle("/r/{hash}", &RedirectPath{store: &mem}).Methods("GET")
	http.ListenAndServe(":8080", r)
}
