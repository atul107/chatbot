package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		AWS  `yaml:"aws"`
		GPT  `yaml:"gpt"`
	}

	HTTP struct {
		Port      string `yaml:"port" env:"HTTP_PORT"`
		Algorithm string `yaml:"algorithm" env:"HTTP_ALGO"`
		Cert      string `yaml:"cert" env:"HTTP_CERT"`
		Key       string `yaml:"key" env:"HTTP_KEY"`
	}

	App struct {
		Name    string `yaml:"name" env:"APP_NAME"`
		Version string `yaml:"version" env:"APP_VERSION"`
		ENV     string `yaml:"env" env:"ENV"`
	}

	Log struct {
		Level string `yaml:"log_level" env:"LOG_LEVEL"`
	}

	AWS struct {
		AccessKeyID     string `yaml:"access_key_id" env:"ACCESS_KEY_ID"`
		AccessKeySecret string `yaml:"access_key_secret" env:"ACCESS_KEY_SECRET"`
		Region          string `yaml:"region" env:"REGION"`
		BucketName      string `yaml:"bucket_name" env:"BUCKET_NAME"`
		UploadTimeout   int    `yaml:"upload_timeout" env:"UPLOAD_TIMEOUT"`
	}

	GPT struct {
		OpenAIKey   string  `yaml:"port" env:"OPENAI_KEY"`
		Model       string  `yaml:"port" env:"GPT_MODEL"`
		Temperature float64 `yaml:"port" env:"GPT_TEMPERATURE"`
		Token       int     `yaml:"port" env:"GPT_MAX_TOKEN"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf(".env file not found or error loading it: %v", err)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf(".env file not found or error loading it: %v", err)
	}
	return cfg, nil
}
