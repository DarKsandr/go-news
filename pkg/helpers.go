package pkg

import (
	"net/http"
	"strconv"
)

func PaginateParams(r *http.Request) (page, pageSize int) {
	q := r.URL.Query()
	page, _ = strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ = strconv.Atoi(q.Get("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 20
	}

	return
}

func PaginateData(r *http.Request, total int) (page int, PageSize int, pages []int, TotalPages int, PrevPage int, NextPage int) {
	page, PageSize = PaginateParams(r)

	TotalPages = (total + PageSize - 1) / PageSize
	pages = make([]int, 0, TotalPages)
	for i := 1; i <= TotalPages; i++ {
		pages = append(pages, i)
	}

	PrevPage = page - 1
	NextPage = page + 1

	return
}
