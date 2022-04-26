package entity

import "api-server/pck/enum"

type UserClubInformation struct {
	UserId     int    `json:"userId"`
	Phone      string `json:"phone"`
	Name       string `json:"name"`
	Note       string `json:"note"`
	IsDeleted  bool   `json:"isDeleted"`
	UpdateAt   string `json:"updateAt"`
	IsVerified bool   `json:"isVerified"`
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
	case enum.ADMIN:
		return true
	case enum.MEMBER:
		return true
	case enum.INSTRUCTOR:
		return true
	case enum.GUEST:
		return true
	default:
		return false
	}
}
