package main

import (
	"go-scoresheet/database"
	"go-scoresheet/router"
	"log"
)

func main() {
	// Inisialisasi koneksi database
	database.StartDB()

	// Inisialisasi rute-rute aplikasi dari package router
	app := router.InitializeRoutesMain()

	// Jalankan aplikasi pada port tertentu
	port := ":3000"
	log.Printf("Server started on %s\n", port)
	log.Fatal(app.Listen(port))

}
