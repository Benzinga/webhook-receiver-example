package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// insert URL you have provided to Benzinga
	http.HandleFunc("/",  WebHookHandler)
	//optional: set port to other number
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("server start")
}

// Data being sent from Benzinga's webhook as a request: will always be json format
type HookContext struct {
	Id        string
	Payload   []byte
}
// Parse request Benzinga sent
func ParseHook( req *http.Request) (*HookContext, error) {
	hc := HookContext{}

// Double check if delivery header is set
	if hc.Id = req.Header.Get("X-BZ-Delivery"); len(hc.Id) == 0 {
		return nil, errors.New("No event Id!")
	}
// read body into reader
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
// establish payload
	hc.Payload = body

	return &hc, nil
}

func WebHookHandler(w http.ResponseWriter, r *http.Request) {
// parse request
	hc, err := ParseHook(r)

	w.Header().Set("Content-type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed processing hook! ('%s')", err)
		io.WriteString(w, "{}")
		return
	}

	log.Printf("Received %s", hc.Id)
	var res Body
	err = json.Unmarshal(hc.Payload, &res)
	if err != nil {
		log.Printf("JSON unmarshal error:", err)
	}
	// parse `hc.Payload` or do additional processing here

	//send 200 to Benzinga to let them know everything is okay!
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "{}")
	return
}

// a sample of a news response / what we unmarshal into for this sample

type Security struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	Primary  bool   `json:"primary"`
}

type Content struct {
	ID         int64  `json:"id"`
	RevisionID int    `json:"revision_id"`
	Type       string `json:"type"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Title      string     `json:"title"`
	Body       string     `json:"body"`
	Authors    []string   `json:"authors"`
	Teaser     string     `json:"teaser"`
	URL        string     `json:"url"`
	Tags       []string   `json:"tags"`
	Securities []Security `json:"securities"`
	Channels   []string   `json:"channels"`
}

type Data struct {
	Action    string    `json:"action"`
	ID        int64     `json:"id"`
	Content   Content   `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type Body struct {
	ID         string `json:"id"`
	ApiVersion string `json:"api_version"`
	Kind       string `json:"kind"`
	Data       Data   `json:"data"`
}