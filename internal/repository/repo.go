package repository

import (
	"api-server/internal/rdb"
)

const (
	Account_Table        = "Account"
	Instructor_Table     = "Instructor"
	MemberTicket_Table   = "MemberTicker"
	Ticker_Table         = "Ticket"
	WorkoutBooking_Table = "WorkoutBooking"
)

var rep rdb.IRDB

func InitRepo(cfg rdb.IRDB) {
	if rep == nil {
		rep = cfg
	}
}
