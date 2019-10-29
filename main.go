package main
import (
	"github.com/gorilla/mux"
	"time"
	"net/http"
	"io"
	"fmt"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
	fmt.Printf("Have connection\n")
	time.Sleep(500 * time.Millisecond)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", r)
}