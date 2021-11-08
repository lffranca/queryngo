package presenter

type CommonHeader struct {
	Sub string `header:"X-Sub" binding:"required"`
}
