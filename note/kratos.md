# kratos

- Refence 

  [Kratos (bilibili)](https://space.bilibili.com/1885628842), 

  [Kratos org](https://go-kratos.dev/), [Kratos (github)](https://github.com/go-kratos/kratos), [Kratos WX](https://mp.weixin.qq.com/s/Wm1pHZAbybHV6BLqDWEPVA), [Kratos engineering](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAwNzgwMzY2Ng==&action=getalbum&album_id=1816622127303753730&scene=126&sessionid=-1820420947&uin=&key=&devicetype=Windows+11+x64&version=63090a13&lang=zh_CN&ascene=0), 

  [gothinkster / realworld](https://github.com/gothinkster/realworld), [realworld docs](https://realworld-docs.netlify.app/docs/specs/backend-specs/introduction/),





- Overview

  供前端 (浏览器) 使用的 web应用

  1. 框架的基本情况
  2. 初始化项目
  3. 定义API
  4. 填充业务逻辑
  5. 中间件使用 (登录时获取用户信息 鉴权)
  6. 自定义接口的返回格式

- 本教程不包含

  微服务相关的内容 (服务注册发现 ...)
  
  



- 项目开发

  [OpenAPI Swagger 使用](https://go-kratos.dev/docs/guide/openapi/), [Swagger UI](https://github.com/swagger-api/swagger-ui)

  



## 项目创建与介绍

### 初始化项目

- 环境准备

  [go](https://golang.org/dl/), (`GO111MODULE`, [`GOPROXY`](https://goproxy.cn/))

  [protoc](https://github.com/protocolbuffers/protobuf) (代码生成), [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go) (插件)

  ```bash
  go version
  go env
  go env -w GOPATH=D:\devenv\go\
  go env -w GOPROXY=https://goproxy.cn,direct
  
  
  go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
  kratos --version
  
  ```

- 脚手架 kratos  ([CLI工具](https://go-kratos.dev/docs/getting-started/usage))

  通过模板快速创建项目

  快速创建与生成 protoc 文件

  使用开发过程中常用的命令

  极大提高开发效率，减轻心智负担

  ```bash
  # 用脚手架拉取项目
  kratos new kratos-realworld
  kratos new kratos-realworld -r https://gitee.com/go-kratos/kratos-layout.git  # 可国内 可自定义
  # 也可以通过环境变量指定源
  KRATOS_LAYOUT_REPO=xxx-layout.git
  kratos new kratos-realworld
  
  
  cd kratos-realworld && code .
  make help
  make init 
  make api  # 现在不需要执行
  
  go mod tidy  
  go mod download
  
  
  cd kratos-realworld
  go generate ./...
  go build -o ./bin/ ./...
  ./bin/kratos-realworld -conf ./configs
  
  
  # 运行项目
  kratos run
  # http://localhost:8000/helloworld/eric
  
  ```

- 依赖注入 [wire](https://github.com/google/wire)

  ```bash
  go install github.com/google/wire/cmd/wire@latest
  
  ```
  
- 项目结构  ([kratos docs](https://go-kratos.dev/docs/intro/layout), [kratos-layout](https://github.com/go-kratos/kratos-layout))

  ```
    .
  ├── Dockerfile  
  ├── LICENSE
  ├── Makefile  
  ├── README.md
  ├── api // 下面维护了微服务使用的proto文件以及根据它们所生成的go文件
  │   └── helloworld
  │       └── v1
  │           ├── error_reason.pb.go  # 编写.proto .go是生成的
  │           ├── error_reason.proto
  │           ├── error_reason.swagger.json
  │           ├── greeter.pb.go
  │           ├── greeter.proto
  │           ├── greeter.swagger.json
  │           ├── greeter_grpc.pb.go
  │           └── greeter_http.pb.go
  ├── cmd  // 整个项目启动的入口文件
  │   └── server
  │       ├── main.go
  │       ├── wire.go  // 我们使用wire来维护依赖注入
  │       └── wire_gen.go
  ├── configs  // 这里通常维护一些本地调试用的样例配置文件
  │   └── config.yaml
  ├── generate.go
  ├── go.mod
  ├── go.sum
  ├── internal  // 该服务所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
  │   ├── biz   // 业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，而 repo 接口在这里定义，使用依赖倒置的原则。
  │   │   ├── README.md
  │   │   ├── biz.go
  │   │   └── greeter.go
  │   ├── conf  // 内部使用的config的结构定义，使用proto格式生成
  │   │   ├── conf.pb.go
  │   │   └── conf.proto
  │   ├── data  // 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。我们可能会把 data 与 dao 混淆在一起，data 偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD 的 infra层。
  │   │   ├── README.md
  │   │   ├── data.go
  │   │   └── greeter.go
  │   ├── server  // http和grpc实例的创建和配置
  │   │   ├── grpc.go
  │   │   ├── http.go
  │   │   └── server.go
  │   └── service  // 实现了 api 定义的服务层，类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑
  │       ├── README.md
  │       ├── greeter.go
  │       └── service.go
  └── third_party  // api 依赖的第三方proto
      ├── README.md
      ├── google
      │   └── api
      │       ├── annotations.proto
      │       ├── http.proto
      │       └── httpbody.proto
      └── validate
          ├── README.md
          └── validate.proto
  ```

  ![Snipaste_2024-05-20_07-42-31](res/Snipaste_2024-05-20_07-42-31.png)
  
  



### 项目设计

- Endpoints

  [Authentication Header:](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#authentication-header), [Authentication:](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#authentication), [Registration:](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#registration)

  [Get Current User](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#get-current-user), [Update User](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#update-user), [Get Profile](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#get-profile), [Follow user](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#follow-user), [Unfollow user](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#unfollow-user)

  [List Articles](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#list-articles), [Feed Articles](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#feed-articles), [Get Article](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#get-article), [Create Article](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#create-article), [Update Article](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#update-article), [Delete Article](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#delete-article)

  [Add Comments to an Article](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#add-comments-to-an-article), [Get Comments from an Article](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#get-comments-from-an-article), [Delete Comment](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#delete-comment)

  [Favorite Article](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#favorite-article), [Unfavorite Article](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#unfavorite-article), [Get Tags](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints#get-tags)

- API Response format

  JSON Objects returned by API:

  [Users (for authentication)](https://realworld-docs.netlify.app/docs/specs/backend-specs/api-response-format#users-for-authentication), [Profile](https://realworld-docs.netlify.app/docs/specs/backend-specs/api-response-format#profile), [Single Article](https://realworld-docs.netlify.app/docs/specs/backend-specs/api-response-format#single-article), [Multiple Articles](https://realworld-docs.netlify.app/docs/specs/backend-specs/api-response-format#multiple-articles)

  [Single Comment](https://realworld-docs.netlify.app/docs/specs/backend-specs/api-response-format#single-comment), [Multiple Comments](https://realworld-docs.netlify.app/docs/specs/backend-specs/api-response-format#multiple-comments), [List of Tags](https://realworld-docs.netlify.app/docs/specs/backend-specs/api-response-format#list-of-tags)

- Error Handling

- CORS (跨域问题)

- Postman

  [Postman collection](https://github.com/gothinkster/realworld/blob/master/api/Conduit.postman_collection.json) 

- Tests





## API定义与生成

### 改名字

- 改名字

  `api/realworld/v1/` (realworld.proto, error_reason.proto)

  模板名称是 helloworld 需要替换 (体力活)

  代码生成器

  ```bash
  make api  # win 有问题
  protoc --proto_path=./api \
         --proto_path=./third_party \
         --go_out=paths=source_relative:./api \
         --go-http_out=paths=source_relative:./api \
         --go-grpc_out=paths=source_relative:./api \
         --openapi_out=fq_schema_naming=true,default_response=false:. \
         ./api/realworld/v1/realworld.proto ./api/realworld/v1/error_reason.proto 
         
  cd cmd/kratos-realworld/  # 到注解文件下 cmd\kratos-realworld\wire.go
  wire
  
  ```

  `internal\server` (...)

  `internal\service` (...)

  ```bash
  kratos run
  # http://localhost:8000/realworld/eric
  
  ```

  



### 写接口

- 写接口  [kratos api](https://go-kratos.dev/docs/component/api)

  postman 代码生成













## 数据库接入与配置修改



















## 项目结构与依赖注入



















## biz层开发和中间件















## 自定义中间件



















## CORS和HTTP中间件自定义

















## HTTP错误返回结构

















## data层开发

















## 错误处理

















## 构建和部署























































