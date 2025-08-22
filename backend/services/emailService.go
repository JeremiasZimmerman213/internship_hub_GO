package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func SendVerificationEmail(email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Verify Your Email - Internship Hub")

	verifyURL := fmt.Sprintf("%s/verify-email?token=%s",
		os.Getenv("FRONTEND_URL"), token)

	body := fmt.Sprintf(`
        <h2>Welcome to Internship Hub!</h2>
        <p>Thank you for signing up! Please click the link below to verify your email address:</p>
        <a href="%s" style="background-color: #4CAF50; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px;">Verify Email</a>
        <p>Or copy and paste this link in your browser: %s</p>
        <p>This link expires in 24 hours.</p>
        <p>If you didn't create an account, please ignore this email.</p>
    `, verifyURL, verifyURL)

	m.SetBody("text/html", body)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASSWORD"),
	)

	return d.DialAndSend(m)
}

func GenerateVerificationToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
