package entity

type Ticket struct {
	TicketID    int64  `json:"ticketID"`
	Information string `json:"information"`
	Price       int64  `json:"price"`
	Sales       int64  `json:"Sales"`
}
type MemberTicket struct {
	Id         int    `json:"id"`
	StartDate  int64  `json:"startDate"`
	EndDate    int64  `json:"endDate"`
	MemberName string `json:"memberName"`
	Phone      string `json:"phone"`
	Note       string `json:"note"`
}
