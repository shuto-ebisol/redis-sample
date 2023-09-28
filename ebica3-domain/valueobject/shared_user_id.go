package valueobject

// SharedUserID 共用ユーザーID
//
//	(1文字以上20文字以下)
type SharedUserID struct {
	value string
}

func NewSharedUserID(v string) (*SharedUserID, error) {
	return &SharedUserID{
		value: v,
	}, nil
}

// Value 共用ユーザーIDを返す
func (vo SharedUserID) Value() string {
	return vo.value
}
