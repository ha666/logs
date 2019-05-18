package logs

import (
	"github.com/ha666/golibs"
	"log"
)

var (
	idNode *golibs.Node
	err    error
)

func init() {
	idNode, err = golibs.NewNode(golibs.Unix() % 1024)
	if err != nil {
		log.Fatalf("【init】加载节点出错:%s", err.Error())
		return
	}
}
