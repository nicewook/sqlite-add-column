package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// type Person struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// type PersonWithHeight struct {
// 	Name   string `json:"name"`
// 	Age    int    `json:"age"`
// 	Height int    `json:"height"`
// }

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create myTable
	createTableOneSQL := `CREATE TABLE IF NOT EXISTS myTable (
		"ID"     INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"Name"   TEXT
	);`

	_, err = db.Exec(createTableOneSQL)
	log.Println(err)

	// check if height column is exist
	rows, err := db.Query(`select name from pragma_table_info("myTable")`)
	if err != nil {
		log.Fatal(err)
		// fmt.Println(err)
	}
	defer rows.Close()

	// log.Println(rows.Columns())
	// log.Println(rows.ColumnTypes())
	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		log.Println("name: ", name)
		names = append(names, name)
	}
	log.Println("names: ", names)
}

// 	for rows.Next() {
// 		var alb Album
// 		if err := rows. .Scan(&alb.ID, &alb.Title, &alb.Artist,
// 				&alb.Price, &alb.Quantity); err != nil {
// 				return albums, err
// 		}
// 		albums = append(albums, album)
// }

// 	var myDepartment string
// 	if err := row.Scan(&myDepartment); err != nil {
// 		log.Println(err)
// 	}
// 	log.Println(myDepartment)

// 	row3 := db.QueryRow("select name from myTable")
// 	if row3 == nil {
// 		log.Fatalln("failed")
// 		return
// 	}

// 	var myName string
// 	if err := row3.Scan(&myName); err != nil {
// 		log.Println(err)
// 	}
// 	log.Println(myName)

// 	// inset 1 data to db
// 	person1 := Person1{
// 		"JHS",
// 		47,
// 	}
// 	if _, err := insertRow(db, person1); err != nil { // Create Database Tables
// 		log.Fatal("fail to add person1 to table:", err)
// 	}

// 	// read db
// 	person1s, err := getAll(db)
// 	if err != nil { // Create Database Tables
// 		log.Fatal("fail to add person1 to table:", err)
// 	}
// 	log.Printf("person1s: %+v\n", person1s)

// 	row2 := db.QueryRow("select department from myTable")
// 	if row2 == nil {
// 		log.Fatalln("failed")
// 		return
// 	}

// 	if err := row2.Scan(&myDepartment); err != nil {
// 		log.Println(err)
// 	}
// 	log.Println(myDepartment)
// check table info
// tableInfo, err := db.Exec("PRAGMA table_info(myTable)")
// if err != nil {
// 	log.Fatal(err)
// }
// log.Printf("%+v\n", tableInfo)

/*
	row := db.QueryRow("PRAGMA user_version")
	if row == nil {
		log.Fatalln("PRAGMA user_version not found")
		return
	}
	var version int
	if err = row.Scan(&version); err != nil {
		log.Fatal(err)
	}
	log.Printf("PRAGMA user_version is %d\n", version)
*/
// row := db.QueryRow("select name from myTable")

// alter db

// add db

// read db

// }
