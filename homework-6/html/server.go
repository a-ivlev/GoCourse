package server

import (
	"log"
	"net/http"
	"html/template"
)

type IndexTemplate struct {
	Title string
	Name string
	Count int
	Products []product
}

type product struct {
	Name string
}

// Формируем html страницу по шаблону
func hello(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	tmpl, err := template.ParseFiles("assets/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	indexTemplate := IndexTemplate{
		Title: "Products page",
		Name:	name,
		Products: []product{
			{
				Name: "Mill",
			},
			{
				Name: "Water",
			},
		},
	}

	indexTemplate.Count = len(indexTemplate.Products)

	err = tmpl.Execute(res, indexTemplate)
	if err != nil {
		log.Fatal(err)
	}
}
// Запускаем http сервер
func ServerGo()  {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}