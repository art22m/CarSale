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
	CreateCarSaleCommandSignature = "create_car_sale"
	ReadCarSaleCommandSignature   = "read_car_sale"
	UpdateCarSaleCommandSignature = "update_car_sale"
	GetCarSaleCommandSignature    = "get_car_sale"
	DeleteCarSaleCommandSignature = "delete_car_sale"
)

// Create

type createCarSaleCommand struct {
	*baseCommand
}

func NewCreateCarSaleCommand(baseCommand *baseCommand) *createCarSaleCommand {
	return &createCarSaleCommand{baseCommand: baseCommand}
}

func (c *createCarSaleCommand) Execute(args ...string) error {
	carSale, err := scanCarSale(c.scanner)
	if err != nil {
		return err
	}

	id, err := c.carSalesRepo.Create(c.ctx, carSale)
	if err != nil {
		return err
	}

	log.Printf("Car sale with id = %v created\n", id)
	return nil
}

// Read

type readCarSaleCommand struct {
	*baseCommand
}

func NewReadCarSaleCommand(baseCommand *baseCommand) *readCarSaleCommand {
	return &readCarSaleCommand{baseCommand: baseCommand}
}

func (c *readCarSaleCommand) Execute(args ...string) error {
	carSales, err := c.carSalesRepo.Read(c.ctx)
	if err != nil {
		return err
	}

	jsonCarSales, err := json.Marshal(carSales)
	if err != nil {
		return err
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, jsonCarSales, "", "    "); err != nil {
		return err
	}

	log.Println(prettyJSON.String())
	return nil
}

// Update

type updateCarSalesCommand struct {
	*baseCommand
}

func NewUpdateCarSalesCommand(baseCommand *baseCommand) *updateCarSalesCommand {
	return &updateCarSalesCommand{baseCommand: baseCommand}
}

func (c *updateCarSalesCommand) Execute(args ...string) error {
	fmt.Print("Enter car_sale id: ")
	c.scanner.Scan()
	carSaleID, err := strconv.ParseInt(c.scanner.Text(), 10, 64)
	if err != nil {
		return err
	}

	carSale, err := scanCarSale(c.scanner)
	if err != nil {
		return err
	}
	carSale.ID = carSaleID

	ok, err := c.carSalesRepo.Update(c.ctx, carSale)
	if err != nil {
		return err
	}

	if !ok {
		log.Println("No such seller found")
		return nil
	}

	log.Println("Car sale updated")
	return nil
}

// Get

type getCarSalesCommand struct {
	*baseCommand
}

func NewGetCarSalesCommand(baseCommand *baseCommand) *getCarSalesCommand {
	return &getCarSalesCommand{baseCommand: baseCommand}
}

func (c *getCarSalesCommand) Execute(args ...string) error {
	fmt.Print("Enter car_sale id: ")
	c.scanner.Scan()
	carSaleID, err := strconv.ParseInt(c.scanner.Text(), 10, 64)
	if err != nil {
		return err
	}

	carSale, err := c.carSalesRepo.GetById(c.ctx, carSaleID)
	if err != nil {
		return err
	}

	jsonCarSale, err := json.Marshal(carSale)
	if err != nil {
		return err
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, jsonCarSale, "", "    "); err != nil {
		return err
	}

	log.Println(prettyJSON.String())
	return nil
}

// Delete

type deleteCarSalesCommand struct {
	*baseCommand
}

func NewDeleteCarSalesCommand(baseCommand *baseCommand) *deleteCarSalesCommand {
	return &deleteCarSalesCommand{baseCommand: baseCommand}
}

func (c *deleteCarSalesCommand) Execute(args ...string) error {
	fmt.Print("Enter car_sale id: ")
	c.scanner.Scan()
	carSaleID, err := strconv.ParseInt(c.scanner.Text(), 10, 64)
	if err != nil {
		return err
	}

	ok, err := c.carSalesRepo.Delete(c.ctx, carSaleID)
	if err != nil {
		return err
	}

	if !ok {
		log.Println("No such seller found")
		return nil
	}

	log.Println("Car sale deleted")
	return nil
}

// Helpers

func scanCarSale(scanner *bufio.Scanner) (*repository.CarSale, error) {
	fmt.Print("Enter car brand: ")
	scanner.Scan()
	brand := scanner.Text()

	fmt.Print("Enter car model: ")
	scanner.Scan()
	model := scanner.Text()

	fmt.Print("Enter seller id: ")
	scanner.Scan()
	sellerID, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		return nil, err
	}

	return &repository.CarSale{Brand: brand, Model: model, SellerID: sellerID}, nil
}
