# zeroswagger

给Go-Zero的API项目加上Swagger文档。

## 使用方法

```
import "github.com/speedphp/zeroswagger"

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	server.AddRoute(zeroswagger.New("/doc").Route())

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
```