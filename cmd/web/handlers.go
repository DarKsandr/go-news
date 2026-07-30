package main

import (
	"html/template"
	"log"
	"main/pkg"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type Header struct {
	ActiveMenu string
}

type Footer struct {
	RecentPosts []*pkg.News
}

type Page struct {
	Data map[string]any
	*Header
	*Footer
}

func NewPage() *Page {
	db := pkg.OpenDB()
	RecentPosts := []*pkg.News{}

	db.Order("created_at desc").Limit(2).Find(&RecentPosts)

	return &Page{
		Data: map[string]any{},
		Header: &Header{
			ActiveMenu: "",
		},
		Footer: &Footer{
			RecentPosts,
		},
	}
}

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
		TopStory   *pkg.News
		MostViews  []*pkg.News
		Features   []*pkg.News
	)
	var g errgroup.Group

	g.Go(func() error {
		return db.Preload("User").Order("created_at desc").Limit(5).Find(&LatestNews).Error
	})

	g.Go(func() error {
		return db.Order("RAND()").First(&TopStory).Error
	})

	g.Go(func() error {
		return db.Preload("User").Order("`views` desc").Limit(5).Find(&MostViews).Error
	})

	g.Go(func() error {
		return db.Preload("Category").Order("RAND()").Limit(4).Find(&Features).Error
	})

	if err := g.Wait(); err != nil {
		ErrorHandler(w, r, err)
		return
	}

	data := NewPage()
	data.Header.ActiveMenu = "home"
	data.Data = map[string]any{
		"LatestNews": LatestNews,
		"TopStory":   TopStory,
		"MostViews":  MostViews,
		"Features":   Features,
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

	data := NewPage()

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

func ErrorHandler(w http.ResponseWriter, _ *http.Request, err error) {
	log.Println(err)

	files := []string{
		"./ui/html/500.html",
		"./ui/html/base.html",
	}

	w.WriteHeader(http.StatusInternalServerError)

	data := NewPage()

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
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

	data := NewPage()
	data.Header.ActiveMenu = "contact"

	if err != nil {
		ErrorHandler(w, r, err)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
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

	db := pkg.OpenDB()
	var g errgroup.Group

	News := []*pkg.News{}
	var NewsTotal int64

	g.Go(func() error {
		return db.Preload("User").Scopes(pkg.Paginate(r)).Order("created_at desc").Find(&News).Error
	})

	g.Go(func() error {
		return db.Model(pkg.News{}).Count(&NewsTotal).Error
	})

	if err := g.Wait(); err != nil {
		ErrorHandler(w, r, err)
		return
	}

	page, _, Pages, TotalPages, PrevPage, NextPage := pkg.PaginateData(r, int(NewsTotal))

	data := NewPage()
	data.Header.ActiveMenu = "news"
	data.Data = map[string]any{
		"News":        News,
		"Pages":       Pages,
		"TotalPages":  TotalPages,
		"CurrentPage": page,
		"PrevPage":    PrevPage,
		"NextPage":    NextPage,
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

	data := NewPage()

	data.Data = map[string]any{
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
