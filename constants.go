package gotimed

import "time"

const (
	TIMED_TCP_LISTENING_HOST = "127.0.0.1"
	TIMED_TCP_LISTENING_PORT = 37
)

var RFC_UNIX_DIFF = time.Unix(0, 0).Sub(
	time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
)
