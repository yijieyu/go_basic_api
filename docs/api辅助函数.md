# api 辅助函数

程序文件路径： api/helper.go

- Bind 解析请求参数到任意结构体
- NewResponse 获得一个通用 Response 对象，保证 api 响应格式统一
- RequestID 获取当前请求的唯一id，可以根据日志追溯请求链
- Logger 获取当前请求唯一的 Logger 实例，该实例默认包含一些重要的属性，便于日志追踪
- ClientIP 获取当前请求用户的客户端 ip
