package config

import (
	"github.com/spf13/viper"
	"log"
	"reflect"
	"strings"
)


type Config struct {
	ParenPath   string
	Resource    string
	Env         string
	Application string
}

// Load function will read config from environment or config file.
func LoadConfig( parentPath string, container string, env string, fileNames ...string) ServerConfig {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.SetConfigType("yaml")

	for _, fileName := range fileNames {
		viper.SetConfigName(fileName)
	}

	viper.AddConfigPath("./" + container + "/")
	if len(parentPath) > 0 {
		viper.AddConfigPath("./" + parentPath + "/" + container + "/")
	}

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Println("config file not found")
		default:
			panic(err)
		}
	}
	if len(env) > 0 {
		env2 := strings.ToLower(env)
		for _, fileName2 := range fileNames {
			name := fileName2 + "-" + env2
			viper.SetConfigName(name)
			viper.MergeInConfig()
		}
	}
	var c ServerConfig
	bindEnvs(c)
	viper.Unmarshal(&c)
	return c
}

// bindEnvs function will bind ymal file to struc model
func bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}