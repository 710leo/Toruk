# Toruk
### 一个简单的go web框架

主要由下面几个第三方库组成
* 路由：github.com/gorilla/mux
* 渲染模板：github.com/unrolled/render
* Cookie封装：github.com/gorilla/securecookie
* 中间件：github.com/codegangsta/negroni
* 上下文传输：github.com/gorilla/context

整个目录结构如下
├── cfg.example.json
├── g
│   ├── config.go
│   └── const.go
├── control
├── handler
│   └── home_handler.go
├── http
│   ├── cookie
│   │   └── cookie.go
│   ├── errors
│   │   └── errors.go
│   ├── middleware
│   │   └── recovery.go
│   ├── param
│   │   └── param.go
│   ├── render
│   │   └── render.go
│   ├── http.go
│   └── routes.go
├── main.go
├── static
└── views
│   └──  home
│   │   └── index.html

