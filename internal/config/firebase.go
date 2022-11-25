package config

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type Config struct {
	Key           string
	Id            string
	ProjectId     string
	ClientEmail   string
	ClientId      string
	AuthURI       string
	TokenURI      string
	ClientCertURL string
}

func SetupFirebase() *auth.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	key := os.Getenv("SERVICE_ACC_KEY")
	id := os.Getenv("SERVICE_ACC_KEY_ID")
	projectId := os.Getenv("SERVICE_ACC_PROJECT_ID")
	clientEmail := os.Getenv("SERVICE_ACC_CLIENT_EMAIL")
	clientId := os.Getenv("SERVICE_ACC_CLIENT_ID")
	authURI := os.Getenv("SERVICE_ACC_AUTH_URI")
	tokenURI := os.Getenv("SERVICE_ACC_TOKEN_URI")
	clientCertURL := os.Getenv("SERVICE_ACC_CLIENT_CERT_URL")

	fmt.Printf(key)

	config := Config{Key: key, Id: id, ProjectId: projectId, ClientEmail: clientEmail, ClientId: clientId, AuthURI: authURI, TokenURI: tokenURI, ClientCertURL: clientCertURL}
	templateString := "{\n  \"type\": \"service_account\",\n  \"project_id\": \"{{.ProjectId}}\",\n  \"private_key_id\": \"{{.Id}}\",\n  \"private_key\": \"{{.Key}}\",\n  \"client_email\": \"{{.ClientEmail}}\",\n  \"client_id\": \"{{.ClientId}}\",\n  \"auth_uri\": \"{{.AuthURI}}\",\n  \"token_uri\": \"{{.TokenURI}}\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"{{.ClientCertURL}}\"\n}\n"
	tmp := template.Must(template.New("config").Parse(templateString))

	f, err := os.Create("serviceAccountKey.json")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = tmp.Execute(f, config)

	if err != nil {
		log.Fatal(err)
	}

	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}
	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("Firebase load error")
	}
	return auth
}
