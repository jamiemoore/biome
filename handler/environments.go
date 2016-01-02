package handler

import (
	"fmt"
	"net/http"
)

func GetEnvironments(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
