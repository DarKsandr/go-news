package main

import (
	"html/template"
	"log"
	"main/pkg"
	"net/http"

	"golang.org/x/sync/errgroup"
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

	var (
		LatestNews []*pkg.News
		TopStory   pkg.News
	)
	var g errgroup.Group

	g.Go(func() error {
		return db.Preload("User").Order("created_at desc").Limit(5).Find(&LatestNews).Error
	})

	g.Go(func() error {
		return db.Order("RAND()").First(&TopStory).Error
	})

	if err := g.Wait(); err != nil {
		ErrorHandler(w, r, err)
		return
	}

	data := map[string]any{
		"LatestNews": LatestNews,
		"TopStory":   TopStory,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		ErrorHandler(w, r, err)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		ErrorHandler(w, r, err)
		return
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/404.html",
		"./ui/html/base.html",
	}

	w.WriteHeader(http.StatusNotFound)

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		ErrorHandler(w, r, err)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		ErrorHandler(w, r, err)
		return
	}
}

func ErrorHandler(w http.ResponseWriter, _ *http.Request, err error) {
	log.Println(err)

	files := []string{
		"./ui/html/500.html",
		"./ui/html/base.html",
	}

	w.WriteHeader(http.StatusInternalServerError)

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/contact.html",
		"./ui/html/base.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		ErrorHandler(w, r, err)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		ErrorHandler(w, r, err)
		return
	}
}

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/news.html",
		"./ui/html/base.html",
		"./ui/html/news-card.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		ErrorHandler(w, r, err)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		ErrorHandler(w, r, err)
		return
	}
}

func NewsDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	db := pkg.OpenDB()

	var news pkg.News

	if err := db.Preload("User").First(&news, id).Error; err != nil {
		NotFoundHandler(w, r)
		return
	}

	var AlsoLike []pkg.News
	if err := db.Order("RAND()").Limit(2).Find(&AlsoLike).Error; err != nil {
		ErrorHandler(w, r, err)
		return
	}

	data := map[string]any{
		"News":     news,
		"Content":  template.HTML(news.Content),
		"AlsoLike": AlsoLike,
	}

	files := []string{
		"./ui/html/news-detail.html",
		"./ui/html/base.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		ErrorHandler(w, r, err)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		ErrorHandler(w, r, err)
		return
	}
}
