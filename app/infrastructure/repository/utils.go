package irepository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func countDocPrefixedKeys(rc *redis.Client, prefix string) (int, error) {
	ctx := context.Background()
	var cursor uint64
	var count int

	for {
		keys, newCursor, err := rc.Scan(ctx, cursor, prefix+"*", 0).Result()
		if err != nil {
			return 0, err
		}

		count += len(keys)

		if newCursor == 0 {
			break
		}
		cursor = newCursor
	}

	return count, nil
}
