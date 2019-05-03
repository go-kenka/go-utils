package server

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	configEnvFileName = "config.env"
)

// Start 加载环境变量和配置文件
func Start() {
	// 加载环境变量
	InitConfigEnv()
}

// InitConfigEnv 初始化env配置
func InitConfigEnv() {
	file := fmt.Sprintf("%v/%v", getCwd(), configEnvFileName)
	if err := godotenv.Load(file); err != nil {
		log.Fatalf("ERROR LOADING DOTENV %v at path %v", err, file)
	}

	log.Print("config.env initialize!")
}

func getCwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// if path is root, return empty instead
	if pwd == "/" {
		pwd = ""
	}

	return pwd
}
