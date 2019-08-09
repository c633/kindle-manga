package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var configDir string

func init() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	configDir = filepath.Join(usr.HomeDir, ".config", "kindle-manga")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		panic(err)
	}
}

type manga struct {
	Name string
	Chap int
	URL  []string
}
type collection struct {
	Manga []manga
}

func main() {
	var coll collection
	configPath := filepath.Join(configDir, "config.toml")
	if _, err := toml.DecodeFile(configPath, &coll); err != nil {
		log.Fatalln(err)
	}

	bot := newRobot(coll, configPath)
	bot.run()
	bot.save()
}
