package presenter

type ListCommonURI struct {
	Offset *int    `form:"offset" json:"offset"`
	Limit  *int    `form:"limit" json:"limit"`
	Search *string `form:"search" json:"search"`
}
