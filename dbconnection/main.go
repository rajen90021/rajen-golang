package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type book struct {
	id     string
	title  string
	author string
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/crudapi")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	/////*******************************************************************//////////////////////////////***********************************
	// insert, err := db.Query("INSERT INTO book VALUES('5','learning gin framework','raj-jainnnn')") ////inserting the value in db
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// fmt.Println("inserted sucessfully")

	//reterving the element
	results, err := db.Query("SELECT * FROM book")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	var books []book //getting all the value from db and storing in slice of structure
	for results.Next() {
		var book book
		err = results.Scan(&book.id, &book.title, &book.author)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, book)
	}

	err = results.Err()
	if err != nil { ///checking error while iteration
		panic(err.Error())
	}
	for _, book := range books { ///////getting all the data
		fmt.Println("id", book.id)
		fmt.Println("Title:", book.title)
		fmt.Println("Author:", book.author)

	}

	//**************************************************************second opproch ******************************************************
	//Execute the DELETE statement
	delete, err := db.Exec("DELETE FROM book WHERE id = ?", 5) // Delete row with ID = 1
	if err != nil {
		panic(err.Error())
	}

	// Check the number of affected rows
	sucesss, err := delete.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Deleted %d row(s) from the table\n", sucesss)

	//***************************************************************************************************updating
	updatedvalue, err := db.Exec("UPDATE book SET title = 'trying2', author = 'rajen' WHERE id = 2")
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := updatedvalue.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Updated %d row(s) in the table\n", rowsAffected)
}
