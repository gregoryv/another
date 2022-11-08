package another

import (
	"encoding/json"
	"net/http"
)

func init() {
	http.HandleFunc("/another", GetGopher)
}

// @Success 200 {object} another.GetGopher.response
// @Router /another [get]
func GetGopher(w http.ResponseWriter, r *http.Request) {
	type gopher struct {
		Color string `json:"color"`
	}
	m := gopher{
		Color: "pink",
	}
	json.NewEncoder(w).Encode(m)
}
