package main

import (
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
	phonedb "github.com/robsantossilva/gophercises/phone/db"
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
	//var db *sql.DB

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, pass)
	must(phonedb.Reset("postgres", psqlInfo, dbnm))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbnm)
	must(phonedb.Migrate("postgres", psqlInfo))

	db, err := phonedb.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(db.Seed())

	// phone, _ := db.GetPhone(1)
	// fmt.Println(phone)

	phones, err := db.AllPhone()
	must(err)

	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		number := normalize(p.Number)
		if number != p.Number {
			fmt.Println("Updating or removing...", number)
			existing, err := db.FindPhone(number)
			must(err)
			if existing != nil {
				must(db.DeletePhone(p.ID))
			} else {
				p.Number = number
				must(db.UpdatePhone(&p))
			}
		} else {
			fmt.Println("No changes required")
		}
		fmt.Printf("\n")
	}
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
