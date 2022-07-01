package main


import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
)


// Database connection for reading balance
var db *sql.DB


// Handler to add money to user's account
func add(ctx *fiber.Ctx) error {
	// send to kafka {name: name, amount: amount, command: command}
	msg := fmt.Sprintf("name is %s, operation is add, amount is %s", ctx.Params("name"), ctx.Params("amount"))
	return ctx.SendString(msg)
}


// Handler to withdraw money from user's account
func withdraw(ctx *fiber.Ctx) error {
	// send to kafka {name: name, amount: amount, command: command}
	msg := fmt.Sprintf("name is %s, operation is withdraw, amount is %s", ctx.Params("name"), ctx.Params("amount"))
	return ctx.SendString(msg)
}


// Handler to get balance of user
func getBalance(ctx *fiber.Ctx) error {
	amount, _ := GetAmount(db, ctx.Params("name"))
	msg := fmt.Sprintf("name is %s, operation is get_balance,\nBalance is %d", ctx.Params("name"), amount)
	return ctx.SendString(msg)
}


func main() {
	db = GetDb()
	defer db.Close()
	app := fiber.New()
	app.Post("/:name/add/:amount", add)
	app.Post("/:name/withdraw/:amount", withdraw)
	app.Get("/:name/balance", getBalance)
	app.Listen(":3000")
}
