# 通用辅助api

| 接口 | 请求方法 | 路由规则 | 参数 | 备注 |
| --- | --- | --- | --- | --- |
| 服务状态检查 | HEAD | / | | |
| 服务信息查看 | GET | / | | 查看服务名称 |

### 内网方法

> 只有内网才可以请求这些接口， pprof 可以通过代理服务进行访问。

| 接口 | 方法 | 路由规则 | 参数 | 备注 |
| --- | --- | --- | --- | --- |
| 设置日志级别 | GET | /log/:level | level=`debug` `info` `warn` `error` `fatal` `panic` | 每次设置生效时间为半小时，半小时后自动设置回配置级别 |
| prometheus采集 | GET | /metrics | | |
| pprof 性能分析 | GET | /pprof/*cmd | cmd=`allocs` `block` `cmdline` `goroutine` `heap` `mutex` `profile` `threadcreate` `trace` ||
| debug | GET | /debug/:cmd | cmd 的范围及参数有相关注册组件确定 | |
| reload | GET | /reload/?events=a,b,c,d | events 为事件名称  | |

