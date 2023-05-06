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

func (s *server) createSeller(ctx context.Context, req *http.Request) (int64, int) {
	seller, err := getSellerData(req.Body)
	if err != nil {
		log.Printf("error while reading req body: %s\n", err)
		return 0, http.StatusBadRequest
	}

	id, err := s.sellersRepo.Create(ctx, &repository.Seller{
		Name:        seller.Name,
		PhoneNumber: seller.PhoneNumber,
	})
	if err != nil {
		log.Printf("internal error occured: %s\n", err)
		return 0, http.StatusInternalServerError
	}

	return id, http.StatusOK
}

func (s *server) updateSeller(ctx context.Context, req *http.Request) (bool, int) {
	seller, err := getSellerData(req.Body)
	if err != nil {
		log.Printf("error while reading req body: %s\n", err)
		return false, http.StatusBadRequest
	}

	ok, err := s.sellersRepo.Update(ctx, &repository.Seller{
		ID:          seller.ID,
		Name:        seller.Name,
		PhoneNumber: seller.PhoneNumber,
	})
	if err != nil {
		log.Printf("internal error occured: %s\n", err)
		return false, http.StatusInternalServerError
	}

	return ok, http.StatusOK
}

func (s *server) deleteSeller(ctx context.Context, req *http.Request) (bool, int) {
	id, err := getSellerID(req.URL)
	if err != nil {
		log.Printf("can't parse id: %s\n", err)
		return false, http.StatusBadRequest
	}

	ok, err := s.sellersRepo.Delete(ctx, id)
	if err != nil {
		log.Printf("internal error occured: %s\n", err)
		return false, http.StatusInternalServerError
	}

	return ok, http.StatusOK
}

func (s *server) getSeller(ctx context.Context, req *http.Request) ([]byte, int) {
	id, err := getSellerID(req.URL)
	if err != nil {
		log.Printf("can't parse id: %s\n", err)
		return nil, http.StatusBadRequest
	}

	seller, err := s.sellersRepo.GetById(ctx, id)
	if err != nil {
		log.Printf("internal error occured: %s\n", err)
		return nil, http.StatusInternalServerError
	}

	ss := &serverSeller{
		ID:          seller.ID,
		Name:        seller.Name,
		PhoneNumber: seller.PhoneNumber,
	}

	data, err := json.Marshal(ss)
	if err != nil {
		log.Printf("can't marshal seller with id: %d. Error: %s\n", id, err)
		return nil, http.StatusInternalServerError
	}

	return data, http.StatusOK
}

// Helpers

func getSellerData(reader io.ReadCloser) (serverSeller, error) {
	body, err := io.ReadAll(reader)
	if err != nil {
		return serverSeller{}, err
	}

	data := serverSeller{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("can't unmarshall data with error: %s\n", err)
		return data, err
	}

	return data, nil
}

func getSellerID(reqUrl *url.URL) (int64, error) {
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
