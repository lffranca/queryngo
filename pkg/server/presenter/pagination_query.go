package presenter

type PaginationQuery struct {
	Offset *int `form:"offset"`
	Limit  *int `form:"limit"`
}
