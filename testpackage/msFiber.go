package testpackage

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

// main class
func MainServer() {
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Provide a minimal config
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
	}))

	//router
	// This route path will match requests to the root route, "/":
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

	// This route path will match requests to "/about":
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("about")
	})

	// This route path will match requests to "/random.txt":
	app.Get("/random.txt", func(c *fiber.Ctx) error {
		return c.SendString("random.txt")
	})

	// Parameters
	app.Get("/user/:name/books/:title", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s\n", c.Params("name"))
		fmt.Fprintf(c, "%s\n", c.Params("title"))
		return nil
	})

	//Getter
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Get("/user/:name", func(c *fiber.Ctx) error {

		str := "hello ==> " + c.Params("name")
		return c.JSON(str)
	})

	//Setter
	app.Post("/", func(c *fiber.Ctx) error {
		p := new(Person)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		log.Println(p.Name) // john
		log.Println(p.Pass) // doe
		str := p.Name + p.Pass
		return c.JSON(str)
	})

	app.Post("/inet", func(c *fiber.Ctx) error {
		a := c.Query("search")
		str := "my search is  " + a
		return c.JSON(str)
	})

	v1.Post("/valid", func(c *fiber.Ctx) error {
		//Connect to database
		type User struct {
			Name     string `json:"name" validate:"required,min=3,max=32"`
			IsActive *bool  `json:"isactive" validate:"required"`
			Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
		}
		user := new(User)
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		validate := validator.New()
		errors := validate.Struct(user)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
		}
		return c.JSON(user)
	})

	app.Listen(":3000")
}
