package service

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/todo-app/internal/config"
	"github.com/todo-app/internal/dto"
	"github.com/todo-app/internal/models"
	"github.com/todo-app/internal/repository"
	"github.com/todo-app/pkg/util"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		repo: repository.NewUserRepository(db),
	}
}

func (s *UserService) ProcessGithubCallback(code string) (string, error) {
	accessToken, err := exchangeCodeForToken(code)
	if err != nil {
		return "", err
	}

	githubUser, err := fetchGithubUser(accessToken)
	if err != nil {
		return "", err
	}

	providerId := strconv.FormatInt(githubUser.ID, 10)
	user, err := s.repo.FindByProviderId(dto.ProviderGithub, providerId)
	var userIdStr string

	if err == nil {
		userIdStr = user.ID.String()
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userDb := models.User{
				ID:         uuid.New(),
				Provider:   dto.ProviderGithub,
				ProviderID: providerId,
				Username:   githubUser.Login,
			}
			err = s.repo.Create(&userDb)
			userIdStr = userDb.ID.String()
		} else {
			return "", err
		}
	}

	token, err := util.GenerateToken(userIdStr)
	return token, err
}

func exchangeCodeForToken(code string) (string, error) {
	cfg := config.AppCfg.GithubConfig

	data := url.Values{}
	data.Set("client_id", cfg.ClientId)
	data.Set("client_secret", cfg.ClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", cfg.RedirectUri)

	req, _ := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		strings.NewReader(data.Encode()),
	)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		AccessToken string `json:"access_token"`
	}

	json.NewDecoder(resp.Body).Decode(&result)
	return result.AccessToken, nil
}

func fetchGithubUser(token string) (*dto.GithubUser, error) {
	req, _ := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var user dto.GithubUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	return &user, nil
}
