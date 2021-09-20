package global

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	var jtime = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(jtime), nil
}

type UserResponse struct {
	Id       int32    `json:"id"`
	Mobile   string   `json:"mobile"`
	Nickname string   `json:"nickname"`
	Gender   int32    `json:"gender"`
	Birthday JsonTime `json:"birthday"`
}
