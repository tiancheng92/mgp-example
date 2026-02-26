package server

import (
	"context"
	"mgp_example/config"
	"mgp_example/internal/controller"
	"mgp_example/pkg/log"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	serverConfig := config.GetConf().Server
	switch serverConfig.Mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	s := &http.Server{
		Addr:    serverConfig.Host,       // 监听地址
		Handler: controller.InitRouter(), // 处理器
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	defer close(quit)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
