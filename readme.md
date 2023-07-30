# 温柔刷题群-gentleman

## 项目介绍：

本项目是温柔刷题群对应的考勤网站。目前支持登录注册，查看一系列考勤表等功能。

涉及到的功能有：

1、配置信息（使用 Viper，支持 .env 和 config 目录 ）

2、API版本，API错误码，API限流

3、注册登录

- 注册
  -  判断QQ是否注册
  - 支持QQ+图片验证码进行注册
- 登录
  - 支持密码登录（QQ，用户名任选）
  - 支持更加安全的Token Refresh机制

4、JWT授权

5、整个应用使用命令行模式

6、图片验证码，防机器人滥用

- 支持通过配置信息自定义复杂度
- 内置Redis驱动，以接口方式编写，支持多驱动

7、日志记录

- 集成zap高性能日志库
- 支持命令行记录（开发时快速定位问题）
- 命令行日志高亮
- 支持文件记录（多文件和日期分隔）
- 记录gorm的query log
- 记录http的请求log
- Panic Recovery 中间件
- 日志等级（debug, info, error, panic, fatal）



## 项目技术栈：

1. 使用gin框架，它提供了高性能的路由和中间件支持，使得构建API变得简单而灵活

2. 使用gorm简化数据库操作，并保持代码的简洁性和可读性，防止sql注入。

3. mysql存储用户相关的数据，redis存储验证码

4. 使用jwt进行鉴权，减轻缓存服务器得压力。

5. 使用base64Captcha和阿里云短信服务，确保用户注册和登录的安全性

6. 为了保护API免受恶意请求，集成了限流器库limiter，限制了API的访问频率

7. 为了优化项目的部署和配置管理，使用Viper库，帮助我们统一管理配置信息

8. 编写数据库迁移和填充脚本，简化项目的部署和初始化过程。
9. 用zap库实现高效的日志记录，它是一个快速、结构化的日志记录库，能够更加方便地跟踪和调试代码。
10. 使用Cobra框架，为项目提供命令行交互功能，快速构建自定义命令，方便管理和维护项目。

11. 代码行数：

```
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              60            602            615           2393
HTML                            13            316            194           1547
CSS                              3             25              0            341
JSON                             1              0              0            186
Markdown                         1             51              0             69
-------------------------------------------------------------------------------
TOTAL                           78            994            809           4536

```



## 使用指南

### 下载

```shell
git clone https://github.com/GGboya/gentleman.git
```

### 配置信息的填充

1、在gentleman/.env.example  ，把数据库得部分自己补充下

```go
APP_NAME=Gentleman
APP_ENV=local
APP_KEY=zBqYyQrPNaIUsnRhsGtHLivjqiMjBVLS
APP_DEBUG=true
APP_URL=http://localhost:9090
APP_PORT=9090

DB_CONNECTION=mysql
DB_HOST=
DB_PORT=
DB_DATABASE=
DB_USERNAME=
DB_PASSWORD=
DB_DEBUG=2

REDIS_HOST=
REDIS_PORT=
REDIS_PASSWORD=
REDIS_CACHE_DB=0
REDIS_MAIN_DB=1

JWT_EXPIRE_TIME=120
JWT_MAX_REFRESH_TIME=86400

VERIFY_CODE_LENGTH=6
VERIFY_CODE_EXPIRE=15

LOG_TYPE=daily
LOG_LEVEL=debug
```



### 下载依赖

```
go mod tidy
```





### 编译运行

docker开启mysql，或者自己本地有mysql是启动状态



编译前，终端输入

```shell
go mod tidy
```



终端执行

```shell
go run main.go
```





