package rest

import(
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func Start(){
	go func() {
		log.Fatal(http.ListenAndServe(":8000", getRouter()))
	}()
}

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/alert",AlertHandler)
	return r
}

