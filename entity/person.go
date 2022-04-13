package entity

const (
	ADMIN      = "Admin"
	MEMBER     = "Member"
	INSTRUCTOR = "Instructor"
	GUEST      = "Guest"
)

type UserClubInformation struct {
	UserId     int    `json:"userId"`
	Phone      string `json:"phone"`
	Name       string `json:"name"`
	Note       string `json:"note"`
	IsDeleted  bool   `json:"isDeleted"`
	UpdateAt   string `json:"updateAt"`
	IsVerified bool   `json:"IsVerified"`
}

type Account struct {
	Email     string `json:"email"`
	Hash      string `json:"hash"`
	Role      string `json:"role"`
	UserId    int    `json:"userId"`
	IsDeleted bool   `json:"isDeleted"`
}

func CheckValidRoles(person *Account) bool {
	role := person.Role
	switch role {
	case ADMIN:
		return true
	case MEMBER:
		return true
	case INSTRUCTOR:
		return true
	case GUEST:
		return true
	default:
		return false
	}
}
