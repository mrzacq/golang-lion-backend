package helper

import "database/sql"

func CommitOrCallBack(tx *sql.Tx) {
	err := recover()
	if(err != nil) {
		errorRollback := tx.Rollback()
		Error(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		Error(errorCommit)
	}
}