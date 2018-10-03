package gontpd

import "time"

func Rfc868ToUnix(rfctime int64) int64 {
	t := time.Unix(rfctime, 0)
	nt := t.Add(RFC_UNIX_DIFF)
	return nt.Unix()
}

func UnixToRfc(rfctime int64) int64 {
	t := time.Unix(rfctime, 0)
	nt := t.Add(-RFC_UNIX_DIFF)
	return nt.Unix()
}
