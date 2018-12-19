package main

import (
    "net/http"
    "text/template"
    "log"
)

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content Type", "text/html")
    path := r.URL.Path[1:]
    log.Println(path)
    tmpl, err := template.New("anyNameForTemplate").Parse(doc)
    if err == nil {
        tmpl.Execute(w, path)
    }
}

const doc = `
<!DOCTYPE html>
<html>
<body>
    <h1>HEY WORLD</h1>
    <p>Accessed URL is {{.}}</p>
</body>
</html>
`

func main() {
    http.HandleFunc("/", myHandlerFunc)
    http.ListenAndServe(":8080", nil)
}
