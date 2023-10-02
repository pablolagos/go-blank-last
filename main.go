package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// Environment variables are defined in service.yaml file.
	// Let's read the ADMIN_USERNAME and ADMIN_PASSWORD example variables:
	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	// Write a message to the console
	fmt.Println("This is a golang example custom service.")
	fmt.Printf("A web server will be started at port 8080\n\n")
	fmt.Println("Also, two example variables have been read:")
	fmt.Printf("Admin username: %s\n", adminUsername)
	fmt.Printf("Admin password: %s\n", adminPassword)
	fmt.Println()
	fmt.Println("Visit web interface to see this message there.")

	// Create and register web server handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "This is a golang example custom service.")
		fmt.Fprintf(w, "This web server is listening at port 8080\n\n")
		fmt.Fprintln(w, "Also, two example variables have been read:")
		fmt.Fprintf(w, "Admin username: %s\n", adminUsername)
		fmt.Fprintf(w, "Admin password: %s\n", adminPassword)

	})

	// Start web server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":8080"), nil))
}
