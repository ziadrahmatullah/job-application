package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/appvalidator"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/database"
	grpchandler "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/handler/grpc_handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/middleware"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/usecase"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
	db := database.ConnectDB()
	validator := appvalidator.NewAppValidatorImpl()
	appvalidator.SetValidator(validator)

	jr := repository.NewJobRepository(db)
	ju := usecase.NewJobUsecase(jr)
	jh := grpchandler.NewJobGRPCHandler(ju, validator)

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uh := grpchandler.NewAuthGRPCHandler(uu, validator)

	ar := repository.NewApplyJobRepository(db)
	au := usecase.NewApplyJobUsecase(ar, ur, jr)
	ah := grpchandler.NewApplyJobGRPCHandler(au, validator)

	list, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("error starting tcp server")
	}

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		middleware.LoggerInterceptor,
		middleware.ErrorInterceptor,
		middleware.AuthInterceptor,
	))

	pb.RegisterAuthServiceServer(server, uh)
	pb.RegisterJobServiceServer(server, jh)
	pb.RegisterApplyJobServiceServer(server, ah)

	log.Println("starting grpc server")

	signCh := make(chan os.Signal, 1)
	signal.Notify(signCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.Serve(list); err != nil {
			signCh <- syscall.SIGINT
		}
	}()
	log.Println("server started")
	<-signCh
	log.Println("server stopped")
}
