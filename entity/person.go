package entity

const (
	ADMIN      = "Admin"
	MEMBER     = "Member"
	INSTRUCTOR = "Instructor"
	GUEST      = "GUEST"
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

/**
 Entity này mục đích lưu trữ quĩ lương của nhân viên
-Salary là lương thực nhật
-InMonth là lương ở tháng mấy
-Note là lưu thông tin khác tùy chọn dạng text (thưởng , số công vv)
**/

//type Salary struct {
//	ID        int64  `json:"instructorID"`
//	Phone     string `json:"phoneNumber"`
//	Name      string `json:"name"`
//	Salary    string `json:"salary"`
//	Note      string `json:"note"`
//	InMonth   int    `json:"inMonth"`
//	IsDeleted bool   `json:"isDeleted"`
//}

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
