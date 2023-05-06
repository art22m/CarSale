package server

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"hw5/internal/pkg/repository"
)

func (s *server) createCarSale(ctx context.Context, req *http.Request) (int64, int) {
	carSale, err := getCarSaleData(req.Body)
	if err != nil {
		log.Printf("error while reading req body: %s\n", err)
		return 0, http.StatusBadRequest
	}

	id, err := s.carSaleRepo.Create(ctx, &repository.CarSale{
		Brand:    carSale.Brand,
		Model:    carSale.Model,
		SellerID: carSale.SellerID,
	})
	if err != nil {
		log.Printf("internal error occured: %s\n", err)
		return 0, http.StatusInternalServerError
	}

	return id, http.StatusOK
}

func (s *server) updateCarSale(ctx context.Context, req *http.Request) (bool, int) {
	carSale, err := getCarSaleData(req.Body)
	if err != nil {
		log.Printf("error while reading req body: %s\n", err)
		return false, http.StatusBadRequest
	}

	ok, err := s.carSaleRepo.Update(ctx, &repository.CarSale{
		ID:       carSale.ID,
		Brand:    carSale.Brand,
		Model:    carSale.Model,
		SellerID: carSale.SellerID,
	})
	if err != nil {
		log.Printf("internal error occured: %s\n", err)
		return false, http.StatusInternalServerError
	}

	return ok, http.StatusOK
}

func (s *server) deleteCarSale(ctx context.Context, req *http.Request) (bool, int) {
	id, err := getCarSaleID(req.URL)
	if err != nil {
		log.Printf("can't parse id: %s\n", err)
		return false, http.StatusBadRequest
	}

	ok, err := s.carSaleRepo.Delete(ctx, id)
	if err != nil {
		log.Printf("internal error occured: %s\n", err)
		return false, http.StatusInternalServerError
	}

	return ok, http.StatusOK
}

func (s *server) getCarSale(ctx context.Context, req *http.Request) ([]byte, int) {
	id, err := getCarSaleID(req.URL)
	if err != nil {
		log.Printf("can't parse id: %s\n", err)
		return nil, http.StatusBadRequest
	}

	carSale, err := s.carSaleRepo.GetById(ctx, id)
	if err != nil {
		log.Printf("internal error occured: %s\n", err)
		return nil, http.StatusInternalServerError
	}

	scs := &serverCarSale{
		ID:       carSale.ID,
		Brand:    carSale.Brand,
		Model:    carSale.Model,
		SellerID: carSale.SellerID,
	}

	data, err := json.Marshal(scs)
	if err != nil {
		log.Printf("can't marshal car sale with id: %d. Error: %s\n", id, err)
		return nil, http.StatusInternalServerError
	}

	return data, http.StatusOK
}

// Helpers

func getCarSaleData(reader io.ReadCloser) (serverCarSale, error) {
	body, err := io.ReadAll(reader)
	if err != nil {
		return serverCarSale{}, err
	}

	data := serverCarSale{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("can't unmarshall data with error: %s\n", err)
		return data, err
	}

	return data, nil
}

func getCarSaleID(reqUrl *url.URL) (int64, error) {
	idStr := reqUrl.Query().Get("id")
	if len(idStr) == 0 {
		return 0, errors.New("can't get id")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, errors.New("can't parse id")
	}

	return id, nil
}
