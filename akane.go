package akane

import (
	"database/sql"
)

type DB struct {
	RawDB *sql.DB
}

func (db *DB) SelectOne(query string, args ...interface{}) (interface{}, error) {
	row := db.RawDB.QueryRow(query, args...)
	dest := new(interface{})

	err := row.Scan(dest)
	if err != nil {
		return nil, err
	}

	switch v := (*dest).(type) {
	case []uint8:
		return string(v), nil
	default:
		return *dest, nil
	}
}

func (db *DB) SelectRow(query string, args ...interface{}) (map[string]interface{}, error) {
	rows, err := db.RawDB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		values := make([]interface{}, len(cols))
		for i := range values {
			values[i] = new(interface{})
		}
		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		dest := map[string]interface{}{}
		for i, col := range cols {
			val := *(values[i].(*interface{}))
			switch v := val.(type) {
			case []uint8:
				dest[col] = string(v)
			default:
				dest[col] = v
			}
		}

		return dest, nil
	}

	return nil, nil
}

func (db *DB) SelectAll(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.RawDB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var dests []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(cols))
		for i := range values {
			values[i] = new(interface{})
		}
		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		dest := map[string]interface{}{}
		for i, col := range cols {
			val := *(values[i].(*interface{}))
			switch v := val.(type) {
			case []uint8:
				dest[col] = string(v)
			default:
				dest[col] = v
			}
		}

		dests = append(dests, dest)
	}

	return dests, nil
}


func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &DB{RawDB: db}, nil
}
