package main

import (
	"os/exec"

	"github.com/rs/zerolog/log"
)

func backup(homeDir, dbPath string) {
	c := exec.Command("sh", "backup.sh", dbPath, homeDir)
	c.Dir = homeDir
	log.Info().Msg(c.String())
	err := c.Run()
	if err != nil {
		log.Error().Err(err).Msg("ctrlc backup")
		return
	}
	log.Info().Msg("ctrlc backup ok ...")
}
