package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"receipt-processor/models"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// RedisStore struct for managing Redis operations
type RedisStore struct {
	Client *redis.Client
}

// NewRedisStore initializes Redis connection
func NewRedisStore(redisAddr string) *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	// Check Redis connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error: Failed to connect to Redis:", err)
	}
	return &RedisStore{Client: client}
}

// SaveReceipt stores receipt data in Redis with error handling
func (r *RedisStore) SaveReceipt(id string, receipt models.Receipt) error {
	data, err := json.Marshal(receipt)
	if err != nil {
		return err
	}
	if err := r.Client.Set(ctx, id, data, 0).Err(); err != nil {
		return fmt.Errorf("failed to save receipt: %v", err)
	}
	return nil
}

// GetReceipt retrieves receipt data by ID with error handling
func (r *RedisStore) GetReceipt(id string) (models.Receipt, error) {
	var receipt models.Receipt
	data, err := r.Client.Get(ctx, id).Result()
	if err == redis.Nil {
		return receipt, fmt.Errorf("receipt not found")
	} else if err != nil {
		return receipt, fmt.Errorf("redis error: %v", err)
	}
	json.Unmarshal([]byte(data), &receipt)
	return receipt, nil
}