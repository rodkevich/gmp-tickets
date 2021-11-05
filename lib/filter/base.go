package filter

import (
	"net/url"
	"strconv"
)

const (
	// Query types
	queryParamFields = "fields"
	queryParamSearch = "search"
	// Paging settings
	queryParamPagingDisable         = "disable_paging"
	queryParamPage                  = "page"
	queryParamSize                  = "size"
	responseItemsDefaultPage        = 1
	responseItemsDefaultSizePerPage = 20
)

type Common struct {
	Extended      bool `json:"extended"`
	Search        bool `json:"search"`
	Page          int  `json:"page"`
	Size          int  `json:"size"`
	PagingDisable bool `json:"paging_disable"`
}

func New(queries url.Values) *Common {
	isExtended := has(queries, queryParamFields)
	isSearch := has(queries, queryParamSearch)
	page, _ := strconv.Atoi(queries.Get(queryParamPage))
	sizePerPage, _ := strconv.Atoi(queries.Get(queryParamSize))
	pagingDisable, _ := strconv.ParseBool(queries.Get(queryParamPagingDisable))

	if !has(queries, queryParamSize) {
		sizePerPage = responseItemsDefaultSizePerPage
	}

	if !has(queries, queryParamPage) {
		page = responseItemsDefaultPage
	}
	// offset calculation:
	// page = (page - 1) * sizePerPage + 1
	page = (page - 1) * sizePerPage

	return &Common{
		Extended:      isExtended,
		Search:        isSearch,
		Page:          page,
		Size:          sizePerPage,
		PagingDisable: pagingDisable,
	}
}

func has(queries url.Values, param string) bool {
	return queries.Get(param) != ""
}
