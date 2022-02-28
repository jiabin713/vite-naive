package utils

import (
	"strconv"

	"github.com/sony/sonyflake"
)

func GenerateID() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		panic("sonyflake not created2")
	}
	str_id := strconv.FormatUint(id, 10)
	return str_id
}
