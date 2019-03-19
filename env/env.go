package env

import (
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// RegisterServices sets up all the web components.
func Load( defaultPath string ) {
	// Make SURE we use same location as executable for the config file
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// NOTE: Environment Variables must be Upper Case
	// Even if config file uses mixed case.

	// Automatically use env variables if supplied
	viper.AutomaticEnv()
	viper.AddConfigPath(filepath.Dir(ex)) // look in executable directory
	viper.AddConfigPath(".")              // alternatively, the Working directory

	if defaultPath !="" {
		viper.AddConfigPath(defaultPath) // search the default path
	}

	// These are used by tests, which run in sub-directory
	// so need to look for env.json in higher directory
	viper.AddConfigPath("..\\..\\..\\..\\..\\")
	viper.AddConfigPath("..\\..\\..\\..\\")
	viper.AddConfigPath("..\\..\\..\\")
	//viper.AddConfigPath("..\\..\\")
	viper.SetConfigName("env") // name of config file (without extension)
	err = viper.ReadInConfig() // Find and read the config file
	// Note we do not panic if this did not exist, we might be able to
	// get all our configuration values from environment variables.
	if err == nil {
		log.Info("Config File Used: "+viper.ConfigFileUsed())

		// If file does exist, set up a watcher
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Info("Config file changed:")
		})
	}

}
