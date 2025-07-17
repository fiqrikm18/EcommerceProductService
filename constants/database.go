package constants

import "time"

const (
	MaxIdleConn     int = 10
	MaxOpenConn     int = 10
	ConnMaxLifetime     = 15 * time.Minute
	ConnMaxIdleTime     = 15 * time.Minute
)
