package domain

type Ticket struct {
	Id        int
	Name      string
	Price     int64
	Note      string
	IsDeleted bool
}

type MemberTicket struct {
	Id        int
	UserId    int
	StartDate string
	EndDate   string
	Note      string
	IsDeleted bool
	UpdateAt  string
}
