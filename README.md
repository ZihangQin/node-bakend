## 登录（Login)
- 验证用户是否正确，并且返回token
````
{//request
    username|emil string,
    password string
}

{//response
    Code num,
    message string,
    Data interface{
        bool,
        token string
     } 
}
````

## 注册（register）
- 验证注册信息是否合法
- 将注册信息持久化到数据库

````
{//request
    username string,
    emil string,
    password string,
    state 0|1 //0代表普通用户，1代表特殊用户。。。      
}

{//response
    Code num,
        message string,
        Data interface{
            bool
         } 
}
````

````
前后端分离的架构模式，在 React 中使用 Web3.js 进行智能合约的连接调用，
同时基于后端 Gin 构建 RESTful API 接口，负责处理与智能合约相关的业务逻辑和安全机制，
这样可以使项目的开发更加简单和规范化，并提高程序的灵活性和可扩展性。
````