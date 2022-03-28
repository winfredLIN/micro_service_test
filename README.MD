# Content
## api
## BFF
services->
lib->
管理调用不同services,来实现功能。
服务聚合、访问控制（可以放在网关中）
例如：将所有服务中的权限控制集中在 BFF 层，使下层服务更加纯粹和独立。
作为应用服务层的API服务就是BFF，从基础业务服务如客户服务、订单服务、产品服务等微服务中获取数据，进行一定的聚合和裁剪返回个某个具体业务线的前端应用。
复用缓存
## cmd
注册并启动Service和Client
### Service
### Client
## config
配置全局参数
## lib
保存数据库文件
数据库操作函数
## services
api->
实现service


micro_service
├─Makefile
├─README.MD
├─go.work
├─go.work.sum
├─test.txt
├─services
|    ├─userService
|    |      ├─client.go
|    |      ├─go.mod
|    |      └server.go
|    ├─stockService
|    |      ├─client.go
|    |      ├─go.mod
|    |      └server.go
|    ├─paymentService
|    |       ├─client.go
|    |       ├─go.mod
|    |       └server.go
|    ├─commodityService
|    |        ├─client.go
|    |        ├─go.mod
|    |        └server.go
├─lib
|  ├─.DS_Store
|  ├─lib_test
|  |    ├─go.mod
|  |    └main.go
|  ├─lib_files
|  |     ├─Commodity_Infor.db
|  |     ├─Commodity_Stock.db
|  |     ├─Shopping_Cart.db
|  |     ├─User_Account.db
|  |     └User_Payment.db
|  ├─User_Payment
|  |      ├─go.mod
|  |      └main.go
|  ├─User_Account
|  |      ├─go.mod
|  |      ├─go.sum
|  |      └main.go
|  ├─Shopping_Cart
|  |       ├─go.mod
|  |       └main.go
|  ├─Commodity_Stock
|  |        ├─go.mod
|  |        ├─go.sum
|  |        └main.go
|  ├─Commodity_Infor
|  |        ├─go.mod
|  |        └main.go
├─config
|   ├─config.go
|   ├─config.yaml
|   ├─go.mod
|   └go.sum
├─cmd
|  ├─.DS_Store
|  ├─greeting
|  |    ├─.DS_Store
|  |    ├─server
|  |    |   ├─go.mod
|  |    |   ├─go.sum
|  |    |   └main.go
|  |    ├─client
|  |    |   ├─go.mod
|  |    |   └main.go
|  ├─Payment
|  ├─Login
|  ├─Browse
├─api
|  ├─protobuf
|  |    ├─user
|  |    |  ├─go.mod
|  |    |  ├─user.pb.go
|  |    |  ├─user.proto
|  |    |  └user_grpc.pb.go
|  |    ├─stock
|  |    |   ├─go.mod
|  |    |   ├─stock.pb.go
|  |    |   ├─stock.proto
|  |    |   └stock_grpc.pb.go
|  |    ├─shoppingCart
|  |    |      ├─go.mod
|  |    |      ├─shoppingCart.pb.go
|  |    |      ├─shoppingCart.proto
|  |    |      └shoppingCart_grpc.pb.go
|  |    ├─payment
|  |    |    ├─go.mod
|  |    |    ├─payment.pb.go
|  |    |    ├─payment.proto
|  |    |    └payment_grpc.pb.go
|  |    ├─commodity
|  |    |     ├─commodity.pb.go
|  |    |     ├─commodity.proto
|  |    |     ├─commodity_grpc.pb.go
|  |    |     └go.mod
├─LaunchService
|       ├─server
|       ├─client
|       |   ├─go.mod
|       |   └main.go
├─GinService
|     ├─.DS_Store
|     ├─getparam
|     |    ├─go.mod
|     |    ├─go.sum
|     |    └main.go
|     ├─LoginService
|     |      ├─ginLogin.go
|     |      ├─go.mod
|     |      ├─go.sum
|     |      └login.html
├─BFF
|  ├─PaymentBFF
|  ├─LoginBFF
|  ├─BrowseBFF