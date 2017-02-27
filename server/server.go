package server

import (
	"chatr/helper"
	"golang.org/x/net/websocket"
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseFiles("index.html"))
}

type socket struct {
	io.ReadWriter
	done chan bool
}

//RootHandler it's the first page in the webapp
func RootHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "localhost:8080")
}

func (s socket) Close() error {
	s.done <- true
	return nil
}

func SocketHandler(ws *websocket.Conn) {
	s := socket{ws, make(chan bool)}
	go helper.Match(s)
	<-s.done
}
