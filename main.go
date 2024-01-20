package main

import (
	"go-scoresheet/database"
	"go-scoresheet/router"
	"log"
	"os"
)

// @title GO-Scoresheet
// @description This is a sample server.
// @version 2.0
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// Inisialisasi koneksi database
	database.StartDB()
	log.Println("Database initialized successfully")

	// Inisialisasi rute-rute aplikasi dari package router
	app := router.InitializeRoutesMain()

	// Jalankan aplikasi pada port tertentu
	port := os.Getenv("APP_PORT")
	log.Printf("Server started on %s\n", port)
	log.Fatal(app.Listen(port))

}
