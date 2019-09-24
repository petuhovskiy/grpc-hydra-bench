//go:generate protoc --go_out=plugins=grpc:./pb  -I ../proto ../proto/auth.proto

package main

import (
	"log"
	"net"
	"url"
	"os"

	"github.com/petuhovskiy/grpc-hydra-bench/auth/impl"
	"github.com/petuhovskiy/grpc-hydra-bench/auth/pb"
	"github.com/petuhovskiy/grpc-hydra-bench/auth/users"
	"github.com/petuhovskiy/grpc-hydra-bench/auth/libauth"
	"github.com/petuhovskiy/grpc-hydra-bench/auth/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/petuhovskiy/grpc-hydra-bench/hydracli/client/admin"
	"github.com/petuhovskiy/grpc-hydra-bench/hydracli/client"
	"github.com/petuhovskiy/grpc-hydra-bench/hydracli"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

)

var schema = users.Schema

func connectDatabase() *sqlx.DB {
	db, err := sqlx.Connect(
		"postgres",
		os.ExpandEnv(
			"host=${DB_HOST} password=${DB_PASSWORD} user=postgres dbname=postgres sslmode=disable",
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	db.MustExec(schema)
	return db
}

func initHydraAdmin() *admin.Client {
	adminURL := url.Parse(os.Getenv("HYDRA_ADMIN_URL"))
	cli := hydra.NewHTTPClientWithConfig(nil, &client.TransportConfig{Schemes: []string{adminURL.Scheme}, Host: adminURL.Host, BasePath: adminURL.Path})
	return cli.Admin
}

func main() {
	db := connectDatabase()
	userRepo := users.NewRepo(db)

	hydraAdmin := initHydraAdmin()

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
			middleware.Auth{
				Validate: libauth.Validator(hydraAdmin),
			}.Interceptor,
		)),
	}

	server := grpc.NewServer(opts...)
	pb.RegisterAuthServer(server, impl.NewAuthServer(
		userRepo,
	))
	reflection.Register(server)

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC server started")
	defer log.Println("gRPC server exited")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
