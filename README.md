# IoT-admin-backend

~~是某次作业的代码，不知道为什么居然有人 star 于是补充一下 readme，大概大家都在为作业苦恼叭（~~

## Run
go run main.go

## Build
go build main.go

## 项目结构
```powershell
├── db
├── middleware
├── models
└── pkg
    ├── api
    └── handler
```

## db
db 中主要存放数据库的连接逻辑
```go
var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the database.
	MongoDBUrl = "mongodb://localhost:27017/IoT-admin"
)

// Connect connects to mongodb
func Connect() {
	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}
```

## middleware
middleware 中主要存放中间件。

### cors
处理跨域问题
```go
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
```

### dbConnector
数据库连接中间件：克隆每一个数据库会话，并且确保 `db` 属性在每一个 handler 里均有效
```go
func Connect(context *gin.Context) {
	s := db.Session.Clone()
	defer s.Clone()

	context.Set("db", s.DB(db.Mongo.Database))
	context.Next()

}
```

### jwt
 JWTAuth 中间件，检查token
```go
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}
```

## models
主要存放数据结构体
其中注意一点，在定义 `ID` 时，即会在 MongoDB 中自动生成的 `_id` ，必须加上 omitempty ，忽略该字段，否则在创建时此字段为空会报错
```go
	ID               bson.ObjectId   `json:"_id,omitempty" bson:"_id,omitempty"`
```

## api
主要存放路由
统一 api prefix `/api/v1alpha1/`
在部分路由前加上中间件 `	v1.Use(middleware.JWTAuth())`
路由遵循 RESTful 规范

## handler
 主要存放业务逻辑

如：

GET
```go
// Get a product
func GetProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var product models.Product

	err := db.C(models.CollectionProduct).
		FindId(bson.ObjectIdHex(c.Param("_id"))).
		One(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   product,
	})
}
```

CREATE
首先从 token 中解析出用户的 id, 从而加到 product 的 CreatedBy 字段中
并且每新增一个 product 都往 customer 和 organization 中的 productCount 字段加一，且把 productId 加到这两张表的 productId 数组中
```go
// Create a product
func CreateProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	product.CreatedBy = claims.ID
	product.ID = bson.NewObjectId()

	err = db.C(models.CollectionProduct).Insert(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	err = db.C(models.CollectionUser).Update(bson.M{"_id": product.CreatedBy},
		bson.M{"$inc": bson.M{"productCount": 1}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	for _, id := range product.CustomerID {
		err = db.C(models.CollectionCustomer).Update(bson.M{"_id": id},
			bson.M{"$inc": bson.M{"productCount": 1}})
		err = db.C(models.CollectionCustomer).Update(bson.M{"_id": id},
			bson.M{"$push": bson.M{"productId": product.ID}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"msg":    err.Error(),
			})
			return
		}
	}

	err = db.C(models.CollectionOrg).Update(bson.M{"_id": product.OrganizationID},
		bson.M{"$inc": bson.M{"productCount": 1}})
	err = db.C(models.CollectionOrg).Update(bson.M{"_id": product.OrganizationID},
		bson.M{"$push": bson.M{"productId": product.ID}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
	})
}
```

PUT
```go
func UpdateProduct(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	// 查找原来的文档
	query := bson.M{
		"_id": bson.ObjectIdHex(c.Param("_id")),
	}

	// 更新
	err = db.C(models.CollectionProduct).Update(query, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   product,
	})
}
```

## 部署
使用 docker 打包整个后端
Dockerfile:（注意：需要设置时区）
```dockerfile
#源镜像
FROM golang:latest
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR $GOPATH/src/IoT-admin-backend
COPY . $GOPATH/src/IoT-admin-backend
RUN go build .
#暴露端口
EXPOSE 9002
#最终运行docker的命令
ENTRYPOINT  ["./IoT-admin-backend"]
```

除了 IoT-admin 以外，还需要 mongo , 直接使用 dockerhub 上的最新 mongo 镜像跑一个 mongo container 之后，使用 docker-compose 跑两个容器
docker-compose: （version 是 2.0 是因为服务器上的 docker 版本较低）
```
version: '2.0'
services:
  api:
    container_name: 'IoT-admin'
    build: '.'
    ports:
      - '9002:9002'
    volumes:
      - '.:/go/src/IoT-admin'
    links:
      - mongo
    environment:
      MONGODB_URL: mongodb://mongo:27017/IoT-admin
  mongo:
    image: 'mongo:latest'
    container_name: 'mongo'
    ports:
      - '27010:27017'
```

