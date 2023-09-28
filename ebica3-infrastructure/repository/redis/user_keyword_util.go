package redis

import (
	"encoding/json"
	"fmt"
	"redis-sample/ebica3-domain/entity"
	vo "redis-sample/ebica3-domain/valueobject"
)

type userKeyword struct {
	UserID      string `json:"user_id"`
	Keyword     string `json:"keyword"`
	MailAddress string `json:"mailaddress"`
}

func (c *userKeyword) valid() error {
	if c.UserID == "" {
		return fmt.Errorf("user_id is empty")
	}
	if c.Keyword == "" {
		return fmt.Errorf("keyword is empty")
	}
	if c.MailAddress == "" {
		return fmt.Errorf("mailaddress is empty")
	}
	return nil
}

func (c *userKeyword) toDomainEntity() (*entity.UserKeyword, error) {
	userID, err := vo.NewSharedUserID(c.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert 'customerID' from '%v': %s", c.UserID, err.Error())
	}

	userKeyword := entity.UserKeyword{
		UserID: *userID,
	}

	return &userKeyword, nil
}

func convertEntityToDocument(e entity.UserKeyword) userKeyword {
	return userKeyword{
		UserID: e.UserID.Value(),
	}
}

func (k *userKeyword) MarshalJSON() ([]byte, error) {
	return json.Marshal(k)
}

func (k *userKeyword) UnmarshalJSON(b []byte) error {
	var userKeyword userKeyword
	err := json.Unmarshal(b, &userKeyword)
	if err != nil {
		fmt.Println(err)
	}
	return userKeyword.valid()
}
