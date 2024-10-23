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

	caCert, err := os.ReadFile("/Users/andrew.stark/Documents/Repos/forks/go-redis/doctests/redis_ca.pem")

	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"redis-15313.c34461.eu-west-2-mz.ec2.cloud.rlrcp.com:15313"},
		Username: "default",
		Password: "MrlnkBuSZqO0s0vicIkLnqJXetbSTCan",
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			RootCAs:    caCertPool,
			ServerName: "redis-15313.c34461.eu-west-2-mz.ec2.cloud.rlrcp.com",
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
