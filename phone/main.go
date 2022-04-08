package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "root"
	pass = "root"
	dbnm = "gophercises_phone"
)

func main() {

	var err error
	var db *sql.DB

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, pass)
	// db, err = sql.Open("postgres", psqlInfo)
	// must(err)
	// err = resetDB(db, dbnm)
	// must(err)
	// db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbnm)
	db, err = sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	// must(createPhoneNumbersTable(db))
	// for _, phone := range getPhoneNumbers() {
	// 	id, err := insertPhone(db, phone)
	// 	must(err)
	// 	fmt.Println("id=", id)
	// }

	// phone, _ := getPhone(db, 1)
	// fmt.Println(phone)

	phones, err := allPhone(db)
	must(err)
	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		number := normalize(p.number)
		if number != p.number {
			fmt.Println("Updating or removing...", number)
			existing, err := findPhone(db, number)
			must(err)
			if existing != nil {
				must(deletePhone(db, p.id))
			} else {
				p.number = number
				must(updatePhone(db, p))
			}
		} else {
			fmt.Println("No changes required")
		}
		fmt.Printf("\n")
	}
}

func getPhone(db *sql.DB, id int) (string, error) {
	var number string
	stmt := `SELECT * FROM phone_numbers WHERE id=$1`
	row := db.QueryRow(stmt, id)
	err := row.Scan(&id, &number)
	if err != nil {
		return "", err
	}
	return number, nil
}

func findPhone(db *sql.DB, number string) (*phone, error) {
	var p phone
	stmt := `SELECT * FROM phone_numbers WHERE value=$1`
	row := db.QueryRow(stmt, number)
	err := row.Scan(&p.id, &p.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &p, nil
}

func updatePhone(db *sql.DB, p phone) error {
	stmt := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(stmt, p.id, p.number)
	return err
}

func deletePhone(db *sql.DB, id int) error {
	stmt := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.Exec(stmt, id)
	return err
}

type phone struct {
	id     int
	number string
}

func allPhone(db *sql.DB) ([]phone, error) {
	stmt := `SELECT * FROM phone_numbers`
	row, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var ret []phone

	for row.Next() {
		var p phone
		if err := row.Scan(&p.id, &p.number); err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}

	return ret, nil
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

func createPhoneNumbersTable(db *sql.DB) error {
	statement := `
	CREATE TABLE IF NOT EXISTS phone_numbers (
		id SERIAL,
		value VARCHAR(255)
	)`
	_, err := db.Exec(statement)
	return err
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

func normalize(phone string) string {
	re := regexp.MustCompile("\\D")
	//re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
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

// func normalize(phone string) string {
// 	var buf bytes.Buffer
// 	for _, ch := range phone {
// 		if ch >= '0' && ch <= '9' {
// 			buf.WriteRune(ch)
// 		}
// 	}
// 	return buf.String()
// }
