package utils

import (
	"fmt"
	"time"
)

func GenerateClientID() string {
	return fmt.Sprintf("client-%d", time.Now().UnixNano())
}
