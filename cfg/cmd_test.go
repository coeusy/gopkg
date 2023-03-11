package cfg

import (
	"flag"
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestInitConfigFromCMD(t *testing.T) {
	InitCMD()
	fmt.Println(flag.Set("cfg.path", "../conf"))
	fmt.Println(flag.Set("cfg.files", "runtime"))
	InitConfigFromCMD()
	fmt.Println(jsoniter.MarshalToString(cmdManager.AllSettings()))
}

func TestGetKafka(t *testing.T) {
	InitCMD()
	fmt.Println(flag.Set("cfg.path", "../conf"))
	fmt.Println(flag.Set("cfg.files", "runtime"))
	InitConfigFromCMD()
	fmt.Println(GetKafka())
}

func TestGetConfig(t *testing.T) {
	InitCMD()
	fmt.Println(flag.Set("cfg.path", "../conf"))
	fmt.Println(flag.Set("cfg.files", "runtime"))
	InitConfigFromCMD()
	fmt.Println(GetConfig())
}

func TestGetElasticSearch(t *testing.T) {
	InitCMD()
	fmt.Println(flag.Set("cfg.path", "../conf"))
	fmt.Println(flag.Set("cfg.files", "runtime"))
	InitConfigFromCMD()
	fmt.Println(GetES())
}

func TestGetInfluxDB(t *testing.T) {
	InitCMD()
	fmt.Println(flag.Set("cfg.path", "../conf"))
	fmt.Println(flag.Set("cfg.files", "runtime"))
	InitConfigFromCMD()
	fmt.Println(GetInfluxDB())
}
