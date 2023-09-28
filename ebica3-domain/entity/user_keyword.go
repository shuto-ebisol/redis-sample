package entity

import (
	vo "redis-sample/ebica3-domain/valueobject"
)

// UserKeyword 合言葉に関する値の構造体
type UserKeyword struct {
	UserID      vo.SharedUserID
	Keyword     string
	MailAddress string
}
