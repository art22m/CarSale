package server

import (
	"context"
	"fmt"
	"hw5/internal/pkg/repository"
	"net/http"
)

type serverSeller struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type serverCarSale struct {
	ID       int64  `json:"id"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	SellerID int64  `json:"seller_id"`
}

type server struct {
	sellersRepo repository.SellersRepo
	carSaleRepo repository.CarSalesRepo
}

func CreateServer(ctx context.Context, sr repository.SellersRepo, csr repository.CarSalesRepo) *http.ServeMux {
	serv := server{
		sellersRepo: sr,
		carSaleRepo: csr,
	}

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/seller", func(res http.ResponseWriter, req *http.Request) {
		if req == nil {
			return
		}

		switch req.Method {
		case http.MethodGet:
			data, status := serv.getSeller(ctx, req)
			res.WriteHeader(status)
			res.Write(data)

		case http.MethodPost:
			_, status := serv.createSeller(ctx, req)
			res.WriteHeader(status)

		case http.MethodPut:
			_, status := serv.updateSeller(ctx, req)
			res.WriteHeader(status)

		case http.MethodDelete:
			_, status := serv.deleteSeller(ctx, req)
			res.WriteHeader(status)

		default:
			fmt.Printf("unsupported method: [%s]", req.Method)
			res.WriteHeader(http.StatusNotImplemented)
		}
	})

	serveMux.HandleFunc("/car_sale", func(res http.ResponseWriter, req *http.Request) {
		if req == nil {
			return
		}

		switch req.Method {
		case http.MethodGet:
			data, status := serv.getCarSale(ctx, req)
			res.WriteHeader(status)
			res.Write(data)

		case http.MethodPost:
			_, status := serv.createCarSale(ctx, req)
			res.WriteHeader(status)

		case http.MethodPut:
			_, status := serv.updateCarSale(ctx, req)
			res.WriteHeader(status)

		case http.MethodDelete:
			_, status := serv.deleteCarSale(ctx, req)
			res.WriteHeader(status)

		default:
			fmt.Printf("unsupported method: [%s]", req.Method)
			res.WriteHeader(http.StatusNotImplemented)
		}
	})

	return serveMux
}
