package entity

type MemberTicket struct {
	Id        int    `json:"ID"`
	MemberID  int    `json:"memberID"`
	StartDate int64  `json:"startDate"`
	EndDate   int64  `json:"endDate"`
	Note      string `json:"note"`
	IsDeleted bool   `json:"isDeleted"`
	UpdateAt  string `json:"updateAt"`
}

/*
Table Danh sách hội viên
*/
type Member struct {
	MemberID   int    `json:"memberID"`
	MemberName string `json:"memberName"`
	Phone      string `json:"phone"`
	BD         string `json:"birthDay"`
	IsDeleted  bool   `json:"isDeleted"`
	Note       string `json:"note"`
}
