package cache

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"time"
)

type (
	Cache interface {
		Set(key string, obj any) error
		Get(key string, obj any) error
		Push(key string, obj any) error
		Pop(key string, obj any) error
	}

	Memo struct {
		rds *redis.Client
	}
)

func (m *Memo) Set(key string, obj any) error {
	objBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return m.rds.Set(context.Background(), key, objBytes, time.Minute*30).Err()
}

func (m *Memo) Get(key string, obj any) error {
	bytesData, err := m.rds.Get(context.Background(), key).Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(bytesData, &obj)
}

func (m *Memo) Push(key string, obj any) error {
	objBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return m.rds.LPush(context.Background(), key, objBytes).Err()
}

func (m *Memo) Pop(key string, obj any) error {
	bytesData, err := m.rds.RPop(context.Background(), key).Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(bytesData, &obj)
}
