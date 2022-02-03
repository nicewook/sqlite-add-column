# Add column to SQLite table

I wanted to add column to existing SQLite table. 

I tried to add column just in case of column is not exist. but it is actually no need to do. 
Because, add column will return error if the column already exist

## addColumn

Just add column

reference: https://bit.ly/3ok87JV

```
ALTER TABLE table_name
  ADD new_column_name column_definition;
```

my code

``` go
func addColumn(db *sql.DB, table, columnName string) error {
	_, err := db.Exec(fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", table, columnName))
	return err
}
```

## existColumn

I leave the code for the record. 
This function will tell you if the columnn is already exist or not.

Keyword is `pragma_table_info`

``` go
func existColumn(db *sql.DB, table, columnName string) bool {

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
```

## Trouble shooting

I encountered error below and solved with the link.
`the spelling of pragma_table_info (with underline) and the table name should be within quotes`

reference link: https://stackoverflow.com/a/50951476

``` bash
Q) sqlite> SELECT * FROM PRAGMA table_info(results); returns error: Error: near "(": syntax error 
A) sqlite> SELECT * FROM pragma_table_info('results'); should work, note
```