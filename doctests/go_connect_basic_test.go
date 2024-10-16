// EXAMPLE: connect_basic

// STEP_START connect_basic
package example_commands_test

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ExampleClient_connect_basic() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-11919.c335.europe-west2-1.gce.redns.redis-cloud.com:11919",
		Username: "default",
		Password: "RSDBhtpPhBHHNNnWrJSQxgHQCcYtoNMf",
		DB:       0,
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
