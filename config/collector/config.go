package parser

import (
	"path/filepath"

	"github.com/reucot/parser/config"

	cfg "github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
	"github.com/gookit/ini/v2/dotenv"
)

type Config struct {
	DB config.DB `json:"db"`
}

var instance Config

func Load() error {
	cfg.WithOptions(cfg.ParseEnv, func(o *cfg.Options) {
		o.DecoderConfig.TagName = "json"
	})

	cfg.AddDriver(json.Driver)

	filename, err := filepath.Abs(filepath.Join("./", ".env"))
	if err != nil {
		return err
	}

	if err := dotenv.LoadFiles(filename); err != nil {
		return err
	}

	filename, err = filepath.Abs(filepath.Join("./", "cmd/parser/config.json"))
	if err != nil {
		return err
	}

	if err := cfg.LoadFiles(filename); err != nil {
		return err
	}

	if err := cfg.BindStruct("", &instance); err != nil {
		return err
	}

	return nil
}

func Get() Config {
	return instance
}
