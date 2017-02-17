# Toruk
### go web 开发脚手架

主要由下面几个第三方库集成，简单、灵活，可以快速开发web项目

* 路由：github.com/gorilla/mux
* 渲染模板：github.com/unrolled/render
* Cookie封装：github.com/gorilla/securecookie
* 中间件：github.com/codegangsta/negroni
* 上下文传输：github.com/gorilla/context 

### 文件结构
<img src="http://x2know.qiniudn.com/toruk2.png" width = "300" style="margin-left:0px" alt="图片描述" align=center />


### 使用方法
通过修改http/ruote.go 将数据传到对应的handler中，然后在通过render将数据返回回去。
ruote.go 

    func configConfRoutes(r *mux.Router) {
		r.HandleFunc("/home", handler.HomeIndex).Methods("GET")
	}

handler.go

    func HomeIndex(w http.ResponseWriter, r *http.Request) {
		render.HTML(r, w, "home/index")
	}

### 初始化

    # set $GOPATH and $GOROOT
    # 比如你的项目名称叫做 awosome
    cd $GOPATH/src
    git clone https://github.com/710leo/Toruk.git
    mv Toruk awosome
    cd awosome
    ./init awosome
    go get ./...

### 编译&运行
    ./control build
    ./control start

### 答疑
交流QQ群：173502733
