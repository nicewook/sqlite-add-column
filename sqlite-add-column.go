package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Height int    `json:"height"`
}

const (
	myTable      = "myTable"
	columnHeight = "height"
)

// existColumn checks if a column is exist in the table
// I did the loop twice to have all the column names to be printed
func existColumn(db *sql.DB, table, columnName string) bool {

	// reference: https://stackoverflow.com/a/50951476
	rows, err := db.Query(fmt.Sprintf(`select name from pragma_table_info("%s")`, table))
	if err != nil {
		log.Println(err)
		return false
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Println(err)
			return false
		}
		names = append(names, name)
	}
	fmt.Printf("--all the columns in the table: %v\n", names)
	for _, name := range names {
		if name == columnName {
			fmt.Printf("--there is column [%s] in table\n", columnName)
			return true
		}
	}
	fmt.Printf("--there is no column [%s] in table\n", columnName)
	return false
}

// addColumn trys to add a column to the table
// reference: https://stackoverflow.com/questions/4253804/insert-new-column-into-table-in-sqlite
func addColumn(db *sql.DB, table, columnName string) error {
	_, err := db.Exec(fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", table, columnName))
	return err
}

func insertPerson(db *sql.DB, table string, p Person) error {

	insertPersonSQL := fmt.Sprintf(`INSERT INTO %s (  
		Name,
		Age,
		Height          
	) VALUES (?, ?, ?)`, table)

	_, err := db.Exec(insertPersonSQL, p.Name, p.Age, p.Height)
	return err
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create myTable and exist check
	createTableOneSQL := `CREATE TABLE IF NOT EXISTS myTable (
		"ID"     INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"Name"   TEXT,
		"Age"    INTEGER
	);`

	if _, err = db.Exec(createTableOneSQL); err != nil {
		log.Fatal(err)
	}
	existColumn(db, myTable, columnHeight)
	for i := 0; i < 3; i++ {
		if err := addColumn(db, myTable, columnHeight); err != nil {
			log.Println(err)
		} else {
			log.Printf("column [%s] added to table [%s]", columnHeight, myTable)
		}
	}

	existColumn(db, myTable, columnHeight)

	// insert row and exist check
	personOne := Person{"Kim", 32, 176}
	if err := insertPerson(db, myTable, personOne); err != nil {
		log.Fatal(err)
	}
	existColumn(db, myTable, "height")

	if err := addColumn(db, myTable, columnHeight); err != nil {
		log.Println(err)
	}
}
