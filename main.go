package main

import (
	"log"
	"os"

	"github.com/armon/go-socks5"
)

func main() {
	// Retrieve the user credentials from environment variables
	user := os.Getenv("PROXY_USER")
	pass := os.Getenv("PROXY_PASS")

	if user == "" || pass == "" {
		log.Fatal("PROXY_USER and PROXY_PASS environment variables must be set")
	}
	cred := socks5.StaticCredentials{
		user:pass,
	}

	cator := socks5.UserPassAuthenticator{Credentials: cred}

	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{cator},
	}

	server, err := socks5.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting SOCKS5 proxy server on :1080")
	if err := server.ListenAndServe("tcp", ":1080"); err != nil {
		log.Fatal(err)
	}
}

