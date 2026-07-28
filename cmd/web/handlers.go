package main

import (
	"html/template"
	"log"
	"main/pkg"
	"net/http"
	"sync"
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
	const jobs = 2

	errCh := make(chan error, jobs)
	var wg sync.WaitGroup
	wg.Add(jobs)

	go func() {
		defer wg.Done()
		if err := db.Preload("User").Order("created_at desc").Limit(5).Find(&LatestNews).Error; err != nil {
			errCh <- err
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.Order("RAND()").First(&TopStory).Error; err != nil {
			errCh <- err
		}
	}()

	wg.Wait()
	close(errCh)

	if err, ok := <-errCh; ok {
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
	files := []string{
		"./ui/html/news-detail.html",
		"./ui/html/base.html",
	}

	// id := r.PathValue("id")

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
