package main

import (
	"airquality-webapp/manualapi"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func main() {
	srv := manualapi.NewServer()
	http.ListenAndServe(":8000", nil)

}

func othermain() {
	welcome := Welcome{"Roshan", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("template/mainview.html"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "mainview.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
