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
实现service