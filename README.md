# 星巴克点餐系统

## 功能

### 使用场景

用户侧：
1. 用户注册 & 充值，使用钱包里的钱进行支付
2. 用户登录，浏览商品，下单

商家侧：
1. 接收订单
2. 根据订单的商品及数量找到对应配方
3. 根据配方定义 & 订单数量，领取原料，进行制作
4. 订单完成，交付商品

### 功能范围

- 用户管理
  - 注册
  - 登录
  - 会员到期提醒(V2)
- 订单管理
  - 创建订单
  - 订单支付（余额）
  - 充会员，会员续费(V2)
- 钱包管理
  - 展示余额
  - 充值
- 产品管理
  - 商品展示
  - 新增商品、修改、下架
- 原料管理 & 库存管理
  - 原料入库
  - 原料使用
- 配方管理
  - 配方创建
  - 配方运行
  - 配方调整
- 评价系统（V2）

## 技术

后端：
- Golang
- go-zero： 可以省掉很多不必要的工作，只关注业务逻辑即可
- mysql
- Redis
- Docker
- K8s 后期调整

前端：
- Vue3
- TypeScript
- ElementPlus
- Pinia

## 记录

- [X] 使用 goctl 进行项目初始化
- [X] 编写 DDL
- [ ] 编写 API
- [ ] 使用 openapi 根据 API 生成前端代码
- [ ] 前端同时进行开发
- [ ] 后端开始实现接口

### V1 实现的功能

1. 用户注册、登录
2. 钱包充值

## 问题

### 使用 goctl-swagger 生成 api 文档

先下载： `go install github.com/zeromicro/goctl-swagger@latest`

具体可参考： https://github.com/zeromicro/goctl-swagger

### openapi 问题（还未解决，goctl-swagger 不支持 api 文件嵌套）

先采用这个 go install github.com/wumitech-com/goctl-openapi3@latest，后期改掉

### JWT 登录控制

1. 实现一个用于校验用户登录与否的功能
2. 编写中间件来校验每个请求的 token