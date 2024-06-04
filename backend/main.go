package main

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/docs"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/author/authortransport"
	"book-store-management-backend/module/book/booktransport"
	"book-store-management-backend/module/booktitle/booktitletransport"
	"book-store-management-backend/module/customer/customertransport/gincustomer"
	"book-store-management-backend/module/dashboard/dashboardtransport/gindashboard"
	"book-store-management-backend/module/feature/featuretransport/ginfeature"
	"book-store-management-backend/module/importnote/importnotetransport/ginimportnote"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotetransport/gininventorychecknote"
	"book-store-management-backend/module/invoice/invoicetransport/gininvoice"
	"book-store-management-backend/module/role/roletransport/ginrole"
	"book-store-management-backend/module/salereport/salereporttransport/ginsalereport"
	"book-store-management-backend/module/shopgeneral/shopgeneraltransport/ginshopgeneral"
	ginstockreports "book-store-management-backend/module/stockreport/stockreporttransport/ginstockreport"
	"book-store-management-backend/module/supplierdebtreport/supplierdebtreporttransport/ginsupplierdebtreport"
	"book-store-management-backend/module/upload/uploadtransport"
	"strconv"
	"time"

	"book-store-management-backend/module/category/categorytransport"
	"book-store-management-backend/module/publisher/publishertransport"
	"book-store-management-backend/module/supplier/suppliertransport/ginsupplier"
	"book-store-management-backend/module/user/usertransport/ginuser"
	"fmt"
	"log"

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

	StaticPath string
	ServerHost string

	DBUsername string
	DBPassword string
	DBHost     string
	DBDatabase string

	SecretKey string

	EmailFrom string
	SMTPUser  string
	SMTPass   string
	SMTHost   string
	SMTPort   int
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
	db, err := connectDatabaseWithRetryIn60s(cfg)
	if err != nil {
		log.Fatalln("Error when connecting to database:", err)
	}

	if cfg.Env == "dev" {
		db = db.Debug()
	}

	appCtx := appctx.NewAppContext(
		db,
		cfg.SecretKey,
		cfg.StaticPath,
		cfg.ServerHost,
		cfg.EmailFrom,
		cfg.SMTPass,
		cfg.SMTHost,
		cfg.SMTPort)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.Recover(appCtx))

	docs.SwaggerInfo.BasePath = "/v1"
	v1 := r.Group("/v1")
	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		uploadtransport.SetupRoutes(v1, appCtx)

		authortransport.SetupRoutes(v1, appCtx)
		categorytransport.SetupRoutes(v1, appCtx)
		booktitletransport.SetupRoutes(v1, appCtx)
		booktransport.SetupRoutes(v1, appCtx)
		publishertransport.SetupRoutes(v1, appCtx)
		gininvoice.SetupRoutes(v1, appCtx)
		ginimportnote.SetupRoutes(v1, appCtx)
		gininventorychecknote.SetupRoutes(v1, appCtx)
		ginsupplier.SetupRoutes(v1, appCtx)
		gincustomer.SetupRoutes(v1, appCtx)
		ginrole.SetupRoutes(v1, appCtx)
		ginfeature.SetupRoutes(v1, appCtx)
		ginuser.SetupRoutes(v1, appCtx)
		ginshopgeneral.SetupRoutes(v1, appCtx)
		gindashboard.SetupRoutes(v1, appCtx)
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

	port, _ := strconv.Atoi(env["SMTPORT"])

	return &appConfig{
		Port:       env["PORT"],
		Env:        env["GO_ENV"],
		StaticPath: env["STATIC_PATH"],
		ServerHost: env["SERVER_HOST"],
		DBUsername: env["DB_USERNAME"],
		DBPassword: env["DB_PASSWORD"],
		DBHost:     env["DB_HOST"],
		DBDatabase: env["DB_DATABASE"],
		SecretKey:  env["SECRET_KEY"],
		EmailFrom:  env["EMAILFROM"],
		SMTPUser:   env["SMTPUSER"],
		SMTPass:    env["SMTPASS"],
		SMTHost:    env["SMTHOST"],
		SMTPort:    port,
	}, nil
}

func connectDatabaseWithRetryIn60s(cfg *appConfig) (*gorm.DB, error) {
	const timeRetry = 60 * time.Second
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
