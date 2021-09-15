package subs

// milestones
/*

 */
import (
	"fmt"

	"github.com/alaref-codes/subs/database"
	"github.com/gofiber/fiber/v2"
)

type Su struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func GetAllSubs(c *fiber.Ctx) error {
	db := database.DBConn
	var sub []Su
	db.Find(&sub)
	return c.JSON(sub)
}

func GetOneSub(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var sub Su
	db.Find(&sub, id)
	return c.JSON(sub)
}

func CreateSub(c *fiber.Ctx) error {
	db := database.DBConn
	var sub Su
	err := c.BodyParser(&sub)

	if err != nil {
		return err
	}
	result := db.Where("email = ?", sub.Email).Find(&sub)
	if result.RowsAffected != 0 {
		return fiber.NewError(503, "Record already exists")
	}
	db.Create(&sub)
	return c.JSON(sub)
}

func DeleteSub(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var sub Su

	// db.Where("id = ?", id).Delete(&sub)  This one works too

	result := db.First(&sub, id)

	if result.RowsAffected == 0 {
		return fiber.NewError(503, "Record doesn't exists")
	}

	result.Delete(&sub)

	return c.SendString("email Deleted successfully")
}

func UpdateSub(c *fiber.Ctx) error {
	db := database.DBConn
	var sub Su
	var newSub Su
	err := c.BodyParser(&newSub)

	if err != nil {
		return err
	}

	result := db.First(&sub, newSub.Id)
	fmt.Println(result.RowsAffected)

	if result.RowsAffected == 0 {
		return fiber.NewError(503, "Record doesn't exists")
	}

	sub.Email = newSub.Email
	db.Save(&sub)

	return c.JSON(sub)
}
