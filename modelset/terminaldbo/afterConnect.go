package terminaldbo

import (
	"context"

	"github.com/jackc/pgx/v4"
)

const (
	sqlInsertTerminal = "sqlInsertTerminal"
	sqlUpdateTerminal = "sqlUpdateTerminal"
	sqlActiveTerminal = "sqlActiveTerminal"
	SqlListenTerminal = "sqlListenTerminal"
)

func afterConnect(ctx context.Context, conn *pgx.Conn) (err error) {
	prepare := map[string]string{
		sqlInsertTerminal: `insert into ums.terminal (serial_number,active_code) values ($1, $2) returning id`,
		sqlUpdateTerminal: `update ums.terminal
set
  active_status=$1
  , max_active_session = $2
  , service_status = $3
where
  id = $4`,

		sqlActiveTerminal: `select
  id
  , serial_number
  , active_code
  , active_status
  , active_date
  , max_active_session
  , access_role
  , service_status
  , service_expiration
from
  ums.Active($1,$2,$3)`,

		SqlListenTerminal: `listen terninal_notify`,
	}

	for key, value := range prepare {
		_, err = conn.Prepare(ctx, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}
