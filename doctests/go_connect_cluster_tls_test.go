// EXAMPLE: connect_cluster_tls
// STEP_START connect_cluster_tls
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

func ExampleClient_connect_cluster_tls() {
	ctx := context.Background()

	caCert, err := os.ReadFile("<path_to_redis_ca.pem_file>")

	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"<host>:<port"},
		Username: "default",
		Password: "<password>",
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			RootCAs:    caCertPool,
			ServerName: "<host>",
		},
	})
	// REMOVE_START
	rdb.Del(ctx, "foo")
	// REMOVE_END

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
