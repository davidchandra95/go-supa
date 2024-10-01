package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
	"log"
	"net/http"
	"os"
)

func main() {
	app := NewApp()

	repos := NewRepository(app.supaClient.DB)
	services := NewService(repos)
	handlers := NewHandler(services)

	//email := "mail@mail.com" // update with your email
	//password := "password" // update with your pass
	//ctx := context.Background()
	//_, err := app.SignIn(ctx, email, password)
	//if err != nil {
	//	log.Fatal(err)
	//}

	router := http.NewServeMux()

	router.HandleFunc("POST /todos", handlers.CreateTaskHandler)
	router.HandleFunc("GET /todos", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Println("hehehehe")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("server is listening...")
	log.Fatal(server.ListenAndServe())
}

type App struct {
	supaClient *supa.Client
}

func NewApp() *App {
	godotenv.Load(".env")
	supabaseUrl := os.Getenv("PROJECT_URL")
	supabaseKey := os.Getenv("API_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey, true)
	return &App{
		supaClient: supabase,
	}
}

func (app *App) SignUp(ctx context.Context, email, password string) (*supa.User, error) {
	user, err := app.supaClient.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return &supa.User{}, err
	}

	return user, nil
}

func (app *App) SignIn(ctx context.Context, email, password string) (*supa.AuthenticatedDetails, error) {
	user, err := app.supaClient.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return &supa.AuthenticatedDetails{}, err
	}

	return user, nil
}

func (app *App) SignOut(ctx context.Context, token string) error {
	err := app.supaClient.Auth.SignOut(ctx, token)
	return err
}
