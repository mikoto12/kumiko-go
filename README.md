# kumiko

Gin + GORM 项目模板

kumiko/
├── cmd/                # 程序入口（main函数）
│   └── main.go
├── config/             # 配置文件（如yaml/json/env）
│   └── config.yaml
├── internal/           # 内部应用逻辑（核心业务）
│   ├── consumer/       # mq消费者层
│   │   └── consumer.go
│   ├── controller/     # 控制器层，处理 HTTP 请求
│   │   └── user.go
│   ├── service/        # 服务层，封装业务逻辑
│   │   └── user.go
│   ├── model/          # 数据模型（struct） + ORM 映射
│   │   └── user.go
│   ├── dao/            # 数据访问层（封装 GORM/sqlx）
│   │   └── user.go
│   ├── middleware/     # Gin 中间件（JWT, 日志, 限流等）
│   │   └── jwt.go
│   ├── router/         # 路由注册（Gin Engine 设置）
│   │   └── router.go
│   └── utils/          # 通用工具函数（加密、响应封装等）
│       └── response.go
├── pkg/                # 第三方包封装（如日志库、数据库初始化）
│   ├── logger/         #日志库
│   └── database/       #数据库
│   └── elasticsearch/  #elasticsearch
│   └── rabbitmq/       #rabbitmq
│   └── redis/          #redis
├── docs/               # Swagger 文档、接口文档等
│   └── swagger.yaml
├── go.mod
├── go.sum
└── README.md
