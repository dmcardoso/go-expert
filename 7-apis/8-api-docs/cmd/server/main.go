package main

import (
	"log"
	"net/http"

	"github.com/dmcardoso/go-expert/7-apis/8-api-docs/configs"
	_ "github.com/dmcardoso/go-expert/7-apis/8-api-docs/docs"
	"github.com/dmcardoso/go-expert/7-apis/8-api-docs/internal/entity"
	"github.com/dmcardoso/go-expert/7-apis/8-api-docs/internal/infra/database"
	"github.com/dmcardoso/go-expert/7-apis/8-api-docs/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Wesley Willians
// @contact.url    http://www.fullcycle.com.br
// @contact.email  atendimento@fullcycle.com.br

// @license.name   Full Cycle License
// @license.url    http://www.fullcycle.com.br

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	userDB := database.NewUser(db)
	productHandler := handlers.NewProductHandler(productDB)
	userHandler := handlers.NewUserHandler(userDB, configs.JWTExpiresIn)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(LogRequest)
	router.Use(commonMiddleware)
	router.Use(middleware.WithValue("jwt", configs.TokenAuth))

	router.Route("/products", func(router chi.Router) {
		router.Use(jwtauth.Verifier(configs.TokenAuth))
		router.Use(jwtauth.Authenticator)

		router.Post("/", productHandler.CreateProduct)
		router.Put("/{id}", productHandler.UpdateProduct)
		router.Delete("/{id}", productHandler.Delete)
		router.Get("/", productHandler.GetProducts)
		router.Get("/{id}", productHandler.GetProduct)
	})

	router.Post("/users", userHandler.Create)
	router.Post("/users/generate-token", userHandler.GetJWT)

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", router)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// r.Context().Value("user")
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
