package main

func main() {
	//dsn := "root:025020@tcp(127.0.0.1:3307)/book_store_management?charset=utf8mb4&parseTime=True&loc=Local"
	//secretKey := "123456789"
	//
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//db = db.Debug()
	//
	//appCtx := appctx.NewAppContext(db, secretKey)
	//
	//r := gin.Default()
	//r.Use(middleware.Recover(appCtx))
	//
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
	//
	//err = r.Run(":8080")
	//if err != nil {
	//	return
	//}
}
