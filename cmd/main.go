package main

import (
	"log"

	"github.com/ocuprum/menu-constructor/internal/handlers"
	repPgSQL "github.com/ocuprum/menu-constructor/internal/repositories/pgsql"
	"github.com/ocuprum/menu-constructor/internal/services"
	"github.com/ocuprum/menu-constructor/pkg/config"
	"github.com/ocuprum/menu-constructor/pkg/http"
	"github.com/ocuprum/menu-constructor/pkg/pgsql"
)

const (
	CONFIG_NAME = "./configs/config"
	CONFIG_EXTENSION = "yaml" 
	CONFIG_PATH = "."
)

func main() {
	// Setting configs
	conf, err := config.LoadConfig(CONFIG_NAME, CONFIG_EXTENSION, CONFIG_PATH)
	if err != nil {
		log.Fatalf("Error loading a config: %v", err)
	}

	// Connecting to database
	db, err := pgsql.NewPgSQLConnection(conf.PgSQL)
	if err != nil {
		log.Fatalf("Error connecting to pgsql db: %v", err)
	}

	log.Println("Database connected succesfully!")

	foodRep := repPgSQL.NewFoodRepository(db)
	foodSvc := services.NewFoodService(foodRep)
	foodHandler := handlers.NewFoodHandler(foodSvc)

	ingredRep := repPgSQL.NewIngredientRepository(db)
	ingredSvc := services.NewIngredientService(ingredRep)
	ingredHandler := handlers.NewIngredientHandler(ingredSvc)

	
	// Creating new server and starting to listen
	srv := http.NewServer(conf.HTTP, foodHandler, ingredHandler)
	
	log.Printf("We are starting on %v", srv.Addr)
	
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}