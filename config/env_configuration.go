package config

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host     string `env:"host" validate:"required"`
	Port     string `env:"port" validate:"required"`
	Database string `env:"database" validate:"required"`
	User     string `env:"user" validate:"required"`
	Password string `env:"password" validate:"required"`
}
type ConnectionStrings struct {
	Mysql MysqlConfig `env:"mysql" validate:"required"`
}
type envConfiguration struct {
	Env               string            `env:"env" validate:"required"`
	Name              string            `env:"name" validate:"required"`
	Host              string            `env:"host" validate:"required"`
	Port              string            `env:"port" validate:"required"`
	ConnectionStrings ConnectionStrings `env:"connectionstrings" validate:"required"`
}

var EnvConfig envConfiguration

func bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("env")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			_ = viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
func InitialEnvConfiguration() (err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("app")
	bindEnvs(EnvConfig)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		default:
			panic(fmt.Errorf("fatal error loading config file: %s", err))
		case viper.ConfigFileNotFoundError:
			log.Println("No config file found.Using defaults and environment variables")
		}
	}

	if err := viper.Unmarshal(&EnvConfig); err != nil {
		log.Fatalf("unable to decode into config struct, %v", err)
	}

	validate := validator.New()
	err = validate.Struct(&EnvConfig)

	return
}
