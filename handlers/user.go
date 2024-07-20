package handlers

import "github.com/gofiber/fiber/v2"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := User{
		ID:    id,
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	user.ID = "12345" // This would usually be generated
	return c.Status(fiber.StatusCreated).JSON(user)
}
