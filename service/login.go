package service

import "fmt"

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

type UserService struct{}

// func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
// 	if nil == global.GVA_DB {
// 		return nil, fmt.Errorf("db not init")
// 	}

// 	var user system.SysUser
// 	err = global.GVA_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
// 	if err == nil {
// 		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
// 			return nil, errors.New("密码错误")
// 		}
// 		MenuServiceApp.UserAuthorityDefaultRouter(&user)
// 	}
// 	return &user, err
// }

func (userService *UserService) Login() string {
	fmt.Println("now in service: login")
	return "login success"
}
