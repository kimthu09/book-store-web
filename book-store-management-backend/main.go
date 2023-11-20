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

func main() {
	env, err := godotenv.Read()

	if err != nil {
		log.Fatalln("Error when loading .env", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env["DB_USERNAME"], env["DB_PASSWORD"], env["DB_HOST"], env["DB_DATABASE"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db = db.Debug()

	secretKey := env["SECRET_KEY"]

	fmt.Println("DB connected", db)
	appCtx := appctx.NewAppContext(db, secretKey)

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

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
