package config

import (
	"github.com/spf13/viper"
	"github.com/tevino/log"
	"os"
	"path/filepath"
	"strings"
)

type Account struct {
	ID   string `yaml:"id"`
	Seed string `yaml:"seed"`
}

var FaucetAccount Account

func LoadConfig(configPath string) {
	curdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if len(configPath) == 0 {
		dirs := strings.Split(curdir, "/")
		configPath = "/" + filepath.Join(dirs[0:len(dirs)-1]...)
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("config file error: %s\n%s\n", err, curdir)
	}
	FaucetAccountMap := viper.GetStringMapString("faucetaccount")
	FaucetAccount.ID = FaucetAccountMap["id"]
	FaucetAccount.Seed = FaucetAccountMap["seed"]
	log.Info(FaucetAccount)
}
