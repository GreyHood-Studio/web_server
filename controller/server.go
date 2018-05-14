package controller

import (
	"github.com/GreyHood-Studio/web_server/router"
	"net/http"
	"fmt"
	"github.com/GreyHood-Studio/server_util/error"
)

func assginServerID(serverType string)  {

}

func checkServerStatus()  {
	for _, gameServer := range router.GameServers {
		res, err := http.Get(fmt.Sprintf("http://%s/ping", gameServer.Address))
		error.NoDeadError(err,"server ping error")
		fmt.Println(res)
	}
}