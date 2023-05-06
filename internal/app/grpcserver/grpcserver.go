package grpcserver

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"hw5/internal/app"
	"hw5/internal/app/pb"
	"hw5/internal/pkg/repository"
)

type Implementation struct {
	pb.UnimplementedBuyCarServiceServer

	sellersRepo  repository.SellersRepo
	carSalesRepo repository.CarSalesRepo
}

func NewImplementation(sellersRepo repository.SellersRepo, carSalesRepo repository.CarSalesRepo) *Implementation {
	return &Implementation{
		sellersRepo:  sellersRepo,
		carSalesRepo: carSalesRepo,
	}
}

/* CarSale */

func (s *Implementation) CreateCarSale(ctx context.Context, in *pb.CreateCarSaleRequest) (*pb.CreateCarSaleResponse, error) {
	id, err := s.carSalesRepo.Create(ctx, &repository.CarSale{
		Brand:    in.Brand,
		Model:    in.Model,
		SellerID: in.SellerID,
	})

	if err != nil {
		return nil, err
	}

	// Prometheus stats
	app.NewCarSalesCounter.Inc()

	return &pb.CreateCarSaleResponse{Id: id}, nil
}

func (s *Implementation) UpdateCarSale(ctx context.Context, in *pb.UpdateCarSaleRequest) (*pb.UpdateCarSaleResponse, error) {
	ok, err := s.carSalesRepo.Update(ctx, &repository.CarSale{
		ID:       in.CarSale.Id,
		Brand:    in.CarSale.Brand,
		Model:    in.CarSale.Model,
		SellerID: in.CarSale.SellerID,
	})

	return &pb.UpdateCarSaleResponse{Ok: ok}, err
}

func (s *Implementation) DeleteCarSale(ctx context.Context, in *pb.DeleteCarSaleRequest) (*pb.DeleteCarSaleResponse, error) {
	ok, err := s.carSalesRepo.Delete(ctx, in.Id)
	return &pb.DeleteCarSaleResponse{Ok: ok}, err
}

func (s *Implementation) GetCarSale(ctx context.Context, in *pb.GetCarSaleRequest) (*pb.GetCarSaleResponse, error) {
	carSale, err := s.carSalesRepo.GetById(ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &pb.GetCarSaleResponse{CarSale: &pb.CarSale{
		Id:       carSale.ID,
		Brand:    carSale.Brand,
		Model:    carSale.Model,
		SellerID: carSale.SellerID,
	}}, nil
}

/* Seller */

func (s *Implementation) CreateSeller(ctx context.Context, in *pb.CreateSellerRequest) (*pb.CreateSellerResponse, error) {
	tr := otel.Tracer("CreateSeller")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String(in.String()))
	defer span.End()

	id, err := s.sellersRepo.Create(ctx, &repository.Seller{
		Name:        in.Name,
		PhoneNumber: in.PhoneNumber,
	})

	if err != nil {
		return nil, err
	}

	// Prometheus stats
	app.NewSellersCounter.Inc()

	return &pb.CreateSellerResponse{Id: id}, nil
}

func (s *Implementation) UpdateSeller(ctx context.Context, in *pb.UpdateSellerRequest) (*pb.UpdateSellerResponse, error) {
	ok, err := s.sellersRepo.Update(ctx, &repository.Seller{
		ID:          in.Seller.Id,
		Name:        in.Seller.Name,
		PhoneNumber: in.Seller.PhoneNumber,
	})

	return &pb.UpdateSellerResponse{Ok: ok}, err
}

func (s *Implementation) DeleteSeller(ctx context.Context, in *pb.DeleteSellerRequest) (*pb.DeleteSellerResponse, error) {
	ok, err := s.sellersRepo.Delete(ctx, in.Id)
	return &pb.DeleteSellerResponse{Ok: ok}, err
}

func (s *Implementation) GetSeller(ctx context.Context, in *pb.GetSellerRequest) (*pb.GetSellerResponse, error) {
	seller, err := s.sellersRepo.GetById(ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &pb.GetSellerResponse{Seller: &pb.Seller{
		Id:          seller.ID,
		Name:        seller.Name,
		PhoneNumber: seller.PhoneNumber,
	}}, nil
}
