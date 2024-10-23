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
	cert, err := tls.LoadX509KeyPair(
		"/Users/andrew.stark/Documents/Repos/forks/go-redis/doctests/redis-db-12592910.crt",
		"/Users/andrew.stark/Documents/Repos/forks/go-redis/doctests/redis-db-12592910.key",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := os.ReadFile("/Users/andrew.stark/Documents/Repos/forks/go-redis/doctests/redis_ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-14669.c338.eu-west-2-1.ec2.redns.redis-cloud.com:14669",
		Username: "default",
		Password: "jj7hRGi1K22vop5IDFvAf8oyeeF98s4h",
		DB:       0,
		TLSConfig: &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		},
	})

	result, err := rdb.Set(ctx, "fox", "bar", 0).Result()

	if err != nil {
		panic(err)
	}

	result, err = rdb.Get(ctx, "foo").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(result) // >>> bar

	// Output:
	// bar
}

// STEP_END
