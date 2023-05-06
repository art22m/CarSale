package commands

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"hw5/internal/pkg/repository"
)

// Signatures

const (
	CreateSellerCommandSignature = "create_seller"
	ReadSellerCommandSignature   = "read_seller"
	UpdateSellerCommandSignature = "update_seller"
	GetSellerCommandSignature    = "get_seller"
	DeleteSellerCommandSignature = "delete_seller"
)

// Create

type createSellerCommand struct {
	*baseCommand
}

func NewCreateSellerCommand(baseCommand *baseCommand) *createSellerCommand {
	return &createSellerCommand{baseCommand: baseCommand}
}

func (c *createSellerCommand) Execute(args ...string) error {
	seller := scanSeller(c.scanner)

	id, err := c.sellersRepo.Create(c.ctx, seller)
	if err != nil {
		return err
	}

	log.Printf("Seller with id = %v created\n", id)
	return nil
}

// Read

type readSellersCommand struct {
	*baseCommand
}

func NewReadSellersCommand(baseCommand *baseCommand) *readSellersCommand {
	return &readSellersCommand{baseCommand: baseCommand}
}

func (c *readSellersCommand) Execute(args ...string) error {
	sellers, err := c.sellersRepo.Read(c.ctx)
	if err != nil {
		return err
	}

	jsonSellers, err := json.Marshal(sellers)
	if err != nil {
		return err
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, jsonSellers, "", "    "); err != nil {
		return err
	}

	log.Println(prettyJSON.String())
	return nil
}

// Update

type updateSellerCommand struct {
	*baseCommand
}

func NewUpdateSellerCommand(baseCommand *baseCommand) *updateSellerCommand {
	return &updateSellerCommand{baseCommand: baseCommand}
}

func (c *updateSellerCommand) Execute(args ...string) error {
	fmt.Print("Enter seller id: ")
	c.scanner.Scan()
	sellerID, err := strconv.ParseInt(c.scanner.Text(), 10, 64)
	if err != nil {
		return err
	}

	seller := scanSeller(c.scanner)
	seller.ID = sellerID

	ok, err := c.sellersRepo.Update(c.ctx, seller)
	if err != nil {
		return err
	}

	if !ok {
		log.Println("No such car sale found")
		return nil
	}

	log.Println("Seller updated")
	return nil
}

// Get

type getSellerCommand struct {
	*baseCommand
}

func NewGetSellerCommand(baseCommand *baseCommand) *getSellerCommand {
	return &getSellerCommand{baseCommand: baseCommand}
}

func (c *getSellerCommand) Execute(args ...string) error {
	fmt.Print("Enter seller id: ")
	c.scanner.Scan()
	sellerID, err := strconv.ParseInt(c.scanner.Text(), 10, 64)
	if err != nil {
		return err
	}

	seller, err := c.sellersRepo.GetById(c.ctx, sellerID)
	if err != nil {
		return err
	}

	jsonSeller, err := json.Marshal(seller)
	if err != nil {
		return err
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, jsonSeller, "", "    "); err != nil {
		return err
	}

	log.Println(prettyJSON.String())
	return nil
}

// Delete

type deleteSellerCommand struct {
	*baseCommand
}

func NewDeleteSellerCommand(baseCommand *baseCommand) *deleteSellerCommand {
	return &deleteSellerCommand{baseCommand: baseCommand}
}

func (c *deleteSellerCommand) Execute(args ...string) error {
	fmt.Print("Enter seller id: ")
	c.scanner.Scan()
	sellerID, err := strconv.ParseInt(c.scanner.Text(), 10, 64)
	if err != nil {
		return err
	}

	ok, err := c.sellersRepo.Delete(c.ctx, sellerID)
	if err != nil {
		return err
	}

	if !ok {
		log.Println("No such seller found")
		return nil
	}

	log.Println("Seller deleted")
	return nil
}

// Helpers

func scanSeller(scanner *bufio.Scanner) *repository.Seller {
	fmt.Print("Enter seller name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter seller phone: ")
	scanner.Scan()
	description := scanner.Text()

	return &repository.Seller{Name: name, PhoneNumber: description}
}
