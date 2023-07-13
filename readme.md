# 温柔刷题群-gentleman

一个基于gin+gorm开发的项目，未来会逐渐引入其他技术栈

目前整个项目用到的技术栈和开源库

1. Air自动重载
2. gin框架，gorm框架
3. 配置信息的设计，实现
4. 身份验证接口设计
5. JSON解析和验证器封装
6. 同学们刷题信息的爬取，考勤表，飞升表的实现



# 使用指南

## 下载

```shell
git clone https://github.com/GGboya/gentleman.git
```

## 配置MySQL

1、在gentleman/.env  文件中按如下提示配置数据库连接

```go
DB_HOST=  你的数据库host地址
DB_PORT=  你的数据库端口
DB_DATABASE= 你的数据库名字
DB_USERNAME=  你的数据库用户名
DB_PASSWORD=  你的数据库密码
```



## 编译运行

docker开启mysql，或者自己本地有mysql是启动状态



编译前，终端输入

```shell
go mod tidy
```



终端执行

```shell
go run main.go
```





## 代码组织结构

1. app/models     基类文件用以存放模型通用属性和方法

2. app/models/user   模型文件里存放的是模型定义，以及对象操作的逻辑代码

3. bootstrap 	目录存放程序初始化的代码

4. config           目录存放配置信息，按照单独的逻辑区分单独的配置文件

5. pkg                对第三方库代码封装

6. routes           目录存放我们所有项目的路由文件

7. templates     目录存放前端代码

8. gentleman/.env       存放环境变量

9. main.go           程序入口

10. readme.md    说明文档

    