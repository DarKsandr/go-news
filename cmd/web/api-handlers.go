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

// NewsGetApiHandler godoc
// @Summary      Get news list
// @Description  Returns a paginated list of news items.
// @Tags         news
// @Accept       json
// @Produce      json
// @Success      200  {array}   pkg.NewsApi
// @Router       /news [get]
func NewsGetApiHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	db := pkg.OpenDB()

	news := []*pkg.NewsApi{}

	db.Preload(clause.Associations).Model(&pkg.News{}).Scopes(pkg.Paginate(r)).Find(&news)

	json.NewEncoder(w).Encode(news)
}
