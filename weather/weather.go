package weather

import (
	"time"
)

// 24 phases per cycle, 1 cycle per week
func phase() int64 {
	return time.Now().Unix() / 25200;
}

