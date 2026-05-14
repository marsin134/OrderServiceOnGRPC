package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"orderServiceGRPC/cmd/app"
	"orderServiceGRPC/internal/config"
	pb "orderServiceGRPC/pkg/generated/order"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := config.LoadEnvFile(".env"); err != nil {
		config.Log.WithFields(logrus.Fields{
			"func":  "main",
			"error": err}).Error("Failed to load .env")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		config.Log.WithFields(logrus.Fields{
			"func":  "main",
			"error": err}).Error("Failed to read .env")
	}

	handler, db := app.App(cfg)
	defer db.Close()

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recoveryInterceptor(),
		),
	)

	pb.RegisterOrderServiceServer(grpcServer, handler)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		config.Log.Fatal("Failed to listen", zap.Error(err))
	}

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		config.Log.Info("Shutting down server...")
		grpcServer.GracefulStop()
		config.Log.Info("Server stopped")
	}()

	config.Log.Info("gRPC server listening on :50051")
	if err = grpcServer.Serve(lis); err != nil {
		config.Log.Fatal("Failed to serve", zap.Error(err))
	}
}

func recoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				config.Log.Error("Recovered from panic",
					zap.String("method", info.FullMethod),
					zap.Any("panic", r),
				)
				err = status.Error(codes.Internal, "internal server error")
			}
		}()

		return handler(ctx, req)
	}
}
