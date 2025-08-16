package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`

	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

func LoadConfig(mode ...string) (*Config, error) {
	var cfg Config

	// 1. Tentukan mode
	envMode := "local"
	if len(mode) > 0 && mode[0] != "" {
		envMode = strings.ToLower(mode[0])
	} else if val := os.Getenv("APP_MODE"); val != "" {
		envMode = strings.ToLower(val)
	}

	// Pastikan ada ekstensi .yaml
	fileName := envMode
	if filepath.Ext(fileName) == "" {
		fileName += ".yaml"
	}

	// 2. Cari working directory
	var dir string
	var err error
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		dir, err = os.Getwd()
	} else {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	}
	if err != nil {
		return nil, fmt.Errorf("failed to determine working directory: %w", err)
	}

	// 3. Bangun path file config
	configPath := filepath.Join(dir, "env", fileName)

	// 4. Baca file YAML
	f, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %s: %w", configPath, err)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

func InitDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}
