package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kirsle/configdir"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	gotoml "github.com/pelletier/go-toml/v2"
	"golang.design/x/hotkey"
)

type Config struct {
	HotKey *HotKey `koanf:"hotkey" toml:"hotkey,dive"`
}

type HotKey struct {
	Modifiers []hotkey.Modifier `koanf:"modifiers" toml:"modifiers"`
	Key       hotkey.Key        `koanf:"key" toml:"key"`
}

var k = koanf.New(".")

var DEFAULT_CONFIG = &Config{
	HotKey: &HotKey{
		Modifiers: []hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift},
		Key:       hotkey.KeyM,
	},
}

func Load(appName string) (*Config, error) {
	// ensure there aren't any spaces in the app name
	appName = strings.ReplaceAll(appName, " ", "_")
	// get the local config path for this app
	configPath := configdir.LocalConfig(appName)
	// ensure the path exists
	err := configdir.MakePath(configPath)
	if err != nil {
		return nil, err
	}
	// generate filename
	fileName := fmt.Sprintf("%s.toml", appName)
	// join to a full path
	configFile := filepath.Join(configPath, fileName)
	// determine if the file exists
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		// create a new default config file if it doesn't
		fh, err := os.Create(configFile)
		if err != nil {
			return nil, err
		}
		defer fh.Close()
		encoder := gotoml.NewEncoder(fh)
		err = encoder.Encode(DEFAULT_CONFIG)
		if err != nil {
			return nil, err
		}
		fh.Close()
		// then return the default config, no need to re-load it from the file
		return DEFAULT_CONFIG, nil
	}
	// load the config file into koanf
	if err = k.Load(file.Provider(configFile), toml.Parser()); err != nil {
		return nil, err
	}
	// then unmarshal
	var conf Config
	err = k.Unmarshal("", &conf)
	if err != nil {
		return nil, err
	}
	// and return the unmarshaled config
	return &conf, nil
}
