package repository

func isExisted(phone string) bool {
	//query :="SELECT COUNT(*) FROM Account WHERE phone = ?"
	return false
}

//ID        int64  `json:"ID"`
//Phone     string `json:"phone"`
//Password  string `json:"password"`
//Name      string `json:"name"`
//Note      string `json:"note"`
//Role      string `json:"role"`
//IsDeleted bool   `json:"isDeleted"`
//UpdateAt  string `json:"updateAt"`

func createAccount(phone string, password string, name string, note string) error {

	return nil
}
func updateAccount(phone string, password string, note string) error {
	return nil
}
func checkLogin(phone string, password string) {

}
