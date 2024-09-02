# Tiny-Socks

This project sets up a SOCKS5 proxy server with username and password authentication using the `armon/go-socks5` library. The proxy server is containerized using Docker.

## Prerequisites

- Docker must be installed on your system.
- Git should be available if you intend to clone this repository.

## Setup

1. **Clone the repository**

   ```bash
   git clone https://github.com/Vlasov/tiny-socks.git
   cd tiny-socks
   ```

2. **Create the project structure:**

   If not using git, you can create the structure manually:

   ```bash
   mkdir tiny-socks
   cd tiny-socks
   touch main.go go.mod Dockerfile
   ```

3. **Add the code files:**

   - **main.go**

     ```go
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
     		user: pass,
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
```
