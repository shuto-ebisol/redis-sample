package main

import (
	"context"
	"fmt"
	"redis-sample/ebica3-domain/entity"
	ebica3_redis "redis-sample/ebica3-infrastructure/repository/redis"

	vo "redis-sample/ebica3-domain/valueobject"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// 接続(redis-clusterを用意する必要があるため、動作確認はしていない)
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{},
	})
	userKeywordClient := ebica3_redis.NewUserKeywordRepository(client)

	userID, _ := vo.NewSharedUserID("user_id")
	userKeyword := entity.UserKeyword{
		UserID:      *userID,
		Keyword:     "keyword",
		MailAddress: "mailaddress",
	}

	// データの保存
	err := userKeywordClient.Put(ctx, "test-key", userKeyword)
	if err != nil {
		fmt.Println(err)
		return
	}

	// データの取得
	result, err := userKeywordClient.Get(ctx, "test-key")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
