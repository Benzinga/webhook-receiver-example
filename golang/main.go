package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/",  WebHookHandler)
	log.Fatal()
	// insert server configuration and run here
	log.Printf("insert server start here")
}

// DATA being sent from Benzinga's webhook
type HookContext struct {
	Id        string
	Payload   []byte
}
// Parse request Benzinga sent
func ParseHook( req *http.Request) (*HookContext, error) {
	hc := HookContext{}


	if hc.Id = req.Header.Get("x-github-delivery"); len(hc.Id) == 0 {
		return nil, errors.New("No event Id!")
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	hc.Payload = body

	return &hc, nil
}

func WebHookHandler(w http.ResponseWriter, r *http.Request) {

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
		log.Printf("JSON unmarshal error")
	}
	// parse `hc.Payload` or do additional processing here
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "{}")
	return
}

// a sample of a news response

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