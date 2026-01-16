package dto

const ProviderGithub = "github"

type GithubUser struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}
