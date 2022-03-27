package entity

const (
	ADMIN      = "Admin"
	MEMBER     = "Member"
	INSTRUCTOR = "Instructor"
)

/*
Entity này lưu trữ dữ liệu của acocunt hệ thống
-Password (hash code) => phục vụ mục đích login
-IsDeleted là soft-delete check xem có tồn tại hay đã bị xóa
-UpdateAt : lưu lại thời gian lần cập nhật gần nhất
*/

type Person struct {
	ID        int64  `json:"ID"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Note      string `json:"note"`
	Role      string `json:"role"`
	IsDeleted bool   `json:"isDeleted"`
	UpdateAt  string `json:"updateAt"`
}

/**
 Entity này mục đích lưu trữ quĩ lương của nhân viên
-Salary là lương thực nhật
-InMonth là lương ở tháng mấy
-Note là lưu thông tin khác tùy chọn dạng text (thưởng , số công vv)
**/

type Salary struct {
	ID        int64  `json:"instructorID"`
	Phone     string `json:"phoneNumber"`
	Name      string `json:"name"`
	Salary    string `json:"salary"`
	Note      string `json:"note"`
	InMonth   int    `json:"inMonth"`
	IsDeleted bool   `json:"isDeleted"`
}

func CheckValidRoles(person *Person) bool {
	role := person.Role
	switch role {
	case ADMIN:
		return true
	case MEMBER:
		return true
	case INSTRUCTOR:
		return true
	default:
		return false
	}
}
