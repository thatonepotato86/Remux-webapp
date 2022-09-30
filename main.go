package main

import (
	"os"

	"github.com/remux-go/remux"
)

type Page struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

var data, _ = os.ReadFile("./static/files/text.txt")
var pagedata = Page{"Water Conservation 🌊", string(data)}

func main() {
	var router = remux.Remux{Port: "4000"}
	router.Get("/", index)
	router.FileServer("/static", "./static")
	router.Serve()
}

func index(e remux.Engine) {
	e.File("./templates/index.html", pagedata)
}
