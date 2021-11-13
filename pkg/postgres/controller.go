package postgres

import (
	"database/sql"
	"errors"
	"github.com/lffranca/queryngo/pkg/postgres/model"
)

func NewController(options *model.Options) (*Controller, error) {
	if options == nil {
		return nil, errors.New("options is required")
	}

	db, err := newDB(&options.PrincipalDB)
	if err != nil {
		return nil, err
	}

	clientDBs := make(map[string]*sql.DB)
	for key, conn := range options.ClientDBs {
		cliDB, err := newDB(&conn)
		if err != nil {
			return nil, err
		}

		clientDBs[key] = cliDB
	}

	controller := new(Controller)
	controller.db = db
	controller.clientDBs = clientDBs

	return controller, nil
}

type Controller struct {
	db        *sql.DB
	clientDBs map[string]*sql.DB
}

func (pkg *Controller) Principal() (*Client, error) {
	return NewClient(pkg.db)
}

func (pkg *Controller) MultiTenant(prefix string) (*Client, error) {
	db, ok := pkg.clientDBs[prefix]
	if !ok {
		return nil, errors.New("")
	}

	return NewClient(db)
}
