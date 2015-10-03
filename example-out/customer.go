package main

import (
	"strings"

	"github.com/jackc/pgx"
)

type CustomerRow struct {
	ID           Int32
	FirstName    String
	LastName     String
	BirthDate    Time
	CreationTime Time
}

func CountCustomer(db Queryer) (int64, error) {
	var n int64
	sql := `select count(*) from customer`
	err := db.QueryRow(sql).Scan(&n)
	return n, err
}

func SelectAllCustomer(db Queryer) ([]CustomerRow, error) {
	sql := `select id, first_name, last_name, birth_date, creation_time
from customer`

	var rows []CustomerRow

	dbRows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	for dbRows.Next() {
		var row CustomerRow
		dbRows.Scan(&row.ID, &row.FirstName, &row.LastName, &row.BirthDate, &row.CreationTime)
		rows = append(rows, row)
	}

	if dbRows.Err() != nil {
		return nil, dbRows.Err()
	}

	return rows, nil
}

func SelectCustomerByID(db Queryer, id int32) (*CustomerRow, error) {
	sql := `select id, first_name, last_name, birth_date, creation_time
from customer
where id=$1`

	var row CustomerRow
	err := db.QueryRow(sql, id).Scan(&row.ID, &row.FirstName, &row.LastName, &row.BirthDate, &row.CreationTime)
	if err != nil {
		return nil, err
	}

	return &row, nil
}

func InsertCustomer(db Queryer, row *CustomerRow) error {
	args := pgx.QueryArgs(make([]interface{}, 0, 5))

	var columns, values []string

	row.ID.addInsert("id", &columns, &values, &args)
	row.FirstName.addInsert("first_name", &columns, &values, &args)
	row.LastName.addInsert("last_name", &columns, &values, &args)
	row.BirthDate.addInsert("birth_date", &columns, &values, &args)
	row.CreationTime.addInsert("creation_time", &columns, &values, &args)

	sql := `insert into customer(` + strings.Join(columns, ", ") + `)
values(` + strings.Join(values, ",") + `)
returning id
  `

	return db.QueryRow(sql, args...).Scan(&row.ID)
}

func UpdateCustomer(db Queryer, id int32, row *CustomerRow) error {
	sets := make([]string, 0, 5)
	args := pgx.QueryArgs(make([]interface{}, 0, 5))
	row.ID.addUpdate("id", &sets, &args)
	row.FirstName.addUpdate("first_name", &sets, &args)
	row.LastName.addUpdate("last_name", &sets, &args)
	row.BirthDate.addUpdate("birth_date", &sets, &args)
	row.CreationTime.addUpdate("creation_time", &sets, &args)

	if len(sets) == 0 {
		return nil
	}

	sql := `update customer set ` + strings.Join(sets, ", ") + ` where id=` + args.Append(id)

	commandTag, err := db.Exec(sql, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return pgx.ErrNoRows
	}
	return nil
}

func DeleteCustomer(db Queryer, id int32) error {
	sql := `delete from customer where id=$1`

	commandTag, err := db.Exec(sql, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return pgx.ErrNoRows
	}
	return nil
}
