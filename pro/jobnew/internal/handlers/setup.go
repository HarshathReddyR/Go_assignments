package handlers

import (
	"fmt"
	"jobnew/internal/auth"
	"jobnew/internal/middlewares"
	"jobnew/internal/models"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func API(a *auth.Auth, c *models.Conn) *gin.Engine {
	r := gin.New()
	m, err := middlewares.NewMid(a)
	ms := models.NewStore(c)
	h := handler{
		s: ms,
		a: a,
	}
	if err != nil {
		log.Panic().Msg("middlewares not set up")
	}
	r.Use(m.Log(), gin.Recovery())
	r.GET("/check", m.Authenticate(check))
	r.POST("/signup", h.Signup)
	r.POST("/login", h.Login)
	return r
}
func check(c *gin.Context) {
	time.Sleep(time.Second * 3)
	select {
	case <-c.Request.Context().Done():
		fmt.Println("user not there")
		return
	default:
		c.JSON(http.StatusOK, gin.H{"msg": "statusOk"})

	}

}
