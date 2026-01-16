package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var AppCfg *AppConfig

func Load() error {
	cfgFile := resolveConfigPath()
	envData, err := os.ReadFile(cfgFile)
	if err != nil {
		return err
	}
	_cfg := AppConfig{}
	if err := yaml.Unmarshal(envData, &_cfg); err != nil {
		return err
	}

	AppCfg = &_cfg
	return nil
}

func resolveConfigPath() string {
	var filename = "config.example.yaml"
	paths := []string{
		filepath.Join("config", filename),
		// docker container
		filepath.Join("/app/config", filename),
	}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}

	return paths[0]
}
