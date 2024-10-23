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
		Addr:     "redis-14669.c338.eu-west-2-1.ec2.redns.redis-cloud.com:14669",
		Username: "default",
		Password: "jj7hRGi1K22vop5IDFvAf8oyeeF98s4h",
		DB:       0,
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
