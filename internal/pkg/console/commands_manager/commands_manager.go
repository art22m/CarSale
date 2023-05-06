package commands_manager

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"strings"

	"hw5/internal/pkg/console/commands"
	"hw5/internal/pkg/repository"
)

type command interface {
	Execute(args ...string) error
}

type commandManager struct {
	commands map[string]command
}

func NewCommandManager(ctx context.Context, sellersRepo repository.SellersRepo, carSalesRepo repository.CarSalesRepo, scanner *bufio.Scanner) *commandManager {
	// All commands
	availableCommands := map[string]command{}

	// Base command
	baseCommand := commands.NewBaseCommand(ctx, sellersRepo, carSalesRepo, scanner)

	// Seller commands
	availableCommands[commands.CreateSellerCommandSignature] = commands.NewCreateSellerCommand(baseCommand)
	availableCommands[commands.ReadSellerCommandSignature] = commands.NewReadSellersCommand(baseCommand)
	availableCommands[commands.UpdateSellerCommandSignature] = commands.NewUpdateSellerCommand(baseCommand)
	availableCommands[commands.GetSellerCommandSignature] = commands.NewGetSellerCommand(baseCommand)
	availableCommands[commands.DeleteSellerCommandSignature] = commands.NewDeleteSellerCommand(baseCommand)

	// Car sale commands
	availableCommands[commands.CreateCarSaleCommandSignature] = commands.NewCreateCarSaleCommand(baseCommand)
	availableCommands[commands.ReadCarSaleCommandSignature] = commands.NewReadCarSaleCommand(baseCommand)
	availableCommands[commands.UpdateCarSaleCommandSignature] = commands.NewUpdateCarSalesCommand(baseCommand)
	availableCommands[commands.GetCarSaleCommandSignature] = commands.NewGetCarSalesCommand(baseCommand)
	availableCommands[commands.DeleteCarSaleCommandSignature] = commands.NewDeleteCarSalesCommand(baseCommand)

	// Spell command
	availableCommands[commands.SpellCommandSignature] = commands.NewSpellCommand(baseCommand)

	// Gofmt command
	availableCommands[commands.GofmtCommandSignature] = commands.NewGofmtCommand(baseCommand)

	// Help command
	availableCommands[commands.HelpCommandSignature] = commands.NewHelpCommand(baseCommand)

	return &commandManager{commands: availableCommands}
}

func (cm commandManager) Handle(input string) {
	inputSplit := strings.Fields(input)
	if len(inputSplit) == 0 {
		log.Println("Empty command")
		return
	}

	cmdName, args := inputSplit[0], inputSplit[1:]
	cmd, ok := cm.commands[cmdName]
	if !ok {
		fmt.Printf("No such command. Write %s to see all available commands\n", commands.HelpCommandSignature)
		return
	}

	if err := cmd.Execute(args...); err != nil {
		log.Println(err)
		return
	}
}
