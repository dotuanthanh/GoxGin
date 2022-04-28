package domain

type UserClubInformation struct {
	UserId     int
	Phone      string
	Name       string
	Note       string
	IsDeleted  bool
	UpdateAt   string
	IsVerified bool
}

type Account struct {
	Email     string
	Hash      string
	Role      string
	UserId    int
	IsDeleted bool
}
