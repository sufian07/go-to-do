package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("My name is Sufian")
	http.HandleFunc("/", taskList)
	http.ListenAndServe(":8082", nil)
}

func taskList(writer http.ResponseWriter, request *http.Request) {
	var taskList = []string{"Create http server", "Add route for task list", "Add route for task add", "Add route for task edit", "Add route for task delete"}
	for _, task := range taskList {
		fmt.Fprintln(writer, task)
	}
}
