package repository

import (
	"context"
	"redis-sample/ebica3-domain/entity"
)

// UserKeywordRepository　合言葉を保存するPersistence Interface
//
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=mock/$GOFILE
type UserKeywordRepository interface {
	// Put 合言葉の保存
	Put(ctx context.Context, key string, value entity.UserKeyword) error
	// Get 合言葉の取得
	Get(ctx context.Context, key string) (*entity.UserKeyword, error)
}
