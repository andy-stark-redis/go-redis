// EXAMPLE: connect_cluster_tls

// STEP_START connect_cluster_tls
package example_commands_test

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ExampleClient_connect_cluster_tls() {
	ctx := context.Background()

	// Load client cert
	cert, err := tls.LoadX509KeyPair("redis_user.crt", "redis_user_private.key")

	if err != nil {
		panic(err)
	}

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"localhost:6379", "localhost:6380"},
		Username: "yourUsername",
		Password: "yourPassword",
		TLSConfig: &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{cert},
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
