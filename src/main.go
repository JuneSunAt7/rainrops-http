package main

import (
	"fmt"
	"net/http"
)

// handler for root path
func mainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html") 
}

func main() {
	http.HandleFunc("/", mainHandler)

	// Handler for all other static files from the web folder
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	
	fmt.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
