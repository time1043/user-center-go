# user-center-go

- Refence 

  [user center (golang)](https://articles.zsxq.com/id_tmiv1m92rndu.html), [user center (golang github)](https://github.com/open-user-center), [user center (golang)](https://wx.zsxq.com/dweb2/index/topic_detail/181485581145842), [go web Scaffolding](https://github.com/miomiora/mio-init),

  [go classmate](https://gitee.com/moxi159753/LearningNotes/tree/master/Golang/), 

  [Kratos (bilibili)](https://space.bilibili.com/1885628842), 





## 背景介绍

- 技术选型

  `Golang` (对应`Java`)

  `Wire` (依赖注入框架，帮助你管理 Go 对象，集成一些其他的内容，对应`spring`)

  `Kratos` (B站微服务框架，提供接口访问、restful接口等能力，对应`springmvc`)

  `Gorm` (Go 操作数据库的框架，持久层框架，不用写 sq 也能实现增删改查，对应`mybatis-plus`)

  `Go-redis` (Go 操作redis的框架)

  `mysql` (数据库)

  `redis` (缓存中间件)

  `grpc` (基于h2的七层通信协议，不过我们只用它的http部分)





## 初始化环境

- 项目初始化

  `sql`: 表结构的定义和测试数据

  `api`: protobuf文件 用于定义restfulAPI 和全局错误码

  `app`: 核心业务代码 

  `app/user/service/cmd/main`: 整个工程的入口 (`main.go`  `wire.go`为依赖注入的定义  `wire_gen.go`为生成的注入代码)

  `app/user/service/configs`: 工程的配置文件 (端口配置 mysql redis 各种全局配置)

  `app/user/service/internal/service`: 整个工程的DTO对象 (数据传输对象), 转DO对象 (业务域对象)

  `app/user/service/internal/server`: 整个工程配置 和拉取http server的逻辑 启动成功后才能监听到8080 接受前端的请求 (grpc暂时忽略)

  `app/user/service/internal/pkg`: 业务逻辑公用的工具函数 (两个对象值的互相深拷贝 go反射)

  `app/user/service/internal/data`: 数据入库的逻辑 (DAO层 操作mysql和redis)

  `third_party`: 整个工程使用的第三方库

  `Makefile`: 各类构建脚本 (类似前端的`package.json`)

  ```bash
  cd user-center-go/
  
  kratos new backend-user-center
  kratos new backend-user-center -r https://gitee.com/go-kratos/kratos-layout.git
  cd backend-user-center
  go mod download
  
  
  mkdir backend-user-center && cd backend-user-center/
  go mod init github.com/time1043/user-center-go/backend-user-center
  
  
  
  ```
  
  



## 注册功能 后端

- 注册逻辑

  用户在前端输入**账号和密码**、以及校验码 (todo)

  **校验**用户的账户、密码、检验密码是否符合要求  (非空检验、账户不小于4位、密码不小于8位、账户不能重复、账户不含特殊字符、密码和校验密码相同)

  对密码进行**加密** (千万不要明文存储数据库)

  向**数据库**插入用户数据

- Q：为什么前端校验了，后端还要校验？

  前端只能拦住正常用户，拦不住攻击，用户可以绕过前端向后端接口发请求





### restfulAPI

- restfulAPI

  ```protobuf
    //用户注册
    rpc UserRegister (UserRegisterReq) returns (UserRegisterReply){
      option (google.api.http) = {
        post: "api/user/register",
        body: "*"
      };
    }
  
  
  message UserRegisterReq{
    string userAccount = 1;
    string userPassword = 2;
    string checkPassword = 3;
  }
  
  message UserRegisterReply{
    User data = 1;
  }
  
  
  message User{
    int32 id = 1;
    string userName = 2;
    string userAccount = 3;
    string avatarUrl = 4;
    string phone = 5;
    string email = 6;
    int32 userStatus = 7;
    int32 gender = 8;
    bool empty = 9;
    int32 userRole = 10;
    string createTime = 11;
  }
  
  ```
  
  



### DTO 转 DO

- DTO 转 DO

  ```go
  func (s *UserService) UserRegister(ctx context.Context, req *v1.UserRegisterReq) (*v1.UserRegisterReply, error) {
  	register := &biz.UserRegister{
  		UserAccount:   req.UserAccount,
  		UserPassword:  req.UserPassword,
  		CheckPassword: req.CheckPassword,
  	}
  	err := s.vc.ParamsValidate(register)
  	if err != nil {
  		return nil, err
  	}
  	id, err := s.ac.UserRegister(ctx, register.UserAccount, register.UserPassword, register.CheckPassword)
  	if err != nil {
  		return nil, err
  	}
  	return &v1.UserRegisterReply{
  		Data: &v1.User{
  			Id: id,
  		},
  	}, nil
  }
  
  ```

  



### biz 业务逻辑

- biz业务逻辑

  ```go
  
  ```

  



### data 数据入库

- data 数据入库

  ```go
  
  ```

  



## 登录功能 后端

- 登录逻辑

  检验用户账户和密码是否**合法** (非空、账户不小于4、密码不小于8、账户不含特殊字符)

  校验**密码**是否输入正确，要和数据库密码密文对比

  记录用户的**登录态** (session)，存到服务器上 (后端springBoot封装的服务器 tomcat)

  返回用户信息 (**脱敏**)





### restfulAPI

- restfulAPI

  ```protobuf
    //用户登录
    rpc UserLogin (UserLoginReq) returns (UserLoginReply){
      option (google.api.http) = {
        post: "api/user/login",
        body: "*"
      };
    }
  
  
  message UserLoginReq{
    string userAccount = 1;
    string userPassword = 2;
  }
  
  message UserLoginReply{
    User data = 1;
  }
  
  ```
  
  



### DTO 转 DO

- DTO 转 DO

  ```go
  
  ```
  
  



### biz 业务逻辑

- biz业务逻辑

  ```go
  
  ```

  



### data 数据入库

- data 数据入库

  ```go
  
  ```
  
  





































































































