package main

import (
	"fmt"
	"time"
)

// TodoItem represents a single todo item
type TodoItem struct {
	ID		  int
	Title     string
	Completed bool
	CreatedAt time.Time
}

// TodoList represents a list of todo items
type TodoList struct {
	Items []*TodoItem
}

// AddItem adds a new todo item to the list
func (tl *TodoList) AddItem(title string) {
	item := &TodoItem{
		ID:        len(tl.Items) + 1,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	tl.Items = append(tl.Items, item)
}

// CompleteItem marks a todo item as completed
func (tl *TodoList) CompleteItem(id int) {
	for _, item := range tl.Items {
		if item.ID == id {
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
		if item.ID == id {
			tl.Items = append(tl.Items[:i], tl.Items[i+1:]...)
			fmt.Printf("Task '%s' deleted.\n", item.Title)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", id)
}


// PrintList prints the TodoList to the console
func (tl *TodoList) PrintList() {
	fmt.Println("Todo List:")
	for _, item := range tl.Items {
		status := "Not Completed"
		if item.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s (%s)\n", item.ID, item.Title, status)
	}
}