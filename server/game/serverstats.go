package main

import "time"

// Some rudimentary server statistics
type ServerStats struct {
	StartTime    int64
	MessageCount uint64
	UserPeak     int
}

func NewServerStats() *ServerStats {
	s := new(ServerStats)
	s.StartTime = time.Nanoseconds()
	return s
}

func (this *ServerStats) Update(users int) {
	this.MessageCount++

	if users > this.UserPeak {
		this.UserPeak = users
	}
}

func (this *ServerStats) Uptime() int64 {
	return time.Nanoseconds() - this.StartTime
}
