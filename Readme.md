# todo_demo

> 发现坏味道

1. 第一阶段：基本功能
   - 添加 Todo 项。
```shell
➜  todo git:(main) ✗ ./todo add 123123
2022/03/26 17:56:42 读取到当前index 5
2022/03/26 17:56:42 {"content":"123123","user_id":0,"status":1}
➜  todo git:(main) ✗ ./todo list      
2022/03/26 17:56:45 读取到当前index 6
                6       |       123123  |       未完成  
                3       |       今天学习3       |       未完成  
                4       |       今天学习4       |       未完成  
```
- 完成 Todo 项。
```shell
➜  todo git:(main) ✗ ./todo list             
2022/03/26 17:55:52 读取到当前index 5
                3       |       今天学习3       |       未完成  
                4       |       今天学习4       |       未完成  
                5       |       今天2   |       未完成  
➜  todo git:(main) ✗ ./todo done 5
2022/03/26 17:55:59 读取到当前index 5
2022/03/26 17:55:59 修改后数据为:{"id":5,"data":"{\"content\":\"今天2\",\"user_id\":0,\"status\":0}"}
2022/03/26 17:55:59 大小为:307
```
  - 查看 Todo 列表，缺省情况下，只列出未完成的 Todo 项。
```shell
➜  todo git:(main) ./todo list
2022/03/26 17:54:53 读取到当前index 4
                3       |       今天学习3       |       未完成  
                4       |       今天学习4       |       未完成  
```

   - 使用 all 参数，查看所有的 Todo 项。
```shell
➜  todo git:(main) ✗ ./todo list --all
2022/03/26 17:57:08 读取到当前index 6
                5       |       今天2          |       已完成  
                6       |       123123        |       未完成  
                1       |       今天学习        |       已完成  
                2       |       今天学习2       |       已完成  
                3       |       今天学习3       |       未完成  
                4       |       今天学习4       |       未完成  
```
   - Todo 项存储在本地文件中；
     - `/.todo-list`
   - Todo 项索引逐一递增。
     - `/.todo-index `
2. 第二阶段：支持多用户
    - 用户登录。
      - todo login -u
    - 用户退出。
      - todo logout
    - 只能看到当前用户的 Todo 列表；
      - /.<username>-todo-list
    - 同一个用户的 Todo 项索引逐一递增；
      - /.todo-index > username:index
    - 当前用户信息存储在配置文件中 ~/.todo-config。
      - /.todo-config > username|password|...
3. 第三阶段：支持 Todo 列表导入和导出
    - Todo 列表导出。
      - 格式化导出 todo export > todolist
    - Todo 列表导入。
      - todo import -f todolist
4. 第四阶段：支持数据库持久化
    - 在配置文件中，配置数据库连接信息。
    - 没有数据库的情况下，使用本地文件；
    - 在有数据库的情况下，使用数据库；
    - 在本地文件已经存在的情况，将本地信息导入到数据库中。
      - 初始化时,同步本地数据到数据库，之后操作都是数据库操作


## 缺乏业务含义的命名

> 好的命名是可以体现业务含义

- 坏味道
  - 不精准的命名
  - 用技术术语命名。


> 用业务语言写代码

 - `TodoListItem`=> `TodoItems
   - list为技术名称
 - Todo=>NewTodo 
 - Todo=>CreateTodo 
 - Todo=>ChangeTodoStatus 
 - Todo=>DeleteTodo 
 - Todo=>SelectByStatus
 - Todo=>Done
 - ...


> 一个良好的团队实践是，建立团队的词汇表，让团队成员有信息可以参考

## 乱用英语

