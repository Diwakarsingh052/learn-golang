package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// bad idea // we will use env variables
const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "root"
	dbname   = "user"
)

var db *sql.DB // bad idea to do in production

func main() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	//insert()
	//insert2()
	//update()
	//delete()
	//querySingleRecord()
	queryMultipleRecord()
}

// CRUD  Operation //Create Read Update Delete

func insert() {

	sqlStatement := `
	Insert INTO users (age, email, first_name, last_name)
	Values ($1 , $2, $3, $4 )
`
	_, err := db.Exec(sqlStatement, 31, "abc@email.com", "Raj", "Ahuja")
	if err != nil {
		panic(err)
	}

}

func insert2() {

	var (
		id    int
		email string
	)

	sqlStatement := `
	Insert INTO users (age, email, first_name, last_name)
	Values ($1 , $2, $3, $4 )
	Returning id, email
`
	err := db.QueryRow(sqlStatement, 32, "xyz@email.com", "Ajay", "Ahuja").Scan(&id, &email)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, email)

}

func update() {

	sqlStatement := `
	Update users
	Set last_name = $2 
	where id = $1;
`

	res, err := db.Exec(sqlStatement, 3, "Singh")
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

func delete() {

	sqlStatement := `
	Delete FROM users
	where id = $1;
`
	res, err := db.Exec(sqlStatement, 3)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted", count)

}

func querySingleRecord() {

	sqlStatement := `Select id, email FROM users where id = $1;`

	var (
		id    int
		email string
	)

	row := db.QueryRow(sqlStatement, 4)
	err := row.Scan(&id, &email)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("now rows returned")
	case nil:
		fmt.Println(id, email)
	default:
		panic(err)
	}

}
func queryMultipleRecord() {

	rows, err := db.Query("Select id, first_name FROM users LIMIT $1", 4)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id        int
			firstName string
		)
		err = rows.Scan(&id, &firstName)

		if err != nil {
			panic(err)
		}
		fmt.Println(id, firstName)

	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

}
