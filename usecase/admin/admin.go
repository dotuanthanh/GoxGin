package admin

func isExisted(phone string) bool {
	//sql := sqlbuilder.Select("count(*)").From(Account_Table)
	return false
}

func CreateAccount(phone string, password string, name string, note string) error {

	return nil
}

func UpdateAccount(phone string, password string, note string) error {

	return nil
}

func CheckLogin(email string, password string) *LoginResponse {
	//TODO
	return &LoginResponse{
		Email: email,
		Token: "token",
	}
}

func ChangePassword() {

}
