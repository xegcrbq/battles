package env

import (
	"github.com/joho/godotenv"
	"sync"
)

var once sync.Once

func InitEnv() error {
	var err error
	once.Do(func() {
		envstr := ".env"
		for i := 0; i < 5; i++ {
			err = godotenv.Load(envstr)
			if err == nil {
				return
			} else {
				envstr = "..\\" + envstr
			}
		}
	})

	return err
}

//func GoDotEnvVariable(key string) string {
//
//	// load .env file
//	err := godotenv.Load(".env")
//
//	if err != nil {
//		log.Fatalf("Error loading .env file")
//	}
//
//	return os.Getenv(key)
//}
