package configs

import (
	"context"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func FirebaseAuthSetup() *auth.Client {
	serviceKey, err := filepath.Abs("configs/serviceKey.json")
	if err != nil {
		panic("Can't load serviceKey.json")
	}

	options := option.WithCredentialsFile(serviceKey)

	app, err := firebase.NewApp(context.Background(), nil, options)
	if err != nil {
		panic(err.Error())
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(err.Error())
	}
	return auth
}
