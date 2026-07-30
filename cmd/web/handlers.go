package main

import (
	"html/template"
	"log"
	"main/pkg"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

type Header struct {
	ActiveMenu string
}

type Footer struct {
	RecentPosts []*pkg.News
}

type Page struct {
	Data       map[string]any
	Date       time.Time
	Categories []*pkg.Category
	*Header
	*Footer
}

func NewPage() (*Page, error) {
	db := pkg.OpenDB()
	var g errgroup.Group

	RecentPosts := []*pkg.News{}
	Categories := []*pkg.Category{}

	g.Go(func() error {
		return db.Order("created_at desc").Limit(2).Find(&RecentPosts).Error
	})

	g.Go(func() error {
		return db.Find(&Categories).Error
	})

	if err := g.Wait(); err != nil {
		return &Page{}, err
	}

	return &Page{
		Data:       map[string]any{},
		Date:       time.Now(),
		Categories: Categories,
		Header: &Header{
			ActiveMenu: "",
		},
		Footer: &Footer{
			RecentPosts,
		},
	}, nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := NewPage()
	if err != nil {
		ErrorHandler(w, r, err, data)
		return
	}

	if r.URL.Path != "/" {
		NotFoundHandler(w, r, data)
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
		WhatNew    []*pkg.Category
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

	WhatNew = data.Categories
	for _, category := range WhatNew {
		g.Go(func() error {
			return db.Limit(6).Where("category_id = ?", category.ID).Find(&category.News).Error
		})
	}

	if err := g.Wait(); err != nil {
		ErrorHandler(w, r, err, data)
		return
	}

	WhatNewFiltered := []*pkg.Category{}
	for _, v := range WhatNew {
		if len(v.News) > 0 {
			WhatNewFiltered = append(WhatNewFiltered, v)
		}
	}

	data.Header.ActiveMenu = "home"
	data.Data = map[string]any{
		"LatestNews": LatestNews,
		"TopStory":   TopStory,
		"MostViews":  MostViews,
		"Features":   Features,
		"WhatNew":    WhatNewFiltered,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		ErrorHandler(w, r, err, data)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		ErrorHandler(w, r, err, data)
		return
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request, data *Page) {
	files := []string{
		"./ui/html/404.html",
		"./ui/html/base.html",
	}

	w.WriteHeader(http.StatusNotFound)

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		ErrorHandler(w, r, err, data)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		ErrorHandler(w, r, err, data)
		return
	}
}

func ErrorHandler(w http.ResponseWriter, _ *http.Request, err error, data *Page) {
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
	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	data, err := NewPage()
	if err != nil {
		ErrorHandler(w, r, err, data)
		return
	}
	data.Header.ActiveMenu = "contact"

	files := []string{
		"./ui/html/contact.html",
		"./ui/html/base.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		ErrorHandler(w, r, err, data)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		ErrorHandler(w, r, err, data)
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

	data, err := NewPage()
	if err != nil {
		ErrorHandler(w, r, err, data)
		return
	}
	data.Header.ActiveMenu = "news"

	News := []*pkg.News{}
	var NewsTotal int64

	g.Go(func() error {
		return db.Preload("User").Scopes(pkg.Paginate(r)).Order("created_at desc").Find(&News).Error
	})

	g.Go(func() error {
		return db.Model(pkg.News{}).Count(&NewsTotal).Error
	})

	if err := g.Wait(); err != nil {
		ErrorHandler(w, r, err, data)
		return
	}

	page, _, Pages, TotalPages, PrevPage, NextPage := pkg.PaginateData(r, int(NewsTotal))

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
		ErrorHandler(w, r, err, data)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		ErrorHandler(w, r, err, data)
		return
	}
}

func NewsDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	db := pkg.OpenDB()

	var news pkg.News

	data, err := NewPage()
	if err != nil {
		ErrorHandler(w, r, err, data)
		return
	}

	if err := db.Preload("User").First(&news, id).Error; err != nil {
		NotFoundHandler(w, r, data)
		return
	}

	var AlsoLike []pkg.News
	if err := db.Order("RAND()").Limit(2).Find(&AlsoLike).Error; err != nil {
		ErrorHandler(w, r, err, data)
		return
	}

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
		ErrorHandler(w, r, err, data)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		ErrorHandler(w, r, err, data)
		return
	}
}
