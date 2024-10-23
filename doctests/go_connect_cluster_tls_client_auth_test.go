// EXAMPLE: connect_cluster_tls_client_auth

// STEP_START connect_cluster_tls_client_auth

package example_commands_test

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func ExampleClient_connect_cluster_tls_client_auth() {
	ctx := context.Background()

	// Load client cert
	cert, err := tls.LoadX509KeyPair(
		"/Users/andrew.stark/Documents/Repos/forks/go-redis/doctests/redis-db-12605866.crt",
		"/Users/andrew.stark/Documents/Repos/forks/go-redis/doctests/redis-db-12605866.key",
	)

	if err != nil {
		panic(err)
	}

	// Load CA cert
	caCert, err := os.ReadFile("/Users/andrew.stark/Documents/Repos/forks/go-redis/doctests/redis_ca.pem")

	if err != nil {
		panic(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"redis-15313.c34461.eu-west-2-mz.ec2.cloud.rlrcp.com:15313"},
		Username: "default",
		Password: "MrlnkBuSZqO0s0vicIkLnqJXetbSTCan",
		TLSConfig: &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
			ServerName:   "redis-15313.c34461.eu-west-2-mz.ec2.cloud.rlrcp.com",
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
