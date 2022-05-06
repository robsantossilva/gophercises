package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Phone struct {
	ID     int
	Number string
}

func Open(driverName, dataSource string) (*DB, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

type DB struct {
	db *sql.DB
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Seed() error {
	for _, phone := range getPhoneNumbers() {
		if _, err := insertPhone(db.db, phone); err != nil {
			return err
		}
	}
	return nil
}

func Migrate(driverName, dataSource string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = createPhoneNumbersTable(db)
	if err != nil {
		return err
	}
	return db.Close()
}

func createPhoneNumbersTable(db *sql.DB) error {
	statement := `
	CREATE TABLE IF NOT EXISTS phone_numbers (
		id SERIAL,
		value VARCHAR(255)
	)`
	_, err := db.Exec(statement)
	return err
}

func Reset(driverName, dataSource, dbName string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = resetDB(db, dbName)
	if err != nil {
		return err
	}
	return db.Close()
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	stmt := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(stmt, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (db *DB) AllPhone() ([]Phone, error) {
	stmt := `SELECT * FROM phone_numbers`
	row, err := db.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var ret []Phone

	for row.Next() {
		var p Phone
		if err := row.Scan(&p.ID, &p.Number); err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}

	return ret, nil
}

func (db *DB) FindPhone(number string) (*Phone, error) {
	var p Phone
	stmt := `SELECT * FROM phone_numbers WHERE value=$1`
	row := db.db.QueryRow(stmt, number)
	err := row.Scan(&p.ID, &p.Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &p, nil
}

func (db *DB) UpdatePhone(p *Phone) error {
	stmt := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.db.Exec(stmt, p.ID, p.Number)
	return err
}

func (db *DB) DeletePhone(id int) error {
	stmt := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.db.Exec(stmt, id)
	return err
}

func getPhoneNumbers() []string {
	return []string{
		"ZZZ 1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
}

func (db *DB) GetPhone(id int) (string, error) {
	var number string
	stmt := `SELECT * FROM phone_numbers WHERE id=$1`
	row := db.db.QueryRow(stmt, id)
	err := row.Scan(&id, &number)
	if err != nil {
		return "", err
	}
	return number, nil
}
