package main
 
import (
"time"
    "fmt"
    "log"
    "net/http"
"html/template"
)
type Fullname struct {
    FirstName string
}
type Application struct {
    Date          time.Time
    Description   string
    
}
type Interview struct {
    OnDate    time.Time
    Fullname   Fullname
    Apps []Application
}
func reminder(fn Fullname) string {
    return fmt.Sprintf("You should attend the interview on the given date only.")
}
func date(t time.Time) string {
    year, month, day := t.Date()
    return fmt.Sprintf("%d/%d/%d", day, month, year)
}

func inter(w http.ResponseWriter, r *http.Request) {
functionmap := template.FuncMap{
        "reminder": reminder,
        "date":date,
}
t :=template.Must(template.New("message.tmpl").Funcs(functionmap).ParseFiles("message.tmpl"))
       
        c :=Interview {
Fullname: Fullname{
            FirstName: "Anusha",
        },
 Apps: []Application{
             Application{
                Date: time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC),
                Description: "Devops engineer",
            },
    },

} 
t.Execute(w, c)
}
//}
func main() {

   http.HandleFunc("/me", inter) 
    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal(err)
    }
}
