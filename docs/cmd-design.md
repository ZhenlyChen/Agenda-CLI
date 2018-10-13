# Agenda-design



## 架构设计

本程序分为四层架构。

### Command

命令层负责将命令解析，将数据交给相应的控制函数

### Controller 

控制层负责检验参数的合法性，并且调用相应的服务，显示处理结果

### Service

服务层负责主要的业务逻辑，根据命令的参数处理业务， 返回处理结果

### Model

数据层负责程序数据（用户、会议、状态）的管理，读取，修改并且持久化数据



此外，还有一些通用辅助模块

### Log

负责全局的日志记录，并且持久化日志

### Util

一些通用的辅助函数



## Model

使用`json`保存数据

### User

```json
{
  "user": "string, 用户名",
  "password": "string, 密码",
  "password_salt"："string, 密码加密盐"
  "email": "string, 用户邮箱",
  "tel": "string, 用户电话",
}
```

### Meeting

```json
{
  "title": "string, 会议标题",
  "host_user": "string, 主持用户名",
  "participator": ["string, 会议参与者"],
  "start": "int64, 开始时间戳",
  "end": "int64, 结束时间戳",
}
```

### Status

```json
{
  "user": "string, 当前登陆用户",
  "expires": "int, 有效期时间戳",
}
```

### Logs

```json
{
  "user": "string, 用户名",
  "command": [{
    "time": "int, 时间戳",
    "command": "string, 命令",
    "output": "string, 输出数据"
  }],
}
```



## 参数

```bash
$ agenda command								#无参数
$ agenda command -uhello
$ agenda command -u hello
$ agenda command -u=hello
$ agenda command --user hello
$ agenda command --user=hello
```

所有缩写默认取其全程的第一位字母

以下均以最后一个样例的形式作为样例



## 具体命令

### 用户

#### 用户登陆

命令：`login`

参数：

- `user`：用户名
- `password`：密码

功能：用户登陆，判断密码是否正确，如果正确则修改当前登陆的用户状态（当前用户登陆状态会在最后一项操作的3小时后自动登出）

```bash
$ agenda login --user=Admin --password=123456
```



#### 用户登出

命令：`logout`/ `exit` / `quit`

参数：无

功能：退出登陆，清理当前登陆状态

```bash
$ agenda logout
```



#### 用户注册

命令：`user register`/`u register`

参数：

- `user`：用户名，唯一，只允许使用26个字母和数字以及`-`和`_`的组合
- `password`：密码，密码使用`Hash`加`Salt`的方法保存在数据库
- `email`：可选，用户邮箱
- `tel`：可选，用户电话

功能：用户注册，判断用户名是否唯一，然后将用户信息存储到数据库

```bash
$ agenda user register --user=Admin --password=123456 [--email=a@zhenly.cn] [--tel=18888888888]
```



#### 用户查询

命令：`user list`/ `u list`

参数：无

功能：列出当前已注册的所有用户的用户名、邮箱及电话信息

```bash
$ agenda user list
```



#### 用户删除

命令：`user delete`/ `u delete`

参数：无

功能：删除当前账户，清理登陆状态，移除相关的会议参与信息，并且删除无效会议

```bash
$ agenda user delete
```



### 会议

#### 创建会议

命令：`meeting create`/`m create`

参数：

- `title`：会议主题
- `participator`：会议参与者（多个参与者用`+`分开）
- `start`：起始时间(`yyyy/MM/dd-hh:mm`，如：1998/03/07-11:23)
- `end`：结束时间

功能：创建会议，检查参与者的合法性以及可行性，并且写入数据库

```bash
$ agenda meeting create --title=Hello --participart=zhen+chen+tp --start=2018/10/13-13:33 --end=2018/10/13-14:44
```





#### 增加会议参与者

命令：`meeting add [participator]`/ `m add [participator]`

参数：

- `participator`：Path参数，新增的参与者

功能：增加会议参与者，检测合法性和可行性

```bash
$ agenda meeting add tp+sq
```



#### 移除会议参与者

命令：`meeting remove [participator]`/ `m remove [participator]`

参数：

- `participator`：Path参数，需要移除的参与者

功能：移除会议参与者，检测移除后的会议合法性

```bash
$ agenda meeting remove tp+sq
```



#### 查询会议

命令： `meeting query`/`m query`/`meeting search` /`m search`

参数：

- `start`：开始的时间，默认为当前时间
- `end`：结束的时间，默认为10年后

功能：查询指定时间段与自己有关的（作为主持者或者参与者）的会议

```bash
$ agenda meeting query --start=2018/10/13-13:33 --end=2018/10/13-14:44
```



#### 取消会议

命令：`meeting delete [title]`/ `m delete [title]`

参数：

- `title`： Path参数，要取消的会议标题

功能：取消指定标题的会议（自己发起的）

```bash
$ agenda meeting delete hello
```



#### 退出会议

命令：`meeting quit [title]`/ `m quit [title]`

参数：

- `title`： Path参数，要退出的会议标题

功能：退出指定标题的会议（自己参与的）

```bash
$ agenda meeting quit hello
```



#### 清空会议

命令：`meeting clear`/ `m clear`

参数：无

功能：清空自己发起的所有会议安排

```bash
$ agenda meeting clear
```



### 日志

#### 查看日志

命令： `log show`

参数：

- `start`：开始的时间，默认为当前时间
- `end`：结束的时间，默认为10年后
- `limit`：限制的显示数量

功能：显示自己的操作日志

```bash
$ agenda log show --start=2018/10/13-13:33 --end=2018/10/13-14:44 --limit=100
```



#### 清空日志

命令：`log clear`

参数：无

功能：清空当前用户的操作日志

```bash
$ agenda log clear
```



### 其他

#### 帮助

命令：`help`/ `h`/ `?`

参数：无

功能：查看帮助信息

```bash
$ agenda help
```



#### 查看版本信息

命令： `version` / `v`

参数：无

功能：显示当前版本信息

```bash
$ agenda version
```

