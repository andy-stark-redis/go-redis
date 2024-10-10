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
		Addrs:    []string{"localhost:6379", "localhost:6380"},
		Username: "yourUsername",
		Password: "yourPassword",
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
