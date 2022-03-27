package cmd

import (
	"log"
	"net/http"
)

type server struct {
	httpContext *http.Client
}

func StartServer() {

}
func handler(s *server) http.Handler {
	h := s.httpContext
	h.Get("/")
	return nil
}
func main() {
	s := &server{}
	log.Fatalln("Server start with port 8000", http.ListenAndServe(":8000", handler(s)))
}