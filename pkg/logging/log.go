package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

var Logger *logrus.Logger

func init() {
	// 設定 json 格式
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 名稱
	filename := path.Join(os.Getenv("LOG_PATH"), os.Getenv("LOG_FILENAME"))

	// write file
	src, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println(filename, " not exists ready to create")
		src, err = os.Create(filename)
		if err != nil {
			fmt.Println("create log failed: ", err)
		}
	}

	// 高於 warning 以上才會紀錄
	logrus.SetLevel(logrus.DebugLevel)

	Logger = logrus.New()

	// 輸出至指定檔案
	Logger.Out = src
}
