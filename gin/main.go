package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	// "context"		// imported, not need
	// "os"
	// "os/signal"

	"github.com/gin-gonic/autotls" // let's encrypt
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/go-playground/validator/v10"
	"golang.org/x/sync/errgroup"
)

// https://gin-gonic.com/zh-cn/docs/examples/

// https://gin-gonic.com/zh-cn/docs/testing/
// net/http/httptest     to test

// 使用 Gin web 框架的知名项目：
//     gorush：Go 编写的通知推送服务器。
//     fnproject：原生容器，云 serverless 平台。
//     photoprism：由 Go 和 Google TensorFlow 提供支持的个人照片管理工具。
//     krakend：拥有中间件的超高性能 API 网关。
//     picfit：Go 编写的图像尺寸调整服务器。
//     gotify：使用实时 web socket 做消息收发的简单服务器。
//     cds：企业级持续交付和 DevOps 自动化开源平台。
//     go-admin: 前后端分离的中后台管理系统脚手架。
// all in github.


func formatAsDate(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", y, m, d)
}

func main() {

	// r := gin.New()
	// 代替
	// // Default 使用 Logger 和 Recovery 中间件
	// r := gin.Default()
	

	// first_demo()

	// http2_server_push()

	// jsonp()

	// multipart_urlencoded_bind()

	// multipart_urlencoded_form()

	// pure_json()

	// query_post_form()

	// secure_json()

	// xml_json_yaml_protobuf()

	// upload_file()

	// update_multi_file()

	// read_data_from_reader()		// http.Get   visit other url

	// good_restart_stop()		// 优雅的停机。(就是会等待一段时间，以便处理完请求，等待的这段时间内，不接受新的请求)

	// basicAuth()

	// useHttpMethod()

	// useMiddleWare()

	// onlyBindQueryParams()

	// useGoroutineInMiddleware()

	// default, gin use one template;
	// use 1.6 block template support multi template. https://github.com/gin-contrib/multitemplate

	// write_log()

	// custome_log_format()

	// bind_request_body_to_different_struct()

	// support_lets_encrypt()

	// map_query_param_or_form_param()		// duo ge canshu

	// query_params()

	// model_bind_verify()

	// bind_checkbox()

	// bind_uri_param()

	// bind_query_params_form_data()

	// bing_formdata_to_custom_struct()

	// custom_http_config()

	// config_hook()

	// config_log_file()

	// custom_validator()

	// set_get_cookie()

	// route_rule()
	// route_rule_group()

	// multi_service_diff_port()

	// redirect()

	// static_file()

	static_resource()

}

func first_demo() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ascii_json", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag": "<br>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})

	// setFuncMap must before load template. otherwise load template will fail, for go don't know function formatAsDate's define
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.GET("/now_format", func(c *gin.Context) {
		c.HTML(http.StatusOK, "date.tmpl", map[string]interface{}{
			// "now": time.Date(2022,02,02,0,0,0,0,time.UTC),
			"now": time.Now(),
		})
	})

	
	r.LoadHTMLGlob("templates/*/*")		// 这个读取读的是 templates的下级的下级中的， templates的下级中的没有读取。
	// same template file name, different directory
	r.GET("/temp_another_index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "another/index.tmpl", gin.H{
			"title": "yes another",
		})
	})

	// r.LoadHTMLGlob("templates/index.tmpl")		// 后面的load会覆盖前面的，所以导致前面的，浏览器访问时返回空白，但是日志中没有错误日志，只有200
	r.GET("/temp_index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "main website",
		})
	})

	// 自定义模板渲染器
/*
	import "html/template"

	html := template.Must(template.PaeseFiles("file1", "file2"))
	r.SetHTMLTemplate(html)
*/

	// 自定义分隔符	应该是指 {{ }} 的替代品。
	// r.Delims("{[{", "}]}")
	// r.LoadHTMLGlob("templates/*/*")

	r.Run()		// 8080

}

var html = template.Must(template.New("https").Parse(`
<html>
<head>
	<title>https test</title>
	<script src="/assets/app.js"></script>
</head>
<body>
	<h1 style="color:red;">welcome, ginner!</h1>
</body>
</html>
`))


// 这是什么功能。。

// 访问 /h2s/123 后返回下面的，其他的没有了，而且 fmt.Println 也没有看到。。start 都没有看到。
// Client sent an HTTP request to an HTTPS server.

// 服务器推送(server push)指的是,还没有收到浏览器的请求,服务器就把各种资源推送给浏览器
// 我应该是少 ./assets ，导致 服务器推了个寂寞。 应该是。

func http2_server_push() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {		// 原先是 "/h2s"， 但是也没用。
		fmt.Println("start")
		if pusher := c.Writer.Pusher(); pusher != nil {
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("failed to push: %v", err)
			}
		}
		fmt.Println("middle")
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
		fmt.Println("end")
	})
	
	r.RunTLS(":8081", "./tls/server.pem", "./tls/server.key")
}



func jsonp() {
	var r = gin.Default()

	r.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		// http://localhost:8080/jsonp?callback=123
		// 会返回 123({"foo":"bar"}) 
		// 123 就是 callback的值。
		c.JSONP(http.StatusOK, data)
	})
	r.Run(":8080")
}



type LoginForm struct {
	User string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// curl -v --form user=user --form password=pwd http://localhost:8080/login

func multipart_urlencoded_bind() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "pwd" {
				c.JSON(200, gin.H{"status": "login success"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	r.Run()
}




func multipart_urlencoded_form() {
	r := gin.Default()

// [qwer@192 ~]$ curl localhost:8080/form -X POST
// {"message":"","nick":"anonymous","status":"posted"}[qwer@192 ~]$ 

	r.POST("/form", func(c *gin.Context) {
		msg := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status": "posted",
			"message": msg,
			"nick": nick,
		})
	})
	
	r.Run()
}


func pure_json() {
	var r = gin.Default()

	// <  > will be convert to unicode format
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"html": "<b>hi, chk</b>",
		})
	})

	// got <b>
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>hi</b>",
		})
	})

	r.Run()
}


func query_post_form() {
	r := gin.Default()

// curl "http://localhost:8080/post?id=2143&page=342" -X POST -d 'name=adwfffwe&msg=12asd23'


	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		msg := c.PostForm("msg")

		fmt.Printf("%s %s %s %s", id, page, name, msg)
	})
	r.Run()
}


// 使用 SecureJSON 防止 json 劫持。如果给定的结构是数组值，则默认预置 "while(1)," 到响应体。
func secure_json() {
	r := gin.Default()

	// 你也可以使用自己的 SecureJSON 前缀
	// r.SecureJsonPrefix(")]}',\n")

	// r.SecureJsonPrefix(")]}',\n")
	
	// ###["lena","auuuu","foooo"]
	r.SecureJsonPrefix("###")

	r.GET("/somejson", func(c *gin.Context) {
		names := []string{"lena", "auuuu", "foooo"}

		// 将输出：while(1);["lena","austin","foo"]  。。这个是不设置 SecureJsonPrefix 的情况下。
		c.SecureJSON(http.StatusOK, names)
	})

	r.Run()
}

func xml_json_yaml_protobuf() {
	r := gin.Default()

	r.GET("/somejson", func(c *gin.Context) {
		// gin.H 是 map[string]interface{} 的一种快捷方式
		c.JSON(http.StatusOK, gin.H{"message": "msgggg", "status": http.StatusOK})
	})

	r.GET("morejson", func(c *gin.Context) {
		var msg struct {
			Name string `json:"user"`	// 6
			Message string
			Number int
		}
		msg.Name = "lena"
		msg.Message = "hey"
		msg.Number = 12345

		// 注意 msg.Name 在 JSON 中变成了 "user"		// ...6
		// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}		// yes user:
		c.JSON(http.StatusOK, msg)
	})

// <map>
// 	<message>xml</message>
// 	<status>200</status>
// </map>
	r.GET("/somexml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "xml", "status": http.StatusOK})
	})

// download a file... content : 
// msg: yaml
// status: 200
	r.GET("/someyaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"msg": "yaml", "status": http.StatusOK})
	})



// testdata/protoexample/test.proto
	// package protoexample;
	// enum FOO {X=17;};
	// message Test {
	//    required string label = 1;
	//    optional int32 type = 2[default=77];
	//    repeated int64 reps = 3;
	//    optional group OptionalGroup = 4{
	// 	 	required string RequiredField = 5;
	//    }
	// }	

	// download a file. test加一些乱码
	r.GET("/someprotobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"

		data := &protoexample.Test {
			Label: &label,
			Reps: reps,
		}

		c.ProtoBuf(http.StatusOK, data)
	})

	r.Run()
}

func upload_file() {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20	// 8 mb
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		// fmt.Println(file)

// curl -X POST http://localhost:8080/upload -F "file=@./upload_file" -H "Content-Type: multipart/form-data"
// curl -X POST http://localhost:8080/upload -F "file=@upload_file" -H "Content-Type: multipart/form-data"

		dst := "./" + file.Filename
		c.SaveUploadedFile(file, dst)		// 放到 本文件 同级。

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded", file.Filename))
	})

	r.Run()
}


func update_multi_file() {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20;

	r.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

// curl -X POST http://localhost:8080/upload -F "upload[]=@./upload_file" -H "Connt-Type: multipart/form-data" -F "upload[]=@upload_file2"

		for _, file := range files {
			log.Println(file.Filename)
			dst := "./" + file.Filename
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d file upload", len(files)))
	})
	r.Run()
}


func read_data_from_reader() {
	r := gin.Default()
	r.GET("/reader", func(c *gin.Context) {
		fmt.Println("111")
		// response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")

		response, err := http.Get("https://pics4.baidu.com/feed/0df431adcbef76090906737e1fdc81c07dd99ecb.jpeg@f_auto?token=a248dc7bac472fade6ade07392eb27dd")


		// 换个图可以了，不过这个图不是png的。
		// localhost访问以后，下载到本地。名字就是下面设置的 gopher.png。
		// 由于不是 png格式， 所以 可以软件 无法按照 png格式 解析。 手动把 .png 删除了就可以了。让软件自己决定什么类型。

		fmt.Println("222")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		fmt.Println("333")
		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		fmt.Println("444")
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

		fmt.Println("555")
	})
	r.Run()
}


func good_restart_stop() {
	// r := gin.Default()		// use  fvbock/endless
	// r.GET("/", handler)
	// // [...]
	// endless.ListenAndServe(":3333", r)

	// >= 1.8 http.Server.Shutdown()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "welcome gin server")
	})

	srv := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

// 有 处理中的请求时， 会延迟5秒关闭。 开始关闭后， 关闭完成前，访问，无法连接到服务器。
// 没有 处理中的请求，直接关闭


	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen : %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("server exiting")
}


// private data
var secrets = gin.H{
	"foo": gin.H{"email": "foo@bar.com", "phone": "123123123"},
	"au": gin.H{"email": "auuuu", "phone": "5555555"},
	"lena": gin.H{"email": "lllllll", "phone": "11111111"},
}
func basicAuth() {
	r := gin.Default()

	auth := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
		"au": "1234",
		"lena": "qwer",
		"maun": "zxcv",
	}))

	// 触发 "localhost:8080/admin/secrets
	auth.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "no secret :("})
		}
	})

	r.Run()
}

func useHttpMethod() {
	// r := gin.Default()

	// router.GET("/someGet", getting)
	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

}

func useMiddleWare() {
	r := gin.New();		// create router without any middleware

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// auth := r.Group("/")
	// auth.Use(AuthRequired())
	// {
	// 	auth.POST("/login", loginEndpoint)
	// 	auth.POST("/submit", submitEndpoint)
	// 	auth.POST("/read", readEndpoint)

	// 	test := auth.Group("test")
	// 	test.GET("/analytics", analytisEndpoint)
	// }
	r.Run()
}


type Person struct {
	Name string `form:"nnnnn"`
	Address string `form:"add"`
}
func onlyBindQueryParams() {
	r := gin.Default()

	// http://localhost:8080/test?nnnnn=qwe%20123&add=qwe2%202r
	// %20 == whitespace	


	r.Any("/test", startPage)		// any.. so accept: get post put patch head options, delete, connect, trace.
	r.Run()
}
func startPage(c *gin.Context) {
	var p Person
	if c.ShouldBindQuery(&p) == nil {
		log.Println("==only bind query string=")
		log.Println(p.Name)
		log.Println(p.Address)
	}
	c.String(http.StatusOK, "ssssuccessss")
}

func useGoroutineInMiddleware() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		ccp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("done, in path " + ccp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Printf("done .." + c.Request.URL.Path)
	})

	r.GET("/mytest", func(c *gin.Context) {
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("done my.. " + c.Request.URL.Path)		// == async.. why long_async c.Copy() ?
		}()
	})

	r.Run()
}


func write_log() {
	gin.DisableConsoleColor()
	// 开启： gin.ForceConsoleColor()   默认是开启的。

	f, _ := os.Create("gin.log")		// main.go 同级
	// gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.Run()
}

func custome_log_format() {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod string, absolutePath string, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})

	r.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	r.Run()
}


type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}
type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}
func bind_request_body_to_different_struct() {
	r := gin.Default()
	// r.GET("/test", aHandler)
	// r.GET("/test2", anHandler)
	r.POST("/test", aHandler)
	r.POST("/test2", anHandler)

	// test 的 foo 可以。
	//
	// curl 'http://localhost:8080/test' -H "Conent-Type:application/json" -X POST -d '{"bar":"adzxc"}'
	// 上面这个不行， no match 。 
	// 但是 换成 test2，  foo/bar 都可以。
	//
	// text bar 下面有原因， 确实不行。


	r.Run()
}
func aHandler(c *gin.Context) {
	obja := formA{}
	objb := formB{}

	// // 一般通过调用 c.Request.Body 方法绑定数据，但不能多次调用这个方法。
	// 所以上面的 text  bar 是不行的， 

	// // c.ShouldBind 使用了 c.Request.Body，不可重用。
	if e := c.ShouldBind(&obja); e == nil {
		c.String(http.StatusOK, `the body is formA`)
	
	// // 因为现在 c.Request.Body 是 EOF，所以这里会报错。
	} else if e := c.ShouldBind(&objb); e == nil {
		c.String(http.StatusOK, `the body is bbbbb`)
	} else {
		c.String(http.StatusBadRequest, `no match`)
	}
}
// c.ShouldBindBodyWith 会在绑定之前将 body 存储到上下文中。 
// 这会对性能造成轻微影响，如果调用一次就能完成绑定的话，那就不要用这个方法。
// 只有某些格式需要此功能，如 JSON, XML, MsgPack, ProtoBuf。 
// 对于其他格式, 如 Query, Form, FormPost, FormMultipart 可以多次调用 c.ShouldBind() 而不会造成任任何性能损失 
func anHandler(c *gin.Context) {
	a := formA{}
	b := formB{}
	// // 读取 c.Request.Body 并将结果存入上下文。
	if e := c.ShouldBindBodyWith(&a, binding.JSON); e == nil {
		c.String(http.StatusOK, `aaaa`)

	// // 这时, 复用存储在上下文中的 body。
	} else if e := c.ShouldBindBodyWith(&b, binding.JSON); e == nil {
		c.String(http.StatusOK, `bbbbbbbbbb`)
	} else {
		c.String(http.StatusBadRequest, `bbbaaaaddddd`)
	}
	// // 可以接受其他格式
	// c.ShouldBindBodyWith(&b, binding.XML)
}



func support_lets_encrypt() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
// 。。“ACME”是Automated Certificate Management Environment (自动化证书管理环境)的缩写。这是一个RFC 8555国际标准
// the ACME protocol requires either port 80 or 443.
// 
// listen tcp :443: bind: permission denied
// 好像要 root 用户。
// 放弃了。  root 没有配环境变量， 所以 /usr/local/go/bin/go run main.go
//		。。 要下 依赖的。。放弃了。。
// 不是必须root的， 还有其他方法，但是 好繁的。。
	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))

	// r := gin.Default()
	// // Ping handler
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.String(200, "pong")
	// })
	// m := autocert.Manager{
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
	// 	Cache:      autocert.DirCache("/var/www/.cache"),
	// }
	// log.Fatal(autotls.RunWithManager(r, &m))	

}

func map_query_param_or_form_param() {
	r := gin.Default()

	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

// curl -X POST 'http://localhost:8080/post?ids[a]=123&ids[b]=hi' -H "Content-type:application/x-www-form-urlencoded" -d 'names[first]=asd&names[second]=wer'
// ids %+v, names: %#v map[a:123 b:hi] map[first:asd second:wer]

		fmt.Println("ids %+v, names: %#v", ids, names)
	})
	r.Run()
} 

func query_params() {
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		fname := c.DefaultQuery("fname", "none")
		lname := c.Query("lname")		// // c.Request.URL.Query().Get("lastname") 的一种快捷方式
// http://localhost:8080/welcome?fname=ouha&lname=chk
		c.String(http.StatusOK, "hi %s %s", fname, lname)
	})
	r.Run()
}


type Login struct {
	User string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
func model_bind_verify() {
	r := gin.Default()

// https://gin-gonic.com/zh-cn/docs/examples/binding-and-validation/


// curl -X POST 'http://localhost:8080/loginjson' -H "Content-type:application/json" -d '{"user":"admin","password":"123"}'
	r.POST("/loginjson", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "admin" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged"})
	})


// curl -X POST 'http://localhost:8080/loginxml' -H "Content-type:application/xml" -d 'xmldata=<root><user>admin</user><password>123</password></root>'
// curl -X POST 'http://localhost:8080/loginxml' -H "Content-type:application/xml" -d '<root><user>admin</user><password>123</password></root>'
	

	r.POST("loginxml", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if xml.User != "admin" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauth"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "logged in"})
	})



// curl -v --form user=admin --form password=123 'http://localhost:8080/loginform'	

	r.POST("/loginform", func(c *gin.Context) {
		var form Login
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}
		if form.User != "admin" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthhhh"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "u login"})
	})

	r.Run()
}


type myForm struct {
	Colors []string `form:"colors[]"`
}
// func bind_checkbox(c *gin.Context) {

// // https://gin-gonic.com/zh-cn/docs/examples/bind-html-checkbox/	

// 	var fakeForm myForm
// 	c.ShouldBind(&fakeForm)
// 	c.JSON(200, gin.H{"colors": fakeForm.Colors})

// }


// func bind_checkbox() {
// 	r := gin.Default()
// 	r.POST("/form231", bind_cb)

// 	r.Run()
// }
// func bind_cb(c *gin.Context) {
// 	var fakeForm myForm
// 	c.ShouldBind(&fakeForm)
// 	c.JSON(200, gin.H{"color": fakeForm.Colors})
// }

type Person2 struct {
	ID string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}
func bind_uri_param() {
	var r = gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		var person Person2
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	r.Run();
}

type Person3 struct {
	Name string `form:"name"`
	Address string `form:"address"`
	BirthDay time.Time `form:"birthday" time_format:"2022-10-01" time_utc:"1"`
}
func bind_query_params_form_data() {
	r := gin.Default()
	r.GET("/testing3", startPage3)
	r.Run(":8080")
}
func startPage3(c *gin.Context) {
	var person Person3
	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
	// 查看更多：https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
	if (c.ShouldBind(&person) == nil) {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.BirthDay)
	}
	c.String(200, "ooooookkk")
}

type SA struct {
	FA string `form:"field_a"`
}
type SB struct {
	NestedStruct SA
	FB string `form:"field_b"`
}
type SC struct {
	NestedStructPointer *SA
	FC string `form:"field_c"`
}
type SD struct {
	NestedAnonyStruct struct {
		FX string `form:"field_x"`
	}
	FD string `form:"field_d"`
}
func GetDataB(c *gin.Context) {
	var b SB
	c.Bind(&b)
	c.JSON(200, gin.H{"a": b.NestedStruct, "b": b.FB})
}
func GetDataC(c *gin.Context) {
	var sc SC
	c.Bind(&sc)
	c.JSON(200, gin.H{"a": sc.NestedStructPointer, "c": sc.FC})
}
func GetDataD(c *gin.Context) {
	var sd SD
	c.Bind(&sd)
	c.JSON(200, gin.H{
		"x": sd.NestedAnonyStruct,
		"d": sd.FD,
	})
}
func bing_formdata_to_custom_struct() {
	r := gin.Default()
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)
	r.Run()
}
// not supper below struct:
// type StructX struct {
//     X struct {} `form:"name_x"` // 有 form
// }

// type StructY struct {
//     Y StructX `form:"name_y"` // 有 form
// }

// type StructZ struct {
//     Z *StructZ `form:"name_z"` // 有 form
// }

func custom_http_config() {
	// r := gin.Default()
	// http.ListenAndServe(":8080", r)

	r2 := gin.Default()
	s := &http.Server{
		Addr: ":8081",
		Handler: r2,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1<<15,
	}
	s.ListenAndServe()
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("vara", "asd123")
		
		// before request
		c.Next()
		// after request

		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}
func config_hook() {
	r := gin.New()
	r.Use(Logger())
	r.GET("/testhook", func(c *gin.Context) {
		t2 := c.MustGet("vara").(string)
		log.Println(t2)
	})
	r.Run()
}


func config_log_file() {
	router := gin.New()
	// LoggerWithFormatter 中间件会写入日志到 gin.DefaultWriter
	// 默认 gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Run(":8080")
}



type Booking struct {
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}
var bookabledate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}
func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "booking date is valid"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
func custom_validator() {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookabledate)
	}
	r.GET("/bookable", getBookable)
	r.Run()
}



func set_get_cookie() {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "notset"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("cookie value: %s\n", cookie)
	})
	r.Run()
}


func route_rule() {
	r := gin.Default()

	// 此 handler 将匹配 /user/john 但不会匹配 /user/ 或者 /user
	r.GET("/usr/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hi %s", name)
	})

	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
	r.GET("/usr/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")			// .... action[0] == '/'
		msg := name + " is " + action
		c.String(http.StatusOK, msg)
	})

	r.Run()
}

func route_rule_group() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/asd", my_log("v1/asd"))		// 闭包有那么一点点可取之处。
		v1.GET("/qwe", my_log("v1/qwe"))
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/123", my_log("v2/123"))
		v2.GET("/234", my_log("v2/234"))
	}

	r.Run()
}

func my_log(id string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("this path is " + id)
	}
}

// ----------------


var (
	g errgroup.Group
)
func r1() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H {
				"code": 200,
				"error": "welcom to server 01",
			},
		)
	})
	return r
}
func r2() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H {
				"code": 200,
				"error": "server 02",
			},
		)
	})
	return r;
}
func multi_service_diff_port() {
	server1 := &http.Server {
		Addr: ":8080",
		Handler: r1(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server2 := &http.Server {
		Addr: ":8081",
		Handler: r2(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server1.ListenAndServe()
	})

	g.Go(func() error {
		return server2.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func redirect() {
	r := gin.Default()

	// 302 + 200   	client receive:200
	r.GET("/redirect1", func(c *gin.Context) {

		// redirect get
		// c.Redirect(http.StatusMovedPermanently, "www.google.com")

		// redirect post ... but in my code, it still get..
		c.Redirect(http.StatusFound, "/foo")
	})

	r.GET("/foo", func(c *gin.Context) {
		c.String(http.StatusOK, "this is fooooooo")
	})

	// ---------
	// 200 + 200	client receive:200
	r.GET("/redirect2", func(c *gin.Context) {
		c.Request.URL.Path = "/foo"
		r.HandleContext(c)
	})

	r.Run()
}


func static_file() {
	// router := gin.Default()
	// router.Static("/assets", "./assets")
	// router.StaticFS("/more_static", http.Dir("my_file_system"))
	// router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// // 监听并在 0.0.0.0:8080 上启动服务
	// router.Run(":8080")
}


func static_resource() {
	// r := gin.New()

	// t, err := loadTemplate()
	// if err != nil {
	// 	panic(err)
	// }
	// r.SetHTMLTemplate(t)

	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "/html/index.tmpl", nil)
	// })
	// r.Run(":8080")
}

// // loadTemplate 加载由 go-assets-builder 嵌入的模板
// func loadTemplate() (*template.Template, error) {
// 	t := template.New("")
// 	for name, file := range Assets.Files {			// use  go-assets
// 		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
// 			continue
// 		}
// 		h, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			return nil, err
// 		}
// 		t, err = t.New(name).Parse(string(h))
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return t, nil
// }
