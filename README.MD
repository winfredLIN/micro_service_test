# 目录结构
# api
## protobuf
### greeting

# cmd
## greeting
### client
### server

# lib
## dbcontext

# services
## greeting

# 梳理
1、api 里面只方法的定义 也就是proto的文件
	1.1任务1：扩展了两个方法并实现
2、cmd里面实现可执行的代码，也就是main.go的文件 main函数 实现服务
	2.1 每一个server或者client的实现要放在一个文件夹下，文件使用main.go命名，与后序wire依赖注入相衔接。（wire根据文件夹进行依赖注入）
	2.2
3、lib（pkg）里面实现 数据库的功能
	3.1 数据库使用GROM
	3.2 任务2：使用GROM设计并实现一个用户表需要用到①增删改查②把数据据库migration
	3.3 提示：设计用户User表需要考虑什么信息是最需要的
4、service 里面实现定义了的方法
	4.1实现任务1中的两个方法
5、其他部分到有使用时再进行增加
6、对于client server要创建一个配置中心，前提是先创建一个配置文件，使用的是Viper
7、go track 服务跟踪
	7.1go track会跟踪所有服务以及其状态，并且呈现在web界面上，可以通过追踪服务来发现服务的各项状态，可有用于发现可优化的服务。比如：一系列服务的响应时间太长，可以发现导致响应时间变长的服务
	7.2 jaeger opentracing
	7.3 案例：order web  -> user-web  -> user rpc  -> order rpc ->
8、健康检查
9、服务注册与服务发现
	9.1 用于发现新的服务以及发现可以结束的服务
10、消息队列
	10.1案例：解决商城抢购的排队问题，排队，支付>>库存减一 取消支付>>返还库存
11、auth鉴权
	11.1例如：User（grpc）调用的是gRPC服务，不鉴权  user-web（gin）gin需要鉴权
12、之后的计划
	实现一个网上商城，完成以下服务
	1.注册/登录 包括：auth鉴权使用jwt
	2.用户信息修改
	3.商品服务
	4.订单服务
	5.用户购买商品
	6.用户查询商品状态
	7.退款
	8. ...
13、未来的规划
	github已经将rest api ->迁移至 graphql
	json -> github最新的api已经完全迁移到graphql了
# 任务
任务1：扩展了两个方法并实现🆗
任务2：使用GROM设计并实现一个用户表需要用到①增删改查②把数据据库migration🆗
任务3：配置vscode 下载字体firacode🆗

///////////////////////////////////////////////////////以下已经整理完了
https://grpc.io/docs/languages/go/quickstart/

数据库用GORM

配置文件  用Viper

开发环境优化

扩展两个

extil 让业务和工具解耦 pull request

做一个User微服务   Gorm  设计一个用户表   增删改查   把数据据库migration

用户 User 表： 考虑什么信息是最需要的

User（grpc）  user-web（gin）gin 最广泛 web框架  user-graphql

rest api -> graphql

json -> github最新的api已经完全迁移到graphql了

postman  postwomen

注册/登录     auth鉴权  jwt

用户信息修改

商品服务

订单服务

用户购买商品

用户查询商品状态

退款

80%