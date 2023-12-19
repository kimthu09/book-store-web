package appctx

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetSecretKey() string
	GetDomain() string
}

type appCtx struct {
	db        *gorm.DB
	secretKey string
	domain    string
}

func NewAppContext(
	db *gorm.DB,
	secretKey string,
	domain string) *appCtx {
	return &appCtx{
		db:        db,
		secretKey: secretKey,
		domain:    domain,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetDomain() string {
	return ctx.domain
}
