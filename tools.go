package gotimed

import "time"

func Rfc868ToUnix(rfctime int64) int64 {
	t := time.Unix(rfctime, 0)
	nt := t.Add(RFC_UNIX_DIFF)
	//	nt := t.AddDate(70, 0, 0)
	return nt.Unix()
}

func UnixToRfc868(rfctime int64) int64 {
	t := time.Unix(rfctime, 0)
	nt := t.Add(-RFC_UNIX_DIFF)
	//	nt := t.AddDate(-70, 0, 0)
	return nt.Unix()
}
