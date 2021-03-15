package main

import (
	"net/http"
	"log"
	"fmt"
	"time"
	"encoding/json"
	"strconv"
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
	"github.com/antage/eventsource"
)

func postMessageHandler(w http.ResponseWriter, r *http.Request){
	msg := r.FormValue("msg")
	name := r.FormValue("name")
	log.Println("postMessageHandler ",msg,name)
	sendMessage(name,msg)
}

func addUserHandler(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("name")
	sendMessage("", fmt.Sprintf("add user : %s",username))
}

type Message struct{
	Name string `json : "name"`
	Msg string `json : "msg"`
}
var msgCh chan Message
func sendMessage(name, msg string){
	//send message to every client
	msgCh <- Message{name,msg}
}

func processMsgCh(es eventsource.EventSource){
	for msg := range msgCh{
		data , _ := json.Marshal(msg)
		es.SendEventMessage(string(data),"",strconv.Itoa(time.Now().Nanosecond()))
	}
}

func leftUserHandler(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	sendMessage("", fmt.Sprintf("left user : %s",username))
}
func main(){
	msgCh = make(chan Message)
	es := eventsource.New(nil,nil)
	defer es.Close()
	go processMsgCh(es)

	mux := pat.New()
	mux.Post("/messages", postMessageHandler)
	mux.Handle("/stream",es)
	mux.Post("/users",addUserHandler)
	mux.Delete("/users", leftUserHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
