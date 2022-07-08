package main

import (
	"fmt"
	"net/http"

	"github.com/matheuslc/guiomar/src/chef"
	"github.com/matheuslc/guiomar/src/context"
	"github.com/matheuslc/guiomar/src/food"

	"github.com/gorilla/mux"

	_ "github.com/matheuslc/guiomar/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Guiomar API
// @version 1.0
// @description Guiomar private and public API docs.
// @termsOfService http://swagger.io/terms/

// @contact.name Matheus Carmo
// @contact.url http://www.swagger.io/support
// @contact.email mematheuslc@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7777
func main() {
	app, _ := context.NewAppContext()
	router := mux.NewRouter()

	router.HandleFunc("/api/chefs", chef.NewChefHandlerWrapper(app.ChefRepository)).Methods("POST")
	router.HandleFunc("/api/foods", food.NewFoodHandlerWrapper(app.FoodRepository)).Methods("POST")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	server := &http.Server{
		Handler: router,
		Addr:    ":7777",
	}

	fmt.Println("We are online! Running on localhost:7777")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
