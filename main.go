package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/handlers"
)

func main() {
	router := mux.NewRouter()

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})

	router.Handle("/socket.io/", server)
	router.Handle("/", http.FileServer(http.Dir("./asset")))
	router.Handle("/assets", http.FileServer(http.Dir("./assets")))

	// provide default cors to the mux
	//handler := cors.Default().Handler(mux)

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(originsOk, headersOk, methodsOk, handlers.AllowCredentials())(router)))
	//log.Fatal(http.ListenAndServe(":5000", handlers.CORS()(router)))
	//log.Fatal(http.ListenAndServe(":5000", handlers.CORS(handlers.AllowedOrigins([]string{"http://127.0.0.1:3000"}))(router)))

}
