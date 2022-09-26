package system

import (
	"log"
	"net/http"
	"server/model/system/request"

	"github.com/gin-gonic/gin"
)

type BaseApi struct {
}

func (b *BaseApi) Register(c *gin.Context) {
	var r request.Register
	_ = c.ShouldBindJSON(&r)
	log.Println(r.Username, r.Password, r.NickName)
	c.String(http.StatusOK, r.Username)
}
