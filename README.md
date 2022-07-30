# IMSystem
This is This is an instant messaging system.
# 1. 技术栈

* 编程语言: Go
* 数据库: MongoDB
* 协议: WebSocket

# 1. 项目模块

* 用户模块
  - 密码登录
  - 发送验证码
  - 用户注册
* 邮箱注册
* 用户详情
* 通讯模块
  - 使用HTTP搭建WebSocket服务
  - 使用Gin框架搭建websocket服务
  - 发送、接受消息
  - 保存消息
  - 获取聊天记录
* 一对一通讯
* 多对多通讯


# 3. 初始化项目

* 创建gomodule： `go mod init IMSystem`
* 下载gin框架： `go get github.com/gin-gonic/gin`
* 下载WebSocket包： `go get github.com/gorilla/websocket`
* 在Docker上安装Mongodb
    - Ubuntu上安装Docker： `sudo apt-get intall -y docker.io`
    - 测试docker是否安装成功： `docker version`
    - 安装mongodb容器：`docker pull mongo:latest`
    - 查看本地镜像检查mongodb是否安装成功： `docker images`
    - 运行容器： `docker run -td --name mongo -p 27017:27017 mongo --auth`
        - -p 27017:27017 ：映射容器服务的 27017 端口到宿主机的 27017 端口。外部可以直接通过 宿主机 ip:27017 访问到 mongo 的服务。
        - --auth：需要密码才能访问容器服务。

    - 看到docker容器是否启动成功：`docker ps`
    - 创建用户名和密码： `docker exec -it mongo mongo admin`
        - 1、创建一个名为admin，密码为123456 的用户
        - `db.createUser({ user:'admin',pwd:'123456',roles:[ { role:'userAdminAnyDatabase', db: 'admin'},"readWriteAnyDatabase"]});`
        - 2、测试连接
        - `db.auth('admin', '123456')`
* 安装mongodb-driver: `go get go.mongodb.org/mongo-driver/mongo `
* 安装jwt包(用于申请token): `go get github.com/dgrijalva/jwt-go`
* 安装email包： `go get github.com/jordan-wright/email`
* 安装redis包(不同版本对应包不同)： `go get github.com/go-redis/redis/v9`


# 4. 建库建表

>  创建`im`库

## 4.1 创建`user_basic`用户集合表
``` json
{
    "account":"账号",
    "password":"密码",
    "nickname":"昵称",
    "sex": 1,  
    "email":"邮箱",
    "avatar":"头像",
    "created_at": 1, 
    "updated_at": 1
}
```

## 4.2 创建`message_basic`消息集合表

``` json
{
    "user_identity":"用户的唯一标识",
    "room_identity":"房间的唯一标识",
    "data":"发送的数据",
    "created_at": 1,
    "updated_at": 1
}
```

## 4.3 创建`room_basic`房间集合表

``` json
{
    "number":"房间号",
    "name":"房间名称",
    "info":"房间简介",
    "user_identity":"房间创建者的唯一标识",
    "created_at": 1,
    "updated_at": 1,
}
```

## 4.4 创建`user_room`用户房间关系表

```json
{
    "user_identity":"用户的唯一标识",
    "room_identity":"房间的唯一标识",
    "message_identity":"消息的唯一标识",
    "created_at": 1,
    "updated_at": 1
}
```


