package presenter

type ParentURI struct {
	ID *int `form:"parent_id" json:"parent_id" binding:"required"`
}
