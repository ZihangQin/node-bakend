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