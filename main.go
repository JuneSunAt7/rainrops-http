package main

import (
	"fmt"
	"net/http"
)

// handler для корневого пути
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет, мир!")
}

func main() {
	// Устанавливаем обработчик для корневого пути
	http.HandleFunc("/", helloHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
