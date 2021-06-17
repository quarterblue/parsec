package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func blockChainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Blockchain Page")
	fmt.Println("Endpoint Hit: Blockchain Page")
}

func blockPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Block Page")
	fmt.Println("Endpoint Hit: BlockPage")
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Parsec API")
	fmt.Println("Endpoint Hit: mainpage")
}

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", mainPage)
	myRouter.HandleFunc("/blockchain", blockChainPage)
	myRouter.HandleFunc("/block", blockPage)
	log.Fatal(http.ListenAndServe(":9200", myRouter))
}
