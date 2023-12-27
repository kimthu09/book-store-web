package appctx

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetSecretKey() string
	GetStaticPath() string
	GetServerHost() string
}

type appCtx struct {
	db         *gorm.DB
	secretKey  string
	staticPath string
	serverHost string
}

func NewAppContext(
	db *gorm.DB,
	secretKey string,
	staticPath string,
	serverHost string,
) *appCtx {
	return &appCtx{
		db:         db,
		secretKey:  secretKey,
		staticPath: staticPath,
		serverHost: serverHost,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetStaticPath() string {
	return ctx.staticPath
}

func (ctx *appCtx) GetServerHost() string {
	return ctx.serverHost
}
