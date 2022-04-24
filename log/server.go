package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

/*
 * 可运行的日志服务
 *
 */
var log *stlog.Logger

type fileLog string

// Write 集成io write
func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

// Run 服务启动
func Run(destination string) {
	// 创建一个log服务
	log = stlog.New(fileLog(destination), "[go]- ", stlog.LstdFlags)
}

// RegisterHandlers 注册HTTP方法，用于日志数据接收
func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

// write 服务端日志打印
func write(message string) {
	log.Printf("%v\n", message)
}
