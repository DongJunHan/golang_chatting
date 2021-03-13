package main

import (
	"net/http"
	"log"
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

func postMessageHandler(w http.ResponseWriter, r *http.Request){
	msg := r.FormValue("msg")
	name := r.FormValue("name")
	log.Println("postMessageHandler ",msg,name)
	snedMessage(name,msg)
}

func sendMessage(name, msg string){
	//send message to every client

}

func main(){
	mux := pat.New()
	mux.Post("/messages", postMessageHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
