// EXAMPLE: connect_basic_tls_client_auth

// STEP_START connect_basic_tls_client_auth
package example_commands_test

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func ExampleClient_connect_basic_tls_client_auth() {
	ctx := context.Background()

	// Load client cert
	cert, err := tls.LoadX509KeyPair("redis_user.crt", "redis_user_private.key")
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := os.ReadFile("redis_ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "yourUsername",
		Password: "yourPassword",
		DB:       0,
		TLSConfig: &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		},
	})

	rdb.Set(ctx, "foo", "bar", 0)
	result, err := rdb.Get(ctx, "foo").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(result) // >>> bar

	// Output:
	// bar
}

// STEP_END
