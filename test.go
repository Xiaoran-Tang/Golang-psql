// author:  Xiaoran Tang
// date:    Nov 03 2017

// This is a sample implementation of postgreSQL in Golang.
// The program simulates the userinfo database for CBRE Group, Inc. 

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	/* initialization: create a database handle */
	fmt.Println("---CBRE userinfo postgreSQL program starts---")
	db, err := sql.Open("postgres", "user=postgres password=password dbname=postgres")
	checkErr(err)



	/* insert userinfo into database */
	fmt.Println("\n---insert userinfo into database---")
	
	//creates a prepared statement for use
	stmt, err := db.Prepare("INSERT INTO userinfo(username,phone,homebranch) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)
   
	//no row will return after Execer
	res, err := stmt.Exec("Tang", "5120001234", "CBRE Dallas")
	checkErr(err)

	var lastInsertId int
	//another way to insert
	err = db.QueryRow("INSERT INTO userinfo(username,phone,homebranch) VALUES($1,$2,$3) returning uid;", "Aziz", "5120001111", "CBRE Austin").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("Last inserted id is: ", lastInsertId)



	/* update data */
	fmt.Println("\n---update data---")
	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err = stmt.Exec("Wong", 1)
	checkErr(err)
   
	//how many rows got affected by the 'update' query
	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("how many rows got affected by the 'update' query? ", affect)



	/* scan all existing rows */
	fmt.Println("\n---scan all existing rows---")
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var phone string
		var homebranch string
		err = rows.Scan(&uid, &username, &phone, &homebranch)
		checkErr(err)
		fmt.Println("record number: ", uid)
		fmt.Println("user name: ", username)
		fmt.Println("phone: ", phone)
		fmt.Println("homebranch: ", homebranch, "\n")
	}



	/* delete all data */
	fmt.Println("\n---delete all data---")
	stmt, err = db.Prepare("delete from userinfo where uid>=$1")
	checkErr(err)

	res, err = stmt.Exec(1)
	checkErr(err)

	//how many rows got affected
	affect, err = res.RowsAffected()
	checkErr(err)
	
	fmt.Println("how many rows got affected by the 'delete' query? ", affect)
	db.Close()

	fmt.Println("\n---CBRE userinfo postgreSQL program ends---")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
