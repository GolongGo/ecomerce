package helper


type PaginationParams struct {
	Page     int    `schema:"page"`
	PageSize int    `schema:"pageSize"`
	Search   string `schema:"search"`
}