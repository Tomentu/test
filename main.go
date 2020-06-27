package main

import (
	"fmt"
	"html/template"
	"net/http"
	"spider"
)

func main() {
	fmt.Println("hi")
	http.HandleFunc("/create", create) //设置访问的路由
	http.HandleFunc("/download",download)
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		//log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("Run app")
}
func create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form["url"])<0 {
		fmt.Fprint(w,"erro")
		return
	}
	url := r.Form["url"][0]
	fmt.Println(url)
	spider.Create(&w,url)
}
func download(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	if len(r.Form["id"])<0 {
		return
	}
	id := r.Form["id"][0]
	fmt.Println(id);
	spider.Download2loacl(id,&w)
}
func home(w http.ResponseWriter, r *http.Request){
	tem,err := template.ParseFiles("./html/index.html")
	if err != nil{
		fmt.Println("读取文件失败,err",err)
		return
	}

	// 利用给定数据渲染模板，并将结果写入w
	tem.Execute(w,"hi")
}
