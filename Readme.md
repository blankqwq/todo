## todo


1. 第一阶段：基本功能
   - 添加 Todo 项。
     - todo add 
   - 完成 Todo 项。
     - todo done <id>
   - 查看 Todo 列表，缺省情况下，只列出未完成的 Todo 项。
     - todo list
   - 使用 all 参数，查看所有的 Todo 项。
     - todo list --all
   - Todo 项存储在本地文件中；
     - /.todo-list
   - Todo 项索引逐一递增。
     - /.todo-index > 行号
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
