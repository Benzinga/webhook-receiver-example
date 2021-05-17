// Now available in package form at https://github.com/rjz/githubhook

package bzWebhook

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)


type HookContext struct {
Id        string
Payload   []byte
}

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

func Handler(w http.ResponseWriter, r *http.Request) {

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