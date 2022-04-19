package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

/**
 * 日志服务
 */
var log *stlog.Logger

type fileLog string

// Write 集成io write
func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

// Run 服务启动
func Run(destination string) {
	// 创建一个log服务
	log = stlog.New(fileLog(destination), "go", stlog.LstdFlags)
}

// RegisterHandlers 注册HTTP方法，用于日志数据接收
func RegisterHandlers() {
	http.HandleFunc("/log", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(request.Body)
			if err != nil || len(msg) == 0 {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

// message 服务端日志打印
func write(message string) {
	log.Printf("%v\n", message)
}
