package main

import (
	"html/template"
	"main/pkg"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		return
	}

	files := []string{
		"./ui/html/index.html",
		"./ui/html/base.html",
		"./ui/html/news-card.html",
	}

	db := pkg.OpenDB()

	//TODO переписать на горутины
	LatestNews := []*pkg.News{}
	db.Preload("User").Order("created_at desc").Limit(5).Find(&LatestNews)

	TopStory := pkg.News{}
	db.Order("RAND()").First(&TopStory)

	data := map[string]any{
		"LatestNews": LatestNews,
		"TopStory":   TopStory,
	}

	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.Execute(w, data)
}

func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	files := []string{
		"./ui/html/404.html",
		"./ui/html/base.html",
	}

	w.WriteHeader(http.StatusNotFound)

	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.Execute(w, nil)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/contact.html",
		"./ui/html/base.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.Execute(w, nil)
}

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/news.html",
		"./ui/html/base.html",
		"./ui/html/news-card.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.Execute(w, nil)
}

func NewsDetailHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/news-detail.html",
		"./ui/html/base.html",
	}

	// id := r.PathValue("id")

	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.Execute(w, nil)
}
