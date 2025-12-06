package database

import (
	"database/sql"
	"eaglebank/internal/infrastructure/dbexecutor"
)

func FetchOneRow(
	db dbexecutor.DbExecutor,
	sql string,
	mapping []any,
	params ...any,
) (bool, error) {
	rows, err := db.Query(sql, params...)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(mapping...)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}

func FetchRowsAnMapToEntities[T, T2 any](
	db dbexecutor.DbExecutor,
	sql string,
	mapperCb func(*T) []any,
	mappedCb func(*T) (T2, error),
	params ...any,
) ([]T2, error) {
	var results []T2
	rows, err := db.Query(sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var input T
		fields := mapperCb(&input)
		err = rows.Scan(fields...)
		if err != nil {
			return nil, err
		}
		entity, err := mappedCb(&input)
		if err != nil {
			return nil, err
		}

		results = append(results, entity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// This runs on a new transaction
func ExecuteSQL(db dbexecutor.DbExecutor, sql string, params ...any) (sql.Result, error) {
	// tx, err := db.Begin()
	// if err != nil {
	// 	return
	// }

	// defer func() {
	// 	if err != nil {
	// 		txErr := tx.Rollback()
	// 		if txErr != nil {
	// 			err = txErr
	// 		}
	// 		return
	// 	}
	// 	txErr := tx.Commit()
	// 	if txErr != nil {
	// 		err = txErr
	// 	}
	// }()

	return db.Exec(sql, params...)
}
