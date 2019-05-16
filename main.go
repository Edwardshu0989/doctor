package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"./libs/config"
	"./models"
	"github.com/gin-gonic/gin"
)

var (
	cfg = flag.String("c", "config/ui.toml", "config file, relative path")
)

func init() {
	flag.Parse()

}

func main() {
	if strings.HasPrefix(*cfg, "/") == false {
		*cfg = path.Join(os.Getenv("WORKDIR"), *cfg)
	}

	// panic(*cfg)
	if err := config.SetFileAndLoad(*cfg); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(config.Get().Db.Server)
	models.Init()
	var handler *gin.Engine
	var err error
	handler, err = Newserver()
	if err != nil {
		panic("启动服务失败")
		os.Exit(-1)
	}
	srv := &http.Server{Addr: ":8888", Handler: handler} //绑定端口
	if err = srv.ListenAndServe(); err != nil {
		fmt.Printf("UiServer stopped: %s\n", err)
		os.Exit(-1)
	}
}
