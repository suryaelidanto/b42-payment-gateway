package main

import (
	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


func main() {
	// env
	errEnv := godotenv.Load()
    if errEnv != nil {
		panic("Failed to load env file")
    }
	
	// initial DB
	mysql.DatabaseInit()
	
	// run migration
	database.RunMigration()
	
	r := mux.NewRouter()
	
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())
	
	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	
	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	// var CONFIG_SMTP_HOST = "smtp.gmail.com"
	// var CONFIG_SMTP_PORT = 587
	// var CONFIG_SENDER_NAME = "DumbMerch <demo.dumbways@gmail.com>"
	// var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
	// var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

	// mailer := gomail.NewMessage()
	// mailer.SetHeader("From", CONFIG_SENDER_NAME)
	// mailer.SetHeader("To", "jody.septiawan5@gmail.com")
	// mailer.SetHeader("Subject", "Transaction Status")
	// mailer.SetBody("text/html", "<h1> Pisang Coklat<h1>")

	// dialer := gomail.NewDialer(
	// 	CONFIG_SMTP_HOST,
	// 	CONFIG_SMTP_PORT,
	// 	CONFIG_AUTH_EMAIL,
	// 	CONFIG_AUTH_PASSWORD,
	// )

	// err := dialer.DialAndSend(mailer)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// log.Println("Mail sent!")
	
	var port = "5000"
	fmt.Println("server running localhost:"+port)
	http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
