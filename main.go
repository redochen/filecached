package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/hprose/hprose-go/hprose"
	"github.com/redochen/filecached/config"
	"github.com/redochen/filecached/core"
)

//主函数
func main() {
	fmt.Println("[main] starting server...")

	svc := hprose.NewHttpService()
	svc.AddFunction("get", core.GetCache)
	svc.AddFunction("post", core.SetCache)

	fmt.Printf("[main] server started! listening on port: %s\n", strconv.Itoa(config.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), svc)
}
