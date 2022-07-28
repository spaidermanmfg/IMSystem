# IMSystem
This is This is an instant messaging system.
# 1. 技术栈

* 编程语言: Go
* 数据库: MongoDB
* 协议: WebSocket

# 1. 项目模块

* 用户模块
* 密码登录
* 邮箱注册
* 用户详情
* 通讯模块
* 一对一通讯
* 多对多通讯
* 消息列表
* 聊天记录列表
# 3. 初始化项目

* 创建gomodule： go mod init IMSystem
* 下载gin框架： go get github.com/gin-gonic/gin
* 下载WebSocket包： go get github.com/gorilla/websocket
* 在Docker上安装Mongodb
    - Ubuntu上安装Docker： sudo apt-get intall -y docker.io
    - 测试docker是否安装成功： docker version
    - 安装mongodb容器：docker pull mongo:latest
    - 查看本地镜像检查mongodb是否安装成功： docker images
    - 运行容器： docker run -td --name mongo -p 27017:27017 mongo --auth
        - -p 27017:27017 ：映射容器服务的 27017 端口到宿主机的 27017 端口。外部可以直接通过 宿主机 ip:27017 访问到 mongo 的服务。
        - --auth：需要密码才能访问容器服务。

    - 看到docker容器是否启动成功：docker ps
    - 创建用户名和密码： docker exec -it mongo mongo admin
        - 1、创建一个名为admin，密码为123456 的用户
        - db.createUser({ user:'admin',pwd:'123456',roles:[ { role:'userAdminAnyDatabase', db: 'admin'},"readWriteAnyDatabase"]});
        - 2、测试连接
        - db.auth('admin', '123456')



