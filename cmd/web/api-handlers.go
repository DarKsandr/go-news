package main

import (
	"encoding/json"
	"main/pkg"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm/clause"
)

var validate = validator.New()

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func validationError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

// NewsGetApiHandler godoc
// @Summary      Get news list
// @Description  Returns a paginated list of news items
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

// NewsCreateApiHandler godoc
// @Summary      Create news
// @Description  Create news
// @Tags         news
// @Accept       json
// @Produce      json
// @Param        news  body      pkg.NewsRequest  true  "News data"
// @Success      200  {object}  pkg.NewsApi
// @Router       /news [post]
func NewsCreateApiHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)

	var req pkg.NewsRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		validationError(w, err)
		return
	}
	db := pkg.OpenDB()

	news := &pkg.NewsApi{
		Title:      req.Title,
		Content:    req.Content,
		ImageUrl:   req.ImageUrl,
		ReadTime:   req.ReadTime,
		Views:      req.Views,
		ShareCount: req.ShareCount,
		UserID:     req.UserID,
		CategoryID: req.CategoryID,
	}

	db.Model(&pkg.News{}).Create(news)

	db.Model(&pkg.News{}).Preload(clause.Associations).First(news, news.ID)

	json.NewEncoder(w).Encode(news)
}
