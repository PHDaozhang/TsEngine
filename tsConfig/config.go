package tsConfig

import (
	_ "fmt"

	"github.com/astaxie/beego/config"
)

func ReadFile(filename, fileType string) (conf config.Configer, err error) {
	return config.NewConfig(fileType, filename)
}
