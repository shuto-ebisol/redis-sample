package redis

import (
	"context"
	"encoding/json"
	"errors"
	"redis-sample/ebica3-domain/entity"
	"time"

	"redis-sample/ebica3-domain/repository"

	"redis-sample/ebica3-domain/service"

	"github.com/redis/go-redis/v9"
)

var (
	redisPutDataError      = errors.New("failed to put data to redis")
	redisDataParsingError  = errors.New("failed to parse data from redis")
	redisDataNotFoundError = errors.New("data not found in redis")
	redisRemoveDataError   = errors.New("failed to remove data from redis")
)

const (
	UserIdKey      = "user_id"
	UserKeywordKey = "keyword"
	MailAddressKey = "mailaddress"
	ProductCodeKey = "product_code"
)

type keywordRedisClient struct {
	client         *redis.ClusterClient
	expirationTime time.Duration
}

// NewUserKeywordRepository UserKeywordRepository interfaceの値を生成する（Redis）
func NewUserKeywordRepository(client *redis.ClusterClient) repository.UserKeywordRepository {
	return &keywordRedisClient{
		client:         client,
		expirationTime: service.KeywordExpirationTime,
	}
}

// Put 合言葉の保存
func (c *keywordRedisClient) Put(ctx context.Context, key string, value entity.UserKeyword) error {
	userKeyword := convertEntityToDocument(value)
	b, err := userKeyword.MarshalJSON()
	if err != nil {
		return errors.Join(redisPutDataError, err)
	}

	err = c.client.Set(ctx, key, string(b), c.expirationTime).Err()
	if err != nil {
		return errors.Join(redisPutDataError, err)
	}

	return nil
}

// Get 合言葉の取得
func (c *keywordRedisClient) Get(ctx context.Context, key string) (*entity.UserKeyword, error) {
	value, err := c.client.Get(ctx, key).Result()
	if err != nil || len(value) == 0 {
		return nil, errors.Join(redisDataNotFoundError, err)
	}
	var userKeyword userKeyword
	err = json.Unmarshal([]byte(value), &userKeyword)
	if err != nil {
		return nil, errors.Join(redisDataParsingError, err)
	}
	return userKeyword.toDomainEntity()
}
