# zeroswagger

给go-zero的API项目加上自动化的Swagger文档。

## 特性

- 零配置，只需增加一行路由。
- 自动搜索当前项目所有.api文件，并生成swagger的接口文档，无需手动执行命令。
- 支持Swagger UI页面显示。
- 使用官方的goctl-swagger来生成Swagger文档逻辑，够稳定。

## 使用方法

```
// 1. 引入zeroswagger
import "github.com/speedphp/zeroswagger"

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	// 2. 增加一行路由
	server.AddRoute(zeroswagger.New("/doc").Route())

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
```

然后访问 http://localhost:8080/doc/swagger.html 即可。

## 感谢

- [go-zero](https://github.com/zeromicro/go-zero) 是一个集成了各种工程实践的 web 和 rpc 框架。通过弹性设计保障了大并发服务端的稳定性，经受了充分的实战检验。
- [goctl-swagger](https://github.com/zeromicro/goctl-swagger) go-zero出品的Swagger文档生成插件。
- [swagger-ui](https://swagger.io/tools/swagger-ui/)