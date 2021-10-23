package domain

type InputQuery struct {
	Query      string
	Schema     string
	Table      string
	Variables  map[string][]int
	Offset     int
	Limit      int
	OrderBy    [][]string
	Search     string
	TemplateID string
}
