package userbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/tokenprovider"
	"book-store-management-backend/module/user/usermodel"
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type MailForgetPasswordRepo interface {
	VerifyUser(
		ctx context.Context,
		email string) error
}

type mailForgetPasswordBiz struct {
	repo          MailForgetPasswordRepo
	tokenProvider tokenprovider.Provider
}

func NewMailForgetPasswordBiz(
	repo MailForgetPasswordRepo,
	tokenProvider tokenprovider.Provider) *mailForgetPasswordBiz {
	return &mailForgetPasswordBiz{
		repo:          repo,
		tokenProvider: tokenProvider,
	}
}

func (biz *mailForgetPasswordBiz) MailForgetPassword(
	ctx context.Context,
	data *usermodel.ReqMailForgotPassword,
	emailFrom string,
	smtpPass string,
	smtpHost string,
	smtpPort int) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.VerifyUser(ctx, data.Email); err != nil {
		return err
	}

	payLoad := tokenprovider.TokenPayloadEmail{
		Email: data.Email,
	}
	token, err := biz.tokenProvider.GenerateTokenForPayLoadEmail(payLoad, 60*60*common.MinuteVerifyEmail)
	if err != nil {
		return common.ErrInternal(err)
	}

	emailData := EmailData{
		URL:       "http://localhost/forgot-password/" + token.Token,
		FirstName: data.Email,
		Subject: "Your password reset token (valid for " +
			strconv.Itoa(common.MinuteVerifyEmail) + " min)",
	}

	err = SendEmail(
		emailFrom,
		data.Email,
		smtpPass,
		smtpHost,
		smtpPort,
		&emailData,
		"resetPassword.html")
	if err != nil {
		return err
	}

	return nil
}

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func SendEmail(
	emailFrom string,
	emailTo string,
	smtpPass string,
	smtpHost string,
	smtpPort int,
	data *EmailData,
	templateName string) error {

	var body bytes.Buffer

	template, err := ParseTemplateDir("mailtemplate")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template = template.Lookup(templateName)
	template.Execute(&body, &data)
	fmt.Println(template.Name())

	m := gomail.NewMessage()

	m.SetHeader("From", emailFrom)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	fmt.Println(emailFrom)
	fmt.Println(smtpPass)

	d := gomail.NewDialer(smtpHost, smtpPort, emailFrom, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	fmt.Println("Am parsing templates...")

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}
