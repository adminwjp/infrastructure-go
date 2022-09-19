package goconfigs

import (
	"github.com/Unknwon/goconfig"
	"log"
)


func  LoadFile(file string, files ...string) *goconfig.ConfigFile {
	cfg, err := goconfig.LoadConfigFile(file, files...)
	if err != nil {
		log.Println("Load %s Fail", err.Error())
	}
	return cfg
}

type configUtil struct {

}
//===== *goconfig.ConfigFile is nil   =====

func (configUtil) GetStringValue(cfg *goconfig.ConfigFile, section string, key string, defValue string) (r string) {
	val, err := cfg.GetValue(section, key)
	if err != nil {
		return defValue
	}
	return val
}

func  (configUtil) GetIntValue(cfg *goconfig.ConfigFile, section string, key string, defValue int) (r int) {
	val, err := cfg.Int(section, key)
	if err != nil {
		return defValue
	}
	return val
}

func (configUtil) GetBoolValue(cfg *goconfig.ConfigFile, section string, key string, defValue bool) (r bool) {
	val, err := cfg.Bool(section, key)
	if err != nil {
		return defValue
	}
	return val
}

func (configUtil) GetFloat64Value(cfg *goconfig.ConfigFile, section string, key string, defValue float64) (r float64) {
	val, err := cfg.Float64(section, key)
	if err != nil {
		return defValue
	}
	return val
}

func  (configUtil) GetInt64Value(cfg *goconfig.ConfigFile, section string, key string, defValue int64) (r int64) {
	val, err := cfg.Int64(section, key)
	if err != nil {
		return defValue
	}
	return val
}

func (configUtil) GetArrayValue(cfg *goconfig.ConfigFile, section string, key string, delim string) (r []string) {
	return cfg.MustValueArray(section, key, delim)
}
