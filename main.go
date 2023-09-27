package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-architecture/handler"
	"github.com/go-architecture/repository"
	"github.com/go-architecture/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initConfic()
	initTimeZone()
	db := initDatabase()

	costomerRepositoryDB := repository.NewCostomerRepositoryDB(db)

	costomerRepositoryMock := repository.NewCostomerRepositoryMock()
	_ = costomerRepositoryMock

	customerService := service.NewCustomerService(costomerRepositoryDB)

	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomerById).Methods(http.MethodGet)

	http.ListenAndServe(":3000", router)
}

func initConfic() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err.Error())
	}
	time.Local = location
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.hostname"),
		viper.GetString("db.port"),
		viper.GetString("db.database"))

	driver := viper.GetString("db.driver")
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		panic(err.Error())
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}