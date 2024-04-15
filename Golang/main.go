package main

import (
	"context"
	"embed"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
	"web/api"
	"web/middleware"
	"web/routes"
	"web/service"
	"web/version"
)

var (
	mHost    = flag.String("l", ":80", "listen addr")
	mVersion = flag.Bool("v", false, "version")
	mDebug   = flag.Bool("dev", false, "Enable development mode")
	mDaemon  = flag.Bool("d", false, "daemon with -d")
	mLog     = flag.String("log", "", "log file")

	//go:embed dist
	distFs embed.FS
)

func main() {
	flag.Parse()
	if *mVersion {
		fmt.Println("Version:", version.Version)
		fmt.Println("BuildTime:", version.BuildTime)
		return
	}
	if *mDebug == false {
		gin.SetMode(gin.ReleaseMode)
	}
	if *mDaemon {
		if err := exec.Command(os.Args[0], flag.Args()...).Start(); err != nil {
			log.Panicln(err)
		}
		os.Exit(0)
	}

	if *mLog != "" {
		elog := middleware.InitLog()
		elog.SetLines(1000)
		elog.SetName(*mLog)
		gin.DefaultWriter = io.MultiWriter(elog, os.Stdout)
	}

	app := gin.Default()

	mySer := service.InitService()
	myApi := api.InitApi(mySer)

	middleware.InitJWT("2D7A68D4-EA9D-4260-900C-C64BB369D709", 1440*7)
	routes.InitRouter(app, myApi)

	assets, _ := fs.Sub(distFs, "dist/assets")
	app.StaticFS("/assets", http.FS(assets))
	app.StaticFileFS("/favicon.ico", "dist/favicon.ico", http.FS(distFs))
	app.StaticFileFS("/logo.png", "dist/logo.png", http.FS(distFs))
	app.GET("/", func(c *gin.Context) {
		index, _ := distFs.ReadFile("dist/index.html")
		c.Data(http.StatusOK, "text/html", index)
	})

	srv := &http.Server{
		Addr:    *mHost,
		Handler: app,
	}
	go func() {
		log.Println("Runing .", *mHost)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicln(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTRAP, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server .")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
	log.Println("Server Exiting .")
}
