package main

import (
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tommylay1902/todo/types"
	views "github.com/tommylay1902/todo/views/page/home"
)

func main() {
	e := echo.New()
	e.Static("/static", "static")

	todos := []types.Todo{{Id: uuid.New(), Description: "Learn HTMX", Completed: true}, {Id: uuid.New(), Description: "Learn Templ", Completed: true}}

	e.GET("/", func(c echo.Context) error {
		return views.Index(todos).Render(c.Request().Context(), c.Response().Writer)
	})

	e.POST("/create", func(c echo.Context) error {
		description := c.FormValue("description")

		newTodo := types.Todo{Description: description, Completed: false, Id: uuid.New()}
		todos = append(todos, newTodo)

		return views.Todo(&newTodo).Render(c.Request().Context(), c.Response().Writer)
	})

	e.DELETE("/delete/:id", func(c echo.Context) error {
		itemID := strings.TrimPrefix(c.Param("id"), "todo-")
		var newTodos []types.Todo
		for _, value := range todos {
			if itemID != value.Id.String() {
				newTodos = append(newTodos, value)
			}
		}

		todos = newTodos

		return views.Todo(nil).Render(c.Request().Context(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":8081"))
}
