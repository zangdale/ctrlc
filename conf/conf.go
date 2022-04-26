package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Conf struct {
	WebDir string `json:"web_dir,omitempty"`
	Port   int    `json:"port,omitempty"`
}

func Init(homeDir string) *Conf {
	if homeDir == "" {
		return &Conf{}
	}
	b, err := ioutil.ReadFile(filepath.Join(homeDir, "config.json"))
	if err != nil {
		return &Conf{}
	}
	res := &Conf{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return &Conf{}
	}
	if res.Port < 0 {
		res.Port = 0
	}
	if res.WebDir != "" {
		fi, err := os.Stat(res.WebDir)
		if err != nil {
			res.WebDir = ""
		} else {
			if !fi.IsDir() {
				res.WebDir = ""
			}
		}
	}

	return res
}
