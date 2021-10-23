package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lffranca/queryngo/domain/importdata"
	"github.com/lffranca/queryngo/domain/querying"
	"github.com/lffranca/queryngo/pkg/util"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
)

func New(conn *string) (*Client, error) {
	if conn == nil {
		return nil, errors.New("conn is required")
	}

	db, err := sql.Open("postgres", *conn)
	if err != nil {
		return nil, err
	}

	client := new(Client)

	client.db = db
	client.common.client = client
	client.File = (*FileService)(&client.common)
	client.Template = (*TemplateService)(&client.common)
	client.Querying = (*QueryingService)(&client.common)

	// domain services
	client.SaveFileKeyImportData = (*SaveFileKeyService)(&client.common)

	return client, nil
}

type service struct {
	client *Client
}

type Client struct {
	db                    *sql.DB
	common                service
	File                  *FileService
	Template              *TemplateService
	Querying              querying.AbstractQuerying
	SaveFileKeyImportData importdata.AbstractDatabase
}

func (pkg *Client) query(ctx context.Context, query string, variables []interface{}) (*sql.Rows, error) {
	if variables == nil || len(variables) <= 0 {
		return pkg.db.QueryContext(ctx, query)
	}

	var args []interface{}
	for _, varItem := range variables {
		args = append(args, pq.Array(varItem))
	}

	return pkg.db.QueryContext(ctx, query, args...)
}

func (pkg *Client) querying(ctx context.Context, query string, variables []interface{}) ([]string, []string, [][]interface{}, error) {
	rows, err := pkg.query(ctx, query, variables)
	if err != nil {
		return nil, nil, nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("ERROR CLOSE ROWS: ", err)
		}
	}()

	return util.SQLRowModel(rows)
}

func (pkg *Client) DB() *sql.DB {
	return pkg.db
}

func (pkg *Client) Close() {
	if err := pkg.db.Close(); err != nil {
		panic(err)
	}
}
