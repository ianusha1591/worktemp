package main

import (
  "html/template"
  "net/http"
)

var testTemplate *template.Template

type Flower struct {
  Rose bool
//Lotus bool
}

type ViewData struct {
  *Flower
}

func main() {
  var err error
  testTemplate, err = template.ParseFiles("flower.tmpl")
  if err != nil {
    panic(err)
  }
http.HandleFunc("/flower", handler)
  http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")

  vd := ViewData{&Flower{true}}
  err := testTemplate.Execute(w, vd)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
