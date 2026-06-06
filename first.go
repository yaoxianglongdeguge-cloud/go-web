package main

import (
	"fmt"
	"net/http"
	"os"
)

func open_file(s string)string{
	f,err:=os.ReadFile(s)
	if err!=nil{
		return "打开文件错误"+err.Error()
	}
	return string(f)
}

func Handleget(w http.ResponseWriter, r *http.Request) {
	r_url:=r.URL.Path[1:]
	fmt.Println(r_url)
	f:=open_file(r_url)
	w.Write([]byte(f))
}

func main() {

	http.HandleFunc("/", Handleget)
	http.ListenAndServe("172.26.27.214:3412", nil)
}
