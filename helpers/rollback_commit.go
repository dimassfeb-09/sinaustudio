package helpers

import (
	"database/sql"
)

func RollbackOrCommit(tx *sql.Tx) {
	err := recover()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return
		}
	} else {
		err := tx.Commit()
		if err != nil {
			return
		}
	}
}
