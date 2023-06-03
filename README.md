 `源码获取路径`：https://github.com/ZihangQin/node-bakend.git

## 登录（/account/login)
- 验证用户是否正确，并且返回token
````
{//GET request
    username|emil string,
    password string
}

{//response
    Code num,
    Msg string,
    Data interface{
        bool,
        token string
     } 
}
````

## 注册（/account/register）
- 验证注册信息是否合法
- 将注册信息持久化到数据库
````
{//POST request
    username string,
    emil string,
    password string,
    phone string, //选填
}

{//response
    Code num,
    Msg string,
    Data interface{
            bool
     } 
}
````

## 主页获取用户信息（/browse/user）
- 用于获取登录账号名，登录账号的积分等
````
{//GET request
    authorization string(is token)
}

{//response
    Code num,
    Msg string,
    Data struct{
        username string,
        calculus int
    }
}
````

##  试题管理页面获取试题列表（/browse/testList）
- 用于从数据库中进行分页获取试题列表，以结构体数组传递给前端
````
{//GET request
    page string, //页数
    token string //用户信息
}

{//response
    Code num,
    Msg string,
    Data {
            TestLists [
                            {
                                ID int,
                                UpdateAt time,
                                Title string,
                                Class string,
                                Score int,
                                TitleType string,
                                Difficulty string,
                                QuestionsSetter string,
                                Answer string
                            }
                            ...
                            ...
                            ...
                         ]
            TitlePages int
         }
}
````

## 新增试题（/browse/saveTest）
- 用于将新添加的试题，持久化保存在数据库中
````
{//POST request
    content: string,    //试题题目
    grade: string,     //测试班级
    score: string,    //得分
    type: string,     //试题类型
    difficulty: string,   //难度系数
    answer: string,     //出题人
    token:string    //token
}

{//response
    Code num,
    Msg "success"|error.Error(),
    Data  nil
}
````

## 删除试题（/browse/deleteTests）
- 对于选中的试题在数据库中进行软删除
````
{//POST request
    strList map[string]string //选中的试题id
}

{//response
    Code num,
    Msg string "success"|error.Error(),
    Data nil
}
````

## 搜索试题（/browse/searchTests）
- 对用户输入的值进行多列模糊搜素
````
{//GET request
    data string //用户输入的搜索内容
}

{//response
    Code num,
    Msg string "success"|error.Error(),
    Data {
            test: [
                    {
                        id uint,
                        updateAt string,
                        title string,
                        class string,
                        score int,
                        title_type string,
                        difficulty string,
                        questionsSetter string,
                        answer string
                     }
                    ...
                    ...
                    ... 
                ]
            totle int //搜索后的试题页数，每页十二条数据
        }

}
````

## 修改试题（/browse/updateTests）
- 对选中的试题进行内容的修改
````
{//POST request
    id string,
    type string,
    grade string,
    content string,
    difficulty string,
    score string,
    answer string,
    token string
}

{//response
    Code num,
    Msg string "success"|error.Error()
    Data nil
}
````


````
前后端分离的架构模式，在 React 中使用 Web3.js 进行智能合约的连接调用，
同时基于后端 Gin 构建 RESTful API 接口，负责处理与智能合约相关的业务逻辑和安全机制，
这样可以使项目的开发更加简单和规范化，并提高程序的灵活性和可扩展性。
````