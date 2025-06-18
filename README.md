# MYGOGIN

## Config 模块
- 细节
    ```bash
  go get -u github.com/gin-gonic/gin@latest
  go get -u gorm.io/gorm@latest
  go get -u gorm.io/driver/mysql@latest
  go get -u github.com/rabbitmq/amqp091-go@latest

  ```
  

## 代码这么分层使用 struct 和 interface 的意义

```bash
┌────────────┐
│ Controller │  -> 负责 HTTP 入口
└─────┬──────┘
      │ 调用 service
┌─────▼──────┐
│  Service   │  -> 业务逻辑
└─────┬──────┘
      │ 调用 repository (interface)
┌─────▼──────┐
│ Repository │  -> 数据持久层 (操作数据库)
│ (UserRepo) │
└─────┬──────┘
      │ 持有 gorm.DB
┌─────▼──────┐
│  GORM/DB   │  -> 具体数据库 ORM
└────────────┘

```