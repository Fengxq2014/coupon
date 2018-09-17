package main

import (
	"context"
	"github.com/Fengxq2014/coupon/common/log"
	"github.com/Fengxq2014/coupon/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Init(gin.Mode())
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Init()
	r := initRouter()
	server := &http.Server{
		Addr:    os.Getenv("SERVER_PORT"),
		Handler: r,
	}

	go func() {
		log.Infof("server listen on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("listen server error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	log.Infof("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//grpc.CloseClient()
	if err := server.Shutdown(ctx); err != nil {
		log.Errorf("server Shutdown err:%v", err)
	}
	log.Infof("server exited")
}
