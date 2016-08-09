# Toruk
### go web 开发脚手架

主要由下面几个第三方库组成

* 路由：github.com/gorilla/mux
* 渲染模板：github.com/unrolled/render
* Cookie封装：github.com/gorilla/securecookie
* 中间件：github.com/codegangsta/negroni
* 上下文传输：github.com/gorilla/context 

### 初始化

    # set $GOPATH and $GOROOT
    # 比如你的项目名称叫做 awesome
    cd $GOPATH/src
    git clone https://github.com/710leo/Toruk.git
    mv Toruk awesome
    cd awesome
    ./init awesome
    go get ./...