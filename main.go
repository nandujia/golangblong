package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"test-blog/models"
)

// 定一个index数据的结构体，里面有两个字符数类型的参数，定义json格式

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	//1. 拿到当前的路径
	path, _ := os.Getwd()
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t, _ = t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)

	//页面上涉及到的所有的数据，必须有定义
	var hr = &models.HomeResponse{}
	t.Execute(w, hr)

}

func main() {
	//程序主入口
	server := http.Server{
		Addr: "127.0.0.1:8888",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err, "本地启动失败")
	}
}
