// the main module
package main

import (
    // import packages
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    // import modules
    "go-fiber-todos/database"
    "go-fiber-todos/todos"
)

// App config => App denotes the Fiber application.
func setupV1(app *fiber.App) {
    // Group is used for Routes with common prefix to define a new sub-router with optional middleware.
    v1 := app.Group("/v1")
    //Each route will have /v1 prefix
    setupTodosRoutes(v1)
}

// Router defines all router handle interface includes app and group router
func setupTodosRoutes(grp fiber.Router) {
    // Group is used for Routes with common prefix => Each route will have /todos prefix
    todosRoutes := grp.Group("/todos")
    // Route for Get all todos -> navigate to => http://127.0.0.1:3000/v1/todos/
    todosRoutes.Get("/", todos.GetAll)
    // Route for Get a todo -> navigate to => http://127.0.0.1:3000/v1/todos/<todo's id>
    todosRoutes.Get("/:id", todos.GetOne)
    // Route for Add a todo -> navigate to => http://127.0.0.1:3000/v1/todos/
    todosRoutes.Post("/", todos.AddTodo)
    // Route for Delete a todo -> navigate to => http://127.0.0.1:3000/v1/todos/<todo's id>
    todosRoutes.Delete("/:id", todos.DeleteTodo)
    // Route for Update a todo -> navigate to => http://127.0.0.1:3000/v1/todos/<todo's id>
    todosRoutes.Patch("/:id", todos.UpdateTodo)
}

// Database Connect function
func initDatabase() {
    // define error here to prevent overshadowing the global DB
    var err error
    // Create todos sqlite file & Config GORM config
    // GORM performs single create, update, delete operations in transactions by default to ensure database data integrity
    database.DBConn, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})

    // Connect to database
    if err != nil {
        // Database was connected
        panic("failed to connect database")
    }

    fmt.Println("Database successfully connected")

    // AutoMigrate run auto migration for gorm model
    database.DBConn.AutoMigrate(&todos.Todo{})
    // Initialize Database connection
    fmt.Println("Database Migrated")
}

// entry point to our program
func main() {
    // call the New() method - used to instantiate a new Fiber App
    app := fiber.New()

    // call the initDatabase() method
    initDatabase()
    // call the setupV1(app) method
    setupV1(app)

    // Simple route => Middleware function
    app.Get("/", func(c *fiber.Ctx) error {
        // Returns plain text.
        return c.SendString("Hello, World!")
        // navigate to => http://127.0.0.1:3000
    })

    // sets up logger
    // Use middlewares for each route
    // This method will match all HTTP verbs: GET, POST, PUT etc Then create a log when every HTTP verb get invoked
    app.Use(logger.New(logger.Config{ // add Logger middleware with config
        Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
    }))

    // listen/Serve the new Fiber app on port 3000
    err := app.Listen(":3000")

    // handle panic errors => panic built-in function that stops the execution of a function and immediately normal execution of that function with an error
    if err != nil {
        panic(err)
    }
}
