package main
 
import (
    "fmt"
    "log"
    "net/http"
"html/template"
)
type NeoCoverage struct {
        Named string
       
    }
type Coverage struct {
          Pagetitle string
        NeeoCoverage []NeoCoverage
    }
type Employee struct {
        Name string
       
    }
type Person struct {
          Names string
        Employer []Employee
NeeoCoverage []NeoCoverage
    }

func hello(w http.ResponseWriter, r *http.Request) {
tmple :=template.Must(template.ParseFiles("mainlayout.tmpl","mainheader.tmpl","mainfooter.tmpl","mainnav1.tmpl","mainnav2.tmpl","mainarticle.tmpl","mainart.tmpl","mainart1.tmpl"))

       // name := r.FormValue("name")
       
        personname :=Person {
Names: "Here we display name and email of employees",
Employer: []Employee{{"joy"},{"god"},{"thomas"},
},
NeeoCoverage: []NeoCoverage{{"jasmine"},{"lilly"},{"rose"},
        },
} 
tmple.Execute(w, personname)
}
//}
func main() {

   http.HandleFunc("/mee", hello) 
    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal(err)
    }
}
