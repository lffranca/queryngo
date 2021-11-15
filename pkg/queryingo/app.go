package queryingo

import (
	"github.com/lffranca/queryngo/pkg/config"
	"sync"
)

func Run() (err error) {
	conf, err := config.New(nil)
	if err != nil {
		return
	}

	wg := &sync.WaitGroup{}
	for _, client := range conf.Servers {
		wg.Add(1)
		go clientRun(wg, client)
	}

	wg.Wait()

	return
}
