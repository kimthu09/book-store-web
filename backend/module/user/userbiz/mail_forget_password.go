package userbiz

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/tokenprovider"
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
)

type mailForgetPasswordBiz struct {
	tokenProvider tokenprovider.Provider
}

func NewMailForgetPasswordBiz(
	tokenProvider tokenprovider.Provider) *mailForgetPasswordBiz {
	return &mailForgetPasswordBiz{
		tokenProvider: tokenProvider,
	}
}

func (biz *mailForgetPasswordBiz) MailForgetPassword(
	ctx context.Context,
	email string,
	emailFrom string,
	smtpPass string,
	smtpUser string,
	smtpHost string,
	smtpPort int) error {
	payLoad := tokenprovider.TokenPayloadEmail{
		Email: email,
	}

	token, err := biz.tokenProvider.GenerateTokenForPayLoadEmail(payLoad, 60*60*15)
	if err != nil {
		return common.ErrInternal(err)
	}

	emailData := EmailData{
		URL:       "http://localhost/resetpassword/" + token.Token,
		FirstName: "",
		Subject:   "Your password reset token (valid for 15min)",
	}

	err = SendEmail(
		emailFrom,
		email,
		smtpPass,
		smtpUser,
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
	smtpUser string,
	smtpHost string,
	smtpPort int,
	data *EmailData,
	templateName string) error {

	var body bytes.Buffer

	template, err := ParseTemplateDir("../../mailtemplate")
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

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
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
