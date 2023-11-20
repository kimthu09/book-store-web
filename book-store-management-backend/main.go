package main

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln("Error when loading config:", err)
	}

	db, err := connectDatabase(cfg)
	if err != nil {
		log.Fatalln("Error when connecting to database:", err)
	}
	if cfg.Env == "dev" {
		db = db.Debug()
	}

	appCtx := appctx.NewAppContext(db, cfg.SecretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//v1 := r.Group("/v1")
	//{
	//	v1.POST("/login", ginuser.Login(appCtx))
	//}
	//
	//users := v1.Group("/users", middleware.RequireAuth(appCtx))
	//{
	//	users.PATCH("", ginuser.CreateUser(appCtx))
	//}
	//
	//suppliers := v1.Group("/suppliers", middleware.RequireAuth(appCtx))
	//{
	//	suppliers.GET("", ginsupplier.ListSupplier(appCtx))
	//	suppliers.GET("/:id", ginsupplier.SeeDetailSupplier(appCtx))
	//	suppliers.POST("", ginsupplier.CreateSupplier(appCtx))
	//	suppliers.PATCH("/:id", ginsupplier.UpdateInfoSupplier(appCtx))
	//	suppliers.POST("/:id/pay", ginsupplier.PaySupplier(appCtx))
	//}
	//
	//importNotes := v1.Group("/importNotes", middleware.RequireAuth(appCtx))
	//{
	//	importNotes.GET("", ginimportnote.ListImportNote(appCtx))
	//	importNotes.GET("/:id", ginimportnote.SeeDetailImportNote(appCtx))
	//	importNotes.POST("", ginimportnote.CreateImportNote(appCtx))
	//	importNotes.PATCH("/:id", ginimportnote.ChangeStatusImportNote(appCtx))
	//}
	//
	//inventoryCheckNotes := v1.Group("/inventoryCheckNotes", middleware.RequireAuth(appCtx))
	//{
	//	inventoryCheckNotes.GET("", gininventorychecknote.ListInventoryCheckNote(appCtx))
	//	inventoryCheckNotes.GET("/:id", gininventorychecknote.SeeDetailInventoryCheckNote(appCtx))
	//	inventoryCheckNotes.POST("", gininventorychecknote.CreateInventoryCheckNote(appCtx))
	//}

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

func connectDatabase(cfg *appConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db.Debug(), nil
}
