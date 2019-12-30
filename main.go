package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/data", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	dat, _ := ioutil.ReadFile("./static/data.json")
	fmt.Fprint(w, string(dat))
}
