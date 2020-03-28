package db

import (
	"database/sql"
	"log"
	"oraTop/db/SQL"
)

type OracleDB struct {
	*sql.DB
}

func (o OracleDB) GetSession() (records [][]string) {
	rows, err := OraDB.Query(SQL.OraSession)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}

	columns, _ := rows.Columns()
	// create record to receive row
	record := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(record))
	for i := range record {
		scanArgs[i] = &record[i]
	}

	records = make([][]string, 0, 40)
	records = append(records, columns)

	for rows.Next() {
		row := make([]string, 0, len(record))
		// scan row
		rows.Scan(scanArgs...)
		// change rawBytes to string
		for _, v := range record {
			row = append(row, string(v))
		}
		records = append(records, row)
	}
	return
}
