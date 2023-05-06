package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"hw5/internal/app/grpcserver"
	"hw5/internal/app/pb"
	"hw5/internal/pkg/console/commands_manager"
	"hw5/internal/pkg/db"
	"hw5/internal/pkg/repository"
	carSale "hw5/internal/pkg/repository/postgresql/car_sale"
	"hw5/internal/pkg/repository/postgresql/seller"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tp, err := tracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	otel.SetTracerProvider(tp)
	log.Println("Jaeger started")

	// Database
	database, err := db.CreateDatabase(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer database.GetPool().Close()

	sellersRepo := seller.NewSellers(database)
	carSalesRepo := carSale.NewCarSales(database)

	log.Println("Successfully connected to DB")

	// Prometheus
	go func() {
		log.Println("Prometheus started")
		http.ListenAndServe(":9091", promhttp.Handler())
	}()

	// gRPC
	lsn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterBuyCarServiceServer(server, grpcserver.NewImplementation(sellersRepo, carSalesRepo))

	log.Println("gRPC started")
	if err = server.Serve(lsn); err != nil {
		log.Fatal(err)
	}

	//handleConsoleCommands(ctx, sellersRepo, carSalesRepo)
}

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	const (
		service     = "api"
		environment = "development"
	)

	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
			attribute.String("environment", environment),
		)),
	)
	return tp, nil
}

func handleConsoleCommands(ctx context.Context, sellersRepo repository.SellersRepo, carSalesRepo repository.CarSalesRepo) {
	scanner := bufio.NewScanner(os.Stdin)
	commandManager := commands_manager.NewCommandManager(ctx, sellersRepo, carSalesRepo, scanner)

	for {
		fmt.Print("> ")

		scanner.Scan()
		cmd := scanner.Text()

		commandManager.Handle(cmd)
	}
}
