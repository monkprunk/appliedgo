package api

import (
	"github.com/gin-gonic/gin"

	"Advance-Golang-Programming/advanced/final/merchant"
	"Advance-Golang-Programming/advanced/final/middleware"
)

type Route struct {
	Name        string
	Path        string
	Method      string
	Endpoint    gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

// func Init() http.Handler {

// 	//merchantRepo := merchant.NewMongoDB(dbConn.MongoDBConn.Session)
// 	merchantRepo := merchant.NewMemoryDB()
// 	merchantSrv := merchant.NewService(merchantRepo)
// 	merchantHandler := merchant.NewHandler(merchantSrv)

// 	apiv1 := []Route{
// 		{
// 			Name:        "register merchant",
// 			Method:      http.MethodPost,
// 			Path:        "/merchant/register",
// 			Endpoint:    merchantHandler.Register,
// 			Middlewares: []gin.HandlerFunc{},
// 		},
// 		{
// 			Name:        "inquiry merchant",
// 			Method:      http.MethodPost,
// 			Path:        "/merchant/:id",
// 			Endpoint:    merchantHandler.Information,
// 			Middlewares: []gin.HandlerFunc{},
// 		},
// 		{
// 			Name:        "List All Products",
// 			Method:      http.MethodPost,
// 			Path:        "/merchant/:id/products",
// 			Endpoint:    merchantHandler.ListAllProduct,
// 			Middlewares: []gin.HandlerFunc{},
// 		},
// 		{
// 			Name:        "Add Product",
// 			Method:      http.MethodPost,
// 			Path:        "/merchant/product/:id",
// 			Endpoint:    merchantHandler.AddProduct,
// 			Middlewares: []gin.HandlerFunc{},
// 		},
// 	}

// 	ro := gin.New()
// 	// ro.Use(cors.Default())
// 	// ro.Use(middleware.Recover)
// 	// ro.Use(middleware.MongoPool(conf))
// 	// ro.Use(middleware.NotFoundOnProduction(conf))

// 	var m middleware.Middleware
// 	ro.Use(m.Request)

// 	for _, route := range apiv1 {
// 		ro.Handle(route.Method, route.Path, append(route.Middlewares, route.Endpoint)...)
// 	}
// 	ro.Run()
// 	return ro
// }

//=====================================================================

func MiddlewareHttp() {

	merchantRepo := merchant.NewMemoryDB()
	merchantSrv := merchant.NewService(merchantRepo)
	merchantHandler := merchant.NewHandler(merchantSrv)

	r := gin.New()

	var m middleware.Middleware
	r.Use(m.Request)

	//merchant := r.Group("/merchant")
	r.POST("/merchant/register", merchantHandler.Register)
	r.GET("/merchant/:id/", merchantHandler.Information)
	r.GET("/merchant/:id/products", merchantHandler.ListAllProduct)
	r.POST("/merchant/product/:id", merchantHandler.AddProduct)
	r.Run()
}
