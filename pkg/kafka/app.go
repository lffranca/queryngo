package kafka

func New() (*Server, error) {
	server := new(Server)
	return server, nil
}

type Server struct{}
