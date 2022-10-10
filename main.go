package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID       int
	Name     string
	Email    string
	Age      int
	Division string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db-go"
)

var (
	db  *sql.DB
	err error
)

func main() {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", sqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	// CreateEmployee()
	// GetEmployees()
	UpdateEmployee()
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `
	INSERT INTO employees (name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning *`

	err = db.QueryRow(sqlStatement, "Kresna Vespri", "kresna@gmail.com", 21, "IT").Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)
}

func GetEmployees() {
	var results = []Employee{}

	sqlStatement := `SELECT * FROM employees`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("Employees :", results)
}

func UpdateEmployee() {
	sqlStatement := `
	UPDATE employees
	SET name = $2, email = $3, age = $4, division = $5
	WHERE id = $1;`

	res, err := db.Exec(sqlStatement, 1, "Kresna VW", "vw@gmail.com", 21, "Software Engineer")

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println("Updated Data Amount :", count)
}
