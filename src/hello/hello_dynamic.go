package hello

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// HelloDynamic is an HTTP Cloud Function with a request parameter.
func HelloDynamic(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello, Dynamic!")
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, Dynamic!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}
