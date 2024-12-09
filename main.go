package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	todoList := TodoList{Items: []*TodoItem{}}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Add Item")
		fmt.Println("2. Print List")
		fmt.Println("3. Complete Item")
		fmt.Println("4. Delete Item")
		fmt.Println("5. Quit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			fmt.Print("Enter Todo item title: ")
			scanner.Scan()
			title := scanner.Text()
			todoList.AddItem(title)
		case "2":
			todoList.PrintList()
		case "3":
			fmt.Print("Enter ID of Todo item to complete: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			todoList.CompleteItem(id)
		case "4":
			fmt.Print("Enter ID of Todo item to delete: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			todoList.DeleteItem(id)
		case "5":
			return
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}