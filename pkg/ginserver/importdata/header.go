package importdata

type headerBind struct {
	Sub string `header:"X-Sub" binding:"required"`
}
