package main

import (
	"encoding/json"
	"main/pkg"
	"net/http"

	"gorm.io/gorm/clause"
)

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func NewsGetApiHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	db := pkg.OpenDB()

	news := []*pkg.NewsApi{}

	db.Preload(clause.Associations).Model(&pkg.News{}).Scopes(pkg.Paginate(r)).Find(&news)

	json.NewEncoder(w).Encode(news)
}
