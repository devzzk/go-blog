package main

import (
	_ "context"
	"log"
	"net/http"
	"os"
	_ "os/signal"
	_ "syscall"
	_ "time"
)

func main() {
	log.Println("Starting server...")

	//router := gin.Default() // warning: Creating an Engine instance with the Logger and Recovery middleware already attached.
	//r := gin.New()
	//r.Use(gin.Logger())
	//
	//r.GET("/", func(c *gin.Context) {
	//	time.Sleep(5 * time.Second)
	//	c.String(http.StatusOK, "Welcome Gin Server")
	//})

	ds, err := InitDS()

	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	router, err := Inject(ds)

	if err != nil {
		log.Fatalf("Failure to inject data sources: %v\n", err)
	}

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	quit := make(chan os.Signal)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}

	log.Println("Server exiting")
}
