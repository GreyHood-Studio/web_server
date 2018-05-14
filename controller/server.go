package controller

import (
	"net/http"
	"fmt"
	"github.com/GreyHood-Studio/server_util/error"
)

func AssginServerID(serverType string)  {

}

func CheckServerStatus(address string)  {
	res, err := http.Get(fmt.Sprintf("http://%s/ping", address))
	error.NoDeadError(err,"server ping error")
	fmt.Println(res)
}