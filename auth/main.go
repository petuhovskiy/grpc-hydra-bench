//go:generate protoc --go_out=plugins=grpc:./pb  -I ../proto ../proto/auth.proto

package main

import (
	"log"
	"net"
	"net/url"
	"os"

	"github.com/petuhovskiy/grpc-hydra-bench/auth/hydracon"
	"github.com/petuhovskiy/grpc-hydra-bench/auth/impl"
	"github.com/petuhovskiy/grpc-hydra-bench/auth/libauth"
	"github.com/petuhovskiy/grpc-hydra-bench/auth/middleware"
	"github.com/petuhovskiy/grpc-hydra-bench/auth/pb"
	"github.com/petuhovskiy/grpc-hydra-bench/auth/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/ory/hydra/sdk/go/hydra/client"
	"github.com/ory/hydra/sdk/go/hydra/client/admin"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

var schema = users.Schema

func connectDatabase() *sqlx.DB {
	db, err := sqlx.Connect(
		"pgx",
		os.ExpandEnv(
			"host=${DB_HOST} password=${DB_PASSWORD} user=postgres dbname=postgres sslmode=disable",
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(schema)
	if err != nil {
		log.Println("Error while executing migartion, it might already been executed.", err)
	}
	return db
}

func initHydraAdmin() *admin.Client {
	adminURL, err := url.Parse(os.Getenv("HYDRA_ADMIN_URL"))
	if err != nil {
		log.Fatal(err)
	}
	cli := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{Schemes: []string{adminURL.Scheme}, Host: adminURL.Host, BasePath: adminURL.Path})
	return cli.Admin
}

func main() {
	log.SetFlags(log.Llongfile)

	db := connectDatabase()
	userRepo := users.NewRepo(db)

	hydraAdmin := initHydraAdmin()

	flowHandler := hydracon.NewHandler(userRepo, hydraAdmin)
	httpserv := hydracon.NewLoginAndConsentServer(flowHandler)

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
			(&middleware.Auth{
				Validate: libauth.Validator(hydraAdmin),
			}).Interceptor,
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

	go func() {
		serv := httpserv
		log.Println("HTTP server started on ", serv.Addr)
		defer log.Fatal("HTTP server exited on ", serv.Addr)

		serv.ListenAndServe()
	}()

	log.Println("gRPC server started")
	defer log.Println("gRPC server exited")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
