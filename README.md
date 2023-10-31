# Go Backend Starter

> 專案使用 `ubergo/fx` 依賴注入，並且學習 Clean architecture 開發流程。

### 參考專案
- [Clean-Gin](https://github.com/dipeshdulal/clean-gin/tree/master)

### 使用到的 pkg
```shell
go get go.uber.org/fx@v1
go get -u go.uber.org/zap
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/joho/godotenv
go get -u github.com/spf13/cobra
go get -u github.com/spf13/viper
```

### some articles
- [构建微服务的十大 Go 框架/库](https://server.51cto.com/article/648455.html)
    - [spf13/cobra - 命令列工具](https://github.com/spf13/cobra): CLI 應用程式的庫，也適用於生成應用程序和命令文件的程式
    - [spf13/viper - 配置讀取器](https://github.com/spf13/viper): 存取配置檔，支援: `JSON`, `TOML`, `YAML`, `HCL`, `INI`, `ENV`, `Java properties config files`

# Initial project
```shell
go mod download
```

# run project
```shell
# 使用下列命令 或是 make run 即可
go run . app:serve
```


### JWT Token & Secret
- [simple-jwt-authentication-for-golang-part-1](https://medium.com/@stle/simple-jwt-authentication-for-golang-part-1-6f314b456aa)
