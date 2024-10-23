// EXAMPLE: connect_cluster

// STEP_START connect_cluster
package example_commands_test

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ExampleClient_connect_cluster() {
	ctx := context.Background()

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"redis-13891.c34425.eu-west-2-mz.ec2.cloud.rlrcp.com:13891"},
		Username: "default",
		Password: "wtpet4pI5EgyJHyldPwR7xM7GaZB0EcG",
	})

	rdb.Set(ctx, "foo", "bar", 0).Result()

	result, err := rdb.Get(ctx, "foo").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(result) // >>> bar

	// Output:
	// bar
}

// STEP_END
