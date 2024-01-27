package api

import (
	"database/sql"
	"errors"
	"os"

	"github.com/avearmin/gorage-sale/internal/database"
	"github.com/avearmin/gorage-sale/internal/oauth2"
	"github.com/joho/godotenv"
)

var (
	ErrorLoadENV            = errors.New("error loading .env file")
	ErrorNoDBConn           = errors.New("DB_CONN_STRING string has not been specified")
	ErrorNoPort             = errors.New("PORT has not been specified")
	ErrorNoJWTSecret        = errors.New("JWT_SECRET has not been specified")
	ErrorNoClientID         = errors.New("CLIENT_ID has not been specified")
	ErrorNoClientSecret     = errors.New("CLIENT_SECRET has not been specified")
	ErrorNoOAuthRedirectURL = errors.New("REDIRECT_URL has not been specified")
)

type config struct {
	DB               *database.Queries
	Port             string
	JwtSecret        string
	StateStore       oauth2.StateStore
	ClientID         string
	ClientSecret     string
	OAuthRedirectURL string
}

func CreateConfig() (config, error) {

	if err := godotenv.Load(); err != nil {
		return config{}, ErrorLoadENV
	}

	dbConn := os.Getenv("DB_CONN_STRING")
	if dbConn == "" {
		return config{}, ErrorNoDBConn
	}
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return config{}, err
	}
	dbQueries := database.New(db)

	port := os.Getenv("PORT")
	if port == "" {
		return config{}, ErrorNoPort
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return config{}, ErrorNoJWTSecret
	}

	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		return config{}, ErrorNoClientID
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		return config{}, ErrorNoClientSecret
	}

	oauthRedirectURL := os.Getenv("REDIRECT_URL")
	if oauthRedirectURL == "" {
		return config{}, ErrorNoOAuthRedirectURL
	}

	stateStore := oauth2.NewStateStore()

	cfg := config{
		DB:               dbQueries,
		Port:             port,
		JwtSecret:        jwtSecret,
		ClientID:         clientID,
		ClientSecret:     clientSecret,
		StateStore:       stateStore,
		OAuthRedirectURL: oauthRedirectURL,
	}

	return cfg, nil
}
