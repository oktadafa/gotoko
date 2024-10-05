package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oktadafa/gotoko/app/models"
	"github.com/oktadafa/gotoko/database/seeder"
	"github.com/urfave/cli"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBUser     string
	DBPort     string
	DBPassword string
	DBName     string
	DBHost     string
}

func (server *Server) InitializeDB(dbConfig DBConfig) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", "localhost", dbConfig.DBPassword, dbConfig.DBUser, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed Connection To Database")
	}
}

func (server *Server) dbMigrate() {
	for _, model := range models.RegiseterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Success create migration")

}

func (server *Server) initCommands(config AppConfig, dbConfig DBConfig) {
	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeder.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	server.InitializeDB(dbConfig)

	fmt.Println("Welcome to " + appConfig.AppName)
	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening To Port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
