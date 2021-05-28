# base_api

应用api

- 需要开启`go modules`功能，默认1.11后默认开启
- HTTP相关基于[gin](https://github.com/gin-gonic/gin)
- Console相关基于[cobra](github.com/spf13/cobra)
- 数据库访问基于[gorm](https://gorm.io/gorm)
- Redis访问基于[redis](https://github.com/garyburd/redigo/redis)
- 开发模式下通过[air](https://github.com/cosmtrek/air)热重启，直接执行`air` 或者`go build ./ && ./go_base_api api -e testing`
- 配置管理基于[viper](https://github.com/spf13/viper)，已对接Apollo， go_apollo测试环境
- 更新依赖后记得执行`go mod vendor`

# 文件结构
- 在docs目录下
