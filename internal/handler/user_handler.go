package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/todo-app/internal/config"
	"github.com/todo-app/internal/service"
	"net/http"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) GithubLogin(c *gin.Context) {
	cfg := config.AppCfg.GithubConfig
	url := "https://github.com/login/oauth/authorize" +
		"?client_id=" + cfg.ClientId +
		"&redirect_uri=" + cfg.RedirectUri +
		"&scope=user:email"
	c.Redirect(http.StatusFound, url)
}

func (h *UserHandler) GithubCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.Error(errors.New("missing query param of code"))
		return
	}
	token, err := h.svc.ProcessGithubCallback(code)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
