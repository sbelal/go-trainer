package main

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

var Todos []Todo
var nextID = 1

func AddTodoItem(title string) Todo {
	newTodo := Todo{ID: nextID, Title: title, Completed: false}
	nextID = nextID + 1
	Todos = append(Todos, newTodo)
	return newTodo
}

func ListTodos() []Todo {
	return Todos
}

func GetTodo(id int) (*Todo, int) {
	for index, todo := range Todos {
		if todo.ID == id {
			return &Todos[index], index
		}
	}

	return nil, -1
}

func ToggleTodoComplete(id int) bool {
	foundTodo, _ := GetTodo(id)
	if foundTodo == nil {
		return false
	}

	foundTodo.Completed = !foundTodo.Completed
	return true
}

func DeleteTodoItem(id int) bool {
	foundTodo, foundIndex := GetTodo(id)
	if foundTodo == nil {
		return false
	}

	Todos = append(Todos[:foundIndex], Todos[foundIndex+1:]...)

	return true
}
