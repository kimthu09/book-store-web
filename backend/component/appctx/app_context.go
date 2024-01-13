package appctx

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetSecretKey() string
	GetStaticPath() string
	GetServerHost() string
	GetEmailFrom() string
	GetSMTPPass() string
	GetSMTPHost() string
	GetSMTPPort() int
}

type appCtx struct {
	db         *gorm.DB
	secretKey  string
	staticPath string
	serverHost string
	emailFrom  string
	smtpPass   string
	smtpHost   string
	smtpPort   int
}

func NewAppContext(
	db *gorm.DB,
	secretKey string,
	staticPath string,
	serverHost string,
	emailFrom string,
	smtpPass string,
	smtpHost string,
	smtpPort int,
) *appCtx {
	return &appCtx{
		db:         db,
		secretKey:  secretKey,
		staticPath: staticPath,
		serverHost: serverHost,
		emailFrom:  emailFrom,
		smtpPass:   smtpPass,
		smtpHost:   smtpHost,
		smtpPort:   smtpPort,
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

func (ctx *appCtx) GetEmailFrom() string {
	return ctx.emailFrom
}

func (ctx *appCtx) GetSMTPPass() string {
	return ctx.smtpPass
}

func (ctx *appCtx) GetSMTPHost() string {
	return ctx.smtpHost
}

func (ctx *appCtx) GetSMTPPort() int {
	return ctx.smtpPort
}
