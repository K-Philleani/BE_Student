package model

import "student/connection"

// Stitches数据库 account表model

type Account struct {
	Id 			int 	`json:"id"`
	UserAccount string 	`json:"user_account"`
	UserPwd 	string 	`json:"user_pwd"`
	UserPhone 	string 	`json:"user_phone"`
}

func (a *Account) CreateAccount() {

}

func (a *Account) GetAccountAll() (accounts []Account, err error){
	err = connection.DB.Table("account").Find(&accounts).Error
	if err != nil {
		return
	}
	return
}