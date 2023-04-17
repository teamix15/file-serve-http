package templates

import (
	"html/template"
	"log"
	"net/http"
)

func HandleTemplate() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Title string
			Host  string
		}{
			Title: "Мой HTTP-сервер на Go",
			Host:  r.Host,
		}

		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		err := tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
		}
	})
}
