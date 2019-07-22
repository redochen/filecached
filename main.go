package main

import (
	"fmt"
	"github.com/hprose/hprose-go/hprose"
	cfg "github.com/redochen/filecached/config"
	. "github.com/redochen/filecached/core"
	"net/http"
	"strconv"
)

//主函数
func main() {
	fmt.Println("[main] starting server...")

	svc := hprose.NewHttpService()
	svc.AddFunction("get", GetCache)
	svc.AddFunction("post", SetCache)

	fmt.Printf("[main] server started! listening on port: %s\n", strconv.Itoa(cfg.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), svc)
}
