package user

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// HelloUser is an HTTP Cloud Function with a request parameter.
func HelloUser(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello, User!")
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, User!")
		return
	}

	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}
