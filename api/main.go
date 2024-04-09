package api

import (
	"fmt"
	"net/http"

	"github.com/xiaokatech/go-lab/api"
)

func main() {
	fmt.Println("test")
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
