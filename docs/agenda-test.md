# Agenda-test

Agenda测试文档



`Command`模块将使用命令进行手动测试

`Service`和`Model`模块使用`go test`配合`ci`进行全自动的测试



## 测试命令

### 用户

#### 用户状态

命令： `status`

参数：无

功能：查看当前已登录的用户

测试：

```bash
Agenda-CLI status # 显示未登录状态
Agenda-CLI user register -uTest -p123456
Agenda-CLI login -uTest -p123456
Agenda-CLI status # 显示当前登陆用户名
Agenda-CLI logout
Agenda-CLI status # 显示未登录状态
```

#### 用户登陆

命令：`login`

参数：

- `user`：用户名
- `password`：密码

功能：用户登陆，判断密码是否正确，如果正确则修改当前登陆的用户状态（当前用户登陆状态会在最后一项操作的3小时后自动登出）

测试：

```bash
Agenda-CLI login --user=Admin --password=123456 # 未注册用户，登陆失败
Agenda-CLI user register --user=Admin --password=123456
Agenda-CLI login --user=Admin  # 参数不全，登陆失败
Agenda-CLI login --password=123456  # 参数不全，登陆失败
Agenda-CLI login --user=Admin --password=123 # 密码错误， 登陆失败
Agenda-CLI login --user=Admin --password=123456 # 登陆成功
```

#### 用户登出

命令：`logout`/ `exit` / `quit`

参数：无

功能：退出登陆，清理当前登陆状态

```bash
Agenda-CLI user register --user=Mega --password=123456
Agenda-CLI login --user=Mega --password=123456
Agenda-CLI status # 显示当前登陆账户
Agenda-CLI logout # 退出登陆
Agenda-CLI status # 显示未登录状态
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

命令：`meeting add `/ `m add`

参数：

- `participator`：新增的参与者
- `title`：会议标题

功能：增加会议参与者，检测合法性和可行性

```bash
$ agenda meeting add --participator=tp+sq --title=test
```

#### 移除会议参与者

命令：`meeting remove `/ `m remove`

参数：

- `participator`：需要移除的参与者
- `title`：会议标题

功能：移除会议参与者，检测移除后的会议合法性

```bash
$ agenda meeting remove --participator=tp+sq --title=test
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

```v
$ agenda meeting clear
```

### 其他

#### 帮助

命令：`help`/ `--help`/

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





## Go test

`Service`和`Model`模块使用`go test`配合`ci`进行全自动的测试

### service_test.go

- 用户正常注册

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 使用非法用户名注册

  - test function： `TestUserRegister_normal`

  - 所测试的方法：`User().Register()`

- 使用重复用户名注册

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 用户正常登录

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 不输入用户名登录

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 登录密码不正确

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 获取当前登录状态

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 清空登录状态

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 会议正常创建

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 使用非法时间创建会议

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 会议开始时间晚于结束时间

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 使用重复名称创建会议

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 会议参与者不存在于用户列表中

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 会议发起者同时存在于参与者列表中

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 会议时间冲突

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 正常添加会议参与者

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 将添加的参与者不在用户列表中

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 执行添加操作的不是会议发起者

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 将添加的参与者已存在于当前参与者中

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 将添加的参与者存在时间冲突

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 将移除的参与者不在用户列表中

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 将移除的参与者不在参与者列表中

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 正常移除参与者

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 查询会议

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 将删除的会议不存在

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 执行删除会议的不是会议发起者

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 正常删除会议

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 将退出的会议不存在

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 退出会议的操作者不是会议的参与者

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 正常退出会议

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 正常清除会议

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`



### model_test.go









