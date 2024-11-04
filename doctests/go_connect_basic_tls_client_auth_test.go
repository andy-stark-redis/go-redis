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

	cert, err := tls.LoadX509KeyPair(
		"<path_to_redis-db-xxxxxxxx.crt_file>",
		"<path_to_redis-db-xxxxxxxx.key_file>",
	)
	if err != nil {
		log.Fatal(err)
	}

	caCert, err := os.ReadFile("<path_to_redis_ca.pem_file>")

	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "<host>:<port>",
		Username: "default",
		Password: "<password>",
		DB:       0,
		TLSConfig: &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		},
	})
	// REMOVE_START
	rdb.Del(ctx, "foo")
	// REMOVE_END

	result, err := rdb.Set(ctx, "foo", "bar", 0).Result()

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
