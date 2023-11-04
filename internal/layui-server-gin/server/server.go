package server

import (
	"github.com/gin-gonic/gin"
)

// Server is used for api
type Server struct {
	Engine *gin.Engine
}

func New() (*Server, error) {
	s := new(Server)
	err := InitAPIServer(s)
	// 不需要打印，因为我们已经把err向上反馈了
	//if err != nil {
	//	zap.L().Error("Init Server failed", zap.Error(err))
	//}
	return s, err
}
