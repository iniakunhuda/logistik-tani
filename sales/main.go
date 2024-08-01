package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/sales/controller"
	"github.com/iniakunhuda/logistik-tani/sales/repository"
	"github.com/iniakunhuda/logistik-tani/sales/router"
	"github.com/iniakunhuda/logistik-tani/sales/service"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// // Define command-line flags
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4001, "HTTP server network port")
	dsn := flag.String("dsn", "root:@tcp(localhost:3306)/crmtani_inventory?parseTime=true", "MySQL data source name")
	flag.Parse()

	// // Create logger for writing information and error messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// // Open database connection
	db, err := openDB(*dsn)
	if err != nil {
		errLog.Fatal(err)
	}

	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)

	validate := validator.New()
	produkRepository := repository.NewSalesRepositoryImpl(db)
	produkService := service.NewInventoryServiceImpl(produkRepository, validate)
	produkController := controller.NewSalesController(produkService)

	routes := router.NewRouter(produkController)
	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errLog,
		Handler:      routes,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", serverURI)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}

func openDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
