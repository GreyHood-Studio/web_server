package controller

import (
	"github.com/GreyHood-Studio/server_util/database"
	"fmt"
)

func LoginUser(userID string, password string) (string, bool) {
	// select
	var nickname string
	err := database.Conns["user"].QueryRow("SELECT nickname FROM tb_user WHERE id=$1 AND password=$2",
		userID, password).Scan(&nickname)
	if err != nil {
		return "", false
	}
	fmt.Println("logined user: ", nickname)
	return nickname, true
}
