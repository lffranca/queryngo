package querying

type queryStringBind struct {
	QueryID  *string `form:"query_id" json:"query_id" binding:"required"`
	FormatID *string `form:"format_id" json:"format_id" binding:"required"`
}
