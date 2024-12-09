package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TodoItem represents a single todo item
type TodoItem struct {
	Title     string
	OriginalID int
	Completed bool
	CreatedAt time.Time
}

// TodoList represents a list of todo items
type TodoList struct {
	Items []*TodoItem
}

// AddItem adds a new todo item to the list
func (tl *TodoList) AddItem(title string) {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(10000) // Generate a random ID between 0 and 9999
	newItem := &TodoItem{
		Title:     title,
		OriginalID: id,
		Completed: false,
		CreatedAt: time.Now(),
	}
	tl.Items = append(tl.Items, newItem)
	fmt.Printf("Task '%s' added with ID %d.\n", title, id)
}

// CompleteItem marks a todo item as completed
func (tl *TodoList) CompleteItem(id int) {
	for _, item := range tl.Items {
		if item.OriginalID == id {
			item.Completed = true
			fmt.Printf("Task '%s' completed.\n", item.Title)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", id)
}

// DeleteItem deletes a todo item from the list
func (tl *TodoList) DeleteItem(id int) {
	for i, item := range tl.Items {
		if item.OriginalID == id {
			tl.Items = append(tl.Items[:i], tl.Items[i+1:]...)
			fmt.Printf("Task '%s' deleted.\n", item.Title)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", id)
}


// PrintList prints the Todo list to the console
func (tl *TodoList) PrintList() {
	fmt.Println("Todo List:")
	for _, item := range tl.Items {
		status := "Not Completed"
		if item.Completed {
			status = "Completed"
		}
		fmt.Printf("> %s (ID: %d, Status: %s)\n", item.Title, item.OriginalID, status)
	}
}