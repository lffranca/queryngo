package queryingo

import (
	"github.com/lffranca/queryngo/pkg/config"
	"sync"
)

func Run() error {
	conf, err := config.New(nil)
	if err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	for _, client := range conf.Servers {
		wg.Add(1)
		go serverClientRun(wg, client)
	}

	wg.Wait()

	return err
}
