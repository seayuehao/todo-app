package config

type AppConfig struct {
	ServerConfig   ServerConfig   `yaml:"server_config"`
	JwtConfig      JwtConfig      `yaml:"jwt_config"`
	DbConfig       DbConfig       `yaml:"db_config"`
	GoogleConfig   GoogleConfig   `yaml:"google_config"`
	GithubConfig   GithubConfig   `yaml:"github_config"`
	FacebookConfig FacebookConfig `yaml:"facebook_config"`
}

type ServerConfig struct {
	Name    string `yaml:"name"`
	Port    int    `yaml:"port"`
	Version string `yaml:"version"`
	Env     string `yaml:"env"`
}

type JwtConfig struct {
	Secret  string `yaml:"secret"`
	Expires int    `yaml:"expires"`
}

type DbConfig struct {
	Driver string `yaml:"driver"`
	Dsn    string `yaml:"dsn"`
}

type GoogleConfig struct {
	RedirectUri  string `yaml:"redirect_uri"`
	ClientId     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

type GithubConfig struct {
	RedirectUri  string `yaml:"redirect_uri"`
	ClientId     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

type FacebookConfig struct {
	RedirectUri  string `yaml:"redirect_uri"`
	ClientId     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}
