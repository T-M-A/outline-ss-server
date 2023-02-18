package postgres

import (
	"context"
	"errors"

	"github.com/Jigsaw-Code/outline-ss-server/internal/dbase"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ dbase.Repo = (*client)(nil)

// func CreateDBConn(url string) (*client, error) {
// 	conn, err := pgx.Connect(context.Background(), url)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		return nil, err
// 	}
// 	defer conn.Close(context.Background())
// 	return &client{conn: conn}, nil
// }

func NewPool(p *pgxpool.Pool) *client {
	return &client{conn: p}
}

type client struct {
	conn *pgxpool.Pool
}

func (c *client) Load() ([]dbase.ConfigRow, error) {
	var cr []dbase.ConfigRow
	ctx := context.Background()
	switch rows, err := c.conn.Query(ctx, "select id,port,cipher,secret from clients"); {
	case errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded):
		return nil, err
	case err != nil:
		return nil, errors.New("ERR")
	default:
		defer rows.Close()
		for rows.Next() {
			v := dbase.ConfigRow{}
			if err := rows.Scan(&v); err != nil {
				return nil, err
			}
			cr = append(cr, v)
		}
		return cr, nil
	}
}
