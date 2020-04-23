package rpc

//this is your entry point server, it will not be generated again.
import (
	"fmt"
	"github.com/benka-me/laruche-server/go-pkg/config"
	"github.com/benka-me/laruche-server/go-pkg/db"
	"github.com/benka-me/laruche-server/go-pkg/larsrv"
	"github.com/benka-me/laruche/go-pkg/discover"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// This structure will be passed to your handlers. Add everything you need inside.
type App struct {
	Clients
	Config *config.Config
	DB     *db.DB
}

var grpcServer *grpc.Server

func Server_2_0(engine discover.Engine) {
	var err error
	port, err := engine.ThisPort("benka-me/laruche-server")
	if err != nil {
		log.Fatal(err)
	}
	app := &App{
		Clients: InitClients(engine, grpc.WithInsecure()), // Init clients of dependencies services
		Config:  config.Init(engine.Dev),
	}
	app.DB = db.Init(app.Config)

	grpcServer = grpc.NewServer()
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	{
		larsrv.RegisterLarsrvServer(grpcServer, app) // Register your service server.
		reflection.Register(grpcServer)
	}

	fmt.Println("benka-me/laruche-server service running on port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
