package main

import (
	"book-store-management-backend/component/appctx"
	docs "book-store-management-backend/docs"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authortransport"
	"book-store-management-backend/module/book/booktransport"
	booktitletransport "book-store-management-backend/module/booktitle/booktitletransport"
	"book-store-management-backend/module/feature/featuretransport/ginfeature"
	"book-store-management-backend/module/importnote/importnotetransport/ginimportnote"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotetransport/gininventorychecknote"
	"book-store-management-backend/module/role/roletransport/ginrole"
	"book-store-management-backend/module/salereport/salereporttransport/ginsalereport"
	ginstockreports "book-store-management-backend/module/stockreport/stockreporttransport/ginstockreport"
	"book-store-management-backend/module/supplierdebtreport/supplierdebtreporttransport/ginsupplierdebtreport"
	"time"

	"book-store-management-backend/module/category/categorytransport"
	"book-store-management-backend/module/publisher/publishertransport"
	"book-store-management-backend/module/supplier/suppliertransport/ginsupplier"
	"book-store-management-backend/module/user/usertransport/ginuser"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type appConfig struct {
	Port string
	Env  string

	DBUsername string
	DBPassword string
	DBHost     string
	DBDatabase string

	SecretKey string
}

// @title           Book Store Management API
// @description     This is a sample server Book Store Management API server.
// @version         1.0

// @contact.name   Bui Vi Quoc
// @contact.url    https://www.facebook.com/bviquoc/
// @contact.email  21520095@gm.uit.edu.vn

// @host localhost:8080
// @BasePath  /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln("Error when loading config:", err)
	}

	fmt.Println("Connecting to database...")
	db, err := connectDatabaseWithRetryIn30s(cfg)
	if err != nil {
		log.Fatalln("Error when connecting to database:", err)
	}

	if cfg.Env == "dev" {
		db = db.Debug()
	}

	appCtx := appctx.NewAppContext(db, cfg.SecretKey)

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	docs.SwaggerInfo.BasePath = "/v1"
	v1 := r.Group("/v1")
	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		authortransport.SetupRoutes(v1, appCtx)
		categorytransport.SetupRoutes(v1, appCtx)
		booktitletransport.SetupRoutes(v1, appCtx)
		booktransport.SetupRoutes(v1, appCtx)
		publishertransport.SetupRoutes(v1, appCtx)
		ginimportnote.SetupRoutes(v1, appCtx)
		gininventorychecknote.SetupRoutes(v1, appCtx)
		ginsupplier.SetupRoutes(v1, appCtx)
		ginrole.SetupRoutes(v1, appCtx)
		ginfeature.SetupRoutes(v1, appCtx)
		ginuser.SetupRoutes(v1, appCtx)

		report := v1.Group("/reports")
		{
			ginstockreports.SetupRoutes(report, appCtx)
			ginsupplierdebtreport.SetupRoutes(report, appCtx)
			ginsalereport.SetupRoutes(report, appCtx)
		}
	}

	if err := r.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatalln("Error running server:", err)
	}
}

func loadConfig() (*appConfig, error) {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatalln("Error when loading .env", err)
	}

	return &appConfig{
		Port:       env["PORT"],
		Env:        env["ENVIRONMENT"],
		DBUsername: env["DB_USERNAME"],
		DBPassword: env["DB_PASSWORD"],
		DBHost:     env["DB_HOST"],
		DBDatabase: env["DB_DATABASE"],
		SecretKey:  env["SECRET_KEY"],
	}, nil
}

func connectDatabaseWithRetryIn30s(cfg *appConfig) (*gorm.DB, error) {
	const timeRetry = 30 * time.Second
	var connectDatabase = func(cfg *appConfig) (*gorm.DB, error) {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBDatabase)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return db.Debug(), nil
	}

	var db *gorm.DB
	var err error

	deadline := time.Now().Add(timeRetry)

	for time.Now().Before(deadline) {
		log.Println("Connecting to database...")
		db, err = connectDatabase(cfg)
		if err == nil {
			return db, nil
		}
		time.Sleep(time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after retrying for 10 seconds: %w", err)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
