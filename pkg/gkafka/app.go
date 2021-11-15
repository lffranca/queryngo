package gkafka

import "errors"

func New(options *Options) (*Server, error) {
	if options == nil {
		return nil, errors.New("options is required")
	}

	if err := options.validate(); err != nil {
		return nil, err
	}

	server := new(Server)
	server.common.Server = server
	server.brokers = options.Brokers
	server.network = options.Network
	server.processFileTopic = options.ProcessFileTopic
	server.processedFileTopic = options.ProcessedFileTopic
	server.processedFileRepository = options.ProcessedFileRepository
	server.ProcessFile = (*ProcessFileService)(&server.common)

	return server, nil
}

type Server struct {
	common                  service
	brokers                 []string
	network                 string
	processFileTopic        string
	processedFileTopic      string
	processedFileRepository ProcessedFileRepository
	ProcessFile             *ProcessFileService
}

type service struct {
	Server *Server
}
