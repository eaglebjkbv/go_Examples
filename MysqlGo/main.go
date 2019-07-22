package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Merhaba")
	db, err := sql.Open("mysql", "bv2:12345@tcp(localhost:3306)/restproje")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("Select * from plcdeger")
	for rows.Next() {
		var (
			id    string
			tarih string
			saat  string
			deger string
		)
		if err := rows.Scan(&id, &tarih, &saat, &deger); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %s tarih: %s  saat:%s deger: %s \n", id, tarih, saat, deger)

	}

}
