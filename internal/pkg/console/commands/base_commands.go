package commands

import (
	"bufio"
	"context"
	"fmt"

	"hw5/internal/pkg/repository"
)

// Base

type baseCommand struct {
	ctx          context.Context
	sellersRepo  repository.SellersRepo
	carSalesRepo repository.CarSalesRepo
	scanner      *bufio.Scanner
}

func NewBaseCommand(ctx context.Context, sellersRepo repository.SellersRepo, carSalesRepo repository.CarSalesRepo, scanner *bufio.Scanner) *baseCommand {
	return &baseCommand{
		ctx:          ctx,
		sellersRepo:  sellersRepo,
		carSalesRepo: carSalesRepo,
		scanner:      scanner,
	}
}

func (c *baseCommand) Execute(args ...string) error { return nil }

// Help

const HelpCommandSignature = "help"

type helpCommand struct {
	*baseCommand
}

func NewHelpCommand(baseCommand *baseCommand) *helpCommand {
	return &helpCommand{baseCommand: baseCommand}
}

func (c *helpCommand) Execute(args ...string) error {
	fmt.Println("-----------------HELP-----------------")

	c.printSellerCommandsDescription()
	c.printCarSaleCommandsDescription()
	c.printSpellCommandDescription()
	c.printGofmtCommandDescription()

	fmt.Println("Helper command:")
	printCommandHeader(HelpCommandSignature, "print help")

	fmt.Println("--------------------------------------")

	return nil
}

func (c *helpCommand) printSellerCommandsDescription() {
	fmt.Println("Seller commands:")
	printCommandHeader(CreateSellerCommandSignature, "create new seller")
	printCommandHeader(ReadSellerCommandSignature, "print all sellers")
	printCommandHeader(UpdateSellerCommandSignature, "update specified seller")
	printCommandHeader(GetSellerCommandSignature, "get seller by id")
	printCommandHeader(DeleteSellerCommandSignature, "delete seller by id")
}

func (c *helpCommand) printCarSaleCommandsDescription() {
	fmt.Println("Car sale commands:")
	printCommandHeader(CreateCarSaleCommandSignature, "create new car sale")
	printCommandHeader(ReadCarSaleCommandSignature, "print all car sales")
	printCommandHeader(UpdateCarSaleCommandSignature, "update specified car sale")
	printCommandHeader(GetCarSaleCommandSignature, "get car sale by id")
	printCommandHeader(DeleteCarSaleCommandSignature, "delete car sale by id")
}

func (c *helpCommand) printSpellCommandDescription() {
	fmt.Println("Spell command:")
	printCommandHeader(fmt.Sprintf("%s <name>", SpellCommandSignature), "prints all the letters of word separated by a space")
	printCommandHeader("", fmt.Sprintf("%s ozon -> \"returns o z o n\"", SpellCommandSignature))
}

func (c *helpCommand) printGofmtCommandDescription() {
	fmt.Println("Gofmt command:")
	printCommandHeader(fmt.Sprintf("%s <path>", GofmtCommandSignature), "simplified gofmt. Accepts only *.txt files.")
}

func printCommandHeader(command, description string) {
	fmt.Printf("\t %-30s --\t %s\n", command, description)
}
