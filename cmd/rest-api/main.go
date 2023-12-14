package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/database"
	restapihandler "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/handler/rest_api_handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/usecase"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
	db := database.ConnectDB()

	addr := os.Getenv("APP_PORT")
	jr := repository.NewJobRepository(db)
	ju := usecase.NewJobUsecase(jr)
	jh := restapihandler.NewJobHandler(ju)

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uh := restapihandler.NewUserHandler(uu)

	ar := repository.NewApplyJobRepository(db)
	au := usecase.NewApplyJobUsecase(ar, ur, jr)
	ah := restapihandler.NewApplyJobHandler(au)

	opts := server.RouterOpts{
		JobHandler:      jh,
		UserHandler:     uh,
		ApplyJobHandler: ah,
	}
	r := server.NewRouter(opts)

	srv := http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
