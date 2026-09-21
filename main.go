package main

import (
	"fmt"
	"net/http"
)


func main() {
	 db, err := connectDB()

    if err != nil {
        panic(err)
    }

    defer db.Close()

    fmt.Println("Connected to Supabase!")

    postgresStorage := PostgresStorage{
        db: db,
    }

    handler := Handler{
        storage: postgresStorage,
    }
	http.HandleFunc("/", handler.redirectHandler)
	http.HandleFunc("/shorten", handler.shortenHandler)
	fmt.Println("Server running on :8080")
	erre := http.ListenAndServe(":8080", nil)

	fmt.Println(erre)
}
