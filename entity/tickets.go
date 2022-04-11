package entity

type Ticket struct {
	Id         int    `json:"id"`
	TicketName string `json:"ticketName"`
	Price      int64  `json:"price"`
	Note       string `json:"note"`
	IsDeleted  bool   `json:"isDeleted"`
}

type MemberTicket struct {
	Id        int    `json:"Id"`
	UserId    int    `json:"userId"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Note      string `json:"note"`
	IsDeleted bool   `json:"isDeleted"`
	UpdateAt  string `json:"updateAt"`
}
