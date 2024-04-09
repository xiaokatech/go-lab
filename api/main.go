package api

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("test")
	srv := NewServer()
	http.ListenAndServe(":8080", srv)
}
