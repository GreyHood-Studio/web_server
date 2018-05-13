package controller

import (
	"fmt"
	"github.com/GreyHood-Studio/server_util/database"
	"github.com/GreyHood-Studio/server_util/error"
)

func SignUpAccount(userID string, password string, nickname string) (){
	// insert
	var uid int
	err := database.Conns["user"].QueryRow("INSERT INTO tb_user(id, password, nickname) VALUES($1,$2,$3) returning uid;", userID, password, nickname).Scan(&uid)
	error.NoDeadError(err, "insert user error")
	fmt.Println("total user count =", uid)
}

func ChangePassword(userID string, password string) (){
	// update
	fmt.Println("# Updating")
	stmt, err := database.Conns["user"].Prepare("update tb_user set password=$1 where id=$2")
	error.NoDeadError(err, "prepare error")

	res, err := stmt.Exec(password, userID)
	error.NoDeadError(err, "update error")

	affect, err := res.RowsAffected()
	error.NoDeadError(err, "rows affected error")

	fmt.Println(affect, "rows changed")
}

func DeleteAccount(userID string)  {
	
}

func VerifyAccount(userID string, password string)  {
	// select
}