# Agenda-test

Agenda测试文档



`Command`模块将使用命令进行手动测试

`Service`和`Model`模块使用`go test`配合`ci`进行全自动的测试



## 测试命令

### 用户

#### 用户注册

命令：`user register`/`u register`

参数：

- `user`：用户名，唯一，只允许使用26个字母和数字以及`-`和`_`的组合
- `password`：密码，密码使用`Hash`加`Salt`的方法保存在数据库
- `email`：可选，用户邮箱
- `tel`：可选，用户电话

功能：用户注册，判断用户名是否唯一，然后将用户信息存储到数据库

测试：注册4个用户test、test1、test2和test3 

```bash
Agenda-CLI user register -utest -p123
Register Success! Hi, test
Agenda-CLI user register -utest1 -p123
Register Success! Hi, test1
Agenda-CLI user register -utest2 -p123
Register Success! Hi, test2
Agenda-CLI user register -utest3 -p123
Register Success! Hi, test3
Agenda-CLI user register -utest1 -p123
Register Failed! The user name already exists. # 注册失败，用户已存在
```

#### 用户登陆

命令：`login`

参数：

- `user`：用户名
- `password`：密码

功能：用户登陆，判断密码是否正确，如果正确则修改当前登陆的用户状态（当前用户登陆状态会在最后一项操作的3小时后自动登出）

测试：

```bash
Agenda-CLI login -utest4 -p123 
Login Failed! Incorrect username or password. #登录失败，用户不存在
Agenda-CLI login -utest -p123
Login Success! Hello, test.
```

#### 用户状态

命令： `status`

参数：无

功能：查看当前已登录的用户

测试：

```
Agenda-CLI stauts
No logged in users.
Agenda-CLI login -utest -p123
Login Success! Hello, test.
Agenda-CLI stauts
Current user: test.
```

#### 用户登出

命令：`logout`/ `exit` / `quit`

参数：无

功能：退出登陆，清理当前登陆状态

```bash
Agenda-CLI login -utest -p123
Login Success! Hello, test.
Agenda-CLI stauts
Current user: test.
Agenda-CLI logout 
sign out.
Agenda-CLI stauts
No logged in users.
```

#### 用户查询

命令：`user list`/ `u list`

参数：无

功能：列出当前已注册的所有用户的用户名、邮箱及电话信息

```bash
Agenda-CLI stauts
No logged in users.
Agenda-CLI user list
You must login first to obtain permission! # 未登录无法获取
Agenda-CLI login -utest -p123
Login Success! Hello, test.
Agenda-CLI user list
Username: test    Email:     Tel: 
Username: test1    Email:     Tel: 
Username: test2    Email:     Tel: 
Username: test3    Email:     Tel: 
```

#### 用户删除

命令：`user delete`/ `u delete`

参数：无

功能：删除当前账户，清理登陆状态，移除相关的会议参与信息，并且删除无效会议

```bash
Agenda-CLI stauts
Current user: test.
Agenda-CLI user delete
Delete User Success! 
Agenda-CLI login -utest1 -p123
Login Success! Hello, test1.
Agenda-CLI user list
Username: test1    Email:     Tel: 
Username: test2    Email:     Tel: 
Username: test3    Email:     Tel: 
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
Agenda-CLI login -utest1 -p123
Login Success! Hello, test1.
Agenda-CLI meeting create -tm1 -ptest2 -s1111/11/11-11:11 -e1111/11/11-12:11
Create Meeting [m1] Success!
```

#### 增加会议参与者

命令：`meeting add `/ `m add`

参数：

- `participator`：新增的参与者
- `title`：会议标题

功能：增加会议参与者，检测合法性和可行性

```bash
Agenda-CLI meeting add -tm1 -ptest3
Add Participator Success!
Agenda-CLI meeting add -tm1 -ptest5
Add Participator Failed! Some Participators are not user . # 不存在该用户
```

#### 移除会议参与者

命令：`meeting remove `/ `m remove`

参数：

- `participator`：需要移除的参与者
- `title`：会议标题

功能：移除会议参与者，检测移除后的会议合法性

```bash
Agenda-CLI meeting remove -tm1 -ptest3
Remove Participator Success!
```

#### 查询会议

命令： `meeting query`/`m query`/`meeting search` /`m search`

参数：

- `start`：开始的时间，默认为当前时间
- `end`：结束的时间，默认为10年后

功能：查询指定时间段与自己有关的（作为主持者或者参与者）的会议

```bash
Agenda-CLI meeting query -s=1111/11/11-11:11 -e=1111/11/11-12:11
Query Meeting Success!
Meeting1:
Title: m1    Start: 1111-11-11 19:11:00 +0800 CST    End: 1111-11-11 20:11:00 +0800 CST
Participator: test2
```

#### 取消会议

命令：`meeting delete [title]`/ `m delete [title]`

参数：

- `title`： Path参数，要取消的会议标题

功能：取消指定标题的会议（自己发起的）

```bash
Agenda-CLI meeting delete -tm1
Delete Meeting [m1] Success!
Agenda-CLI meeting query -s=1111/11/11-11:11 -e=1111/11/11-12:11
Query Meeting Success! #已删除，无会议
```

#### 退出会议

命令：`meeting quit [title]`/ `m quit [title]`

参数：

- `title`： Path参数，要退出的会议标题

功能：退出指定标题的会议（自己参与的）

```bash
Agenda-CLI meeting create -tm1 -ptest2 -s1111/11/11-11:11 -e1111/11/11-12:11
Create Meeting [m1] Success!
Agenda-CLI logout
sign out.
Agenda-CLI login -utest2 -p123 
Login Success! Hello, test2.
Agenda-CLI meeting quit -tm1
Quit Meeting [m1] Success!
Agenda-CLI logout
sign out.
login -utest1 -p123
Login Success! Hello, test1.
meeting query -s=1111/11/11-11:11 -e=1111/11/11-12:11
Query Meeting Success! #无会议显示，退出后参与者为0，被删除
```

#### 清空会议

命令：`meeting clear`/ `m clear`

参数：无

功能：清空自己发起的所有会议安排

```bash
Agenda-CLI meeting create -tm1 -ptest2 -s1111/11/11-11:11 -e1111/11/11-12:11
Create Meeting [m1] Success!
Agenda-CLI meeting clear
Clear Meeting Success!
meeting query -s=1111/11/11-11:11 -e=1111/11/11-12:11
Query Meeting Success!
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

测试流程：

- 用户正常注册

  - test function： `TestUserRegister_normal`
  - 所测试的方法：`User().Register()`

- 使用非法用户名注册

  - test function： `TestUserRegister_IllegalName`

  - 所测试的方法：`User().Register()`

- 使用重复用户名注册

  - test function： `TestUserRegister_DuplicateName`
  - 所测试的方法：`User().Register()`

- 用户正常登录

  - test function： `TestUserLogin_normal`
  - 所测试的方法：`User().Login()`

- 不输入用户名登录

  - test function： `TestUserLogin_NullUser`
  - 所测试的方法：`User().Login()`

- 登录密码不正确

  - test function： `TestUserLogin_ErrorPassword`
  - 所测试的方法：`User().Login()`

- 获取当前登录状态

  - test function： `TestStatusGetLoginUser`
  - 所测试的方法：`Status().GetLoginUser()`

- 清空登录状态

  - test function： `TestStatusClearStatus`
  - 所测试的方法：`Status().ClearStatus()`

- 会议正常创建

  - test function： `TestMeetingCreate`
  - 所测试的方法：`Meeting().Create()`

- 使用非法时间创建会议

  - test function： `TestMeetingCreate_TimeOutOfRange`
  - 所测试的方法：`Meeting().Create()`

- 会议开始时间晚于结束时间

  - test function： `TestMeetingCreate_EndTimeEarly`
  - 所测试的方法：`Meeting().Create()`

- 使用重复名称创建会议

  - test function： `TestMeetingCreate_DuplicateTitle`
  - 所测试的方法：`Meeting().Create()`

- 会议参与者不存在于用户列表中

  - test function： `TestMeetingCreate_UserNotExist`
  - 所测试的方法：`Meeting().Create()`

- 会议发起者同时存在于参与者列表中

  - test function： `TestMeetingCreate_BothPresenterAndParticipator`
  - 所测试的方法：`Meeting().Create()`

- 会议时间冲突

  - test function： `TestMeetingCreate_Overlap`
  - 所测试的方法：`Meeting().Create()`

- 正常添加会议参与者

  - test function： `TestMeetingAddParticipator`
  - 所测试的方法：`Meeting().AddParticipator()`

- 将添加的参与者不在用户列表中

  - test function： `TestMeetingAddParticipator_NotExist`
  - 所测试的方法：`Meeting().AddParticipator()`

- 执行添加操作的不是会议发起者

  - test function： `TestMeetingAddParticipator_NotPresenter`
  - 所测试的方法：`Meeting().AddParticipator()`

- 将添加的参与者已存在于当前参与者中

  - test function： `TestMeetingAddParticipator_ParticipatorExist`
  - 所测试的方法：`Meeting().AddParticipator()`

- 将添加的参与者存在时间冲突

  - test function： `TestMeetingAddParticipator_Overlap`
  - 所测试的方法：`Meeting().AddParticipator()`

- 将移除的参与者不在用户列表中

  - test function： `TestMeetingRemoveParticipator_NotExist`
  - 所测试的方法：`Meeting().RemoveParticipator`

- 执行移除操作的不是会议发起者
- 将移除的参与者不在参与者列表中

  - test function： `TestMeetingRemoveParticipator_ParticipatorNotExist`
  - 所测试的方法：`Meeting().RemoveParticipator()`

- 正常移除参与者

  - test function： `TestMeetingRemoveParticipator`
  - 所测试的方法：`Meeting().RemoveParticipator()`

- 查询会议

  - test function： `TestMeetingQuery`
  - 所测试的方法：`Meeting().Query()`

- 将删除的会议不存在

  - test function： `TestMeetingDelete_NotExist`
  - 所测试的方法：`Meeting().Delete()`

- 执行删除会议的不是会议发起者

  - test function： `TestMeetingDelete_NotPresenter`
  - 所测试的方法：`Meeting().Delete()`

- 正常删除会议

  - test function： `TestMeetingDelete`
  - 所测试的方法：`Meeting().Delete()`

- 将退出的会议不存在

  - test function： `TestMeetingQuit_NotExist`
  - 所测试的方法：`Meeting().Quit()`

- 退出会议的操作者不是会议的参与者

  - test function： `TestMeetingQuit_ParticipatorNotExist`
  - 所测试的方法：`Meeting().Quit()`

- 正常退出会议

  - test function： `TestMeetingQuit`
  - 所测试的方法：`Meeting().Quit()`

- 正常清除会议

  - test function： `TestMeetingClear`
  - 所测试的方法：`Meeting().Clear()`



### model_test.go

测试流程：

- 添加用户
  - test function： `TestUserAdd`
  - 所测试的方法：`User().Add()`
- 根据用户名获取用户
  - test function： `TestUserGetByName`
  - 所测试的方法：`User().GetByName()`
- 检测用户是否存在
  - test function： `TestUserExist`
  - 所测试的方法：`User().Exist()`
- 设置当前用户登录状态
  - test function： `TestStatusSetUser`
  - 所测试的方法：`Status().SetUser()`
- 更新登录有效时间
  - test function： `TestRefreshTime`
  - 所测试的方法：`Status().RefreshTime()`
- 清除当前登录状态
  - test function： `TestClearStatus`
  - 所测试的方法：`Status().ClearStatus()`
- 添加会议
  - test function： `TestMeetingAdd`
  - 所测试的方法：`Meeting().Add()`
- 检测会议是否存在
  - test function： `TestMeetingExist`
  - 所测试的方法：`Meeting().Exist()`
- 查询会议
  - test function： `TestMeetingQuery`
  - 所测试的方法：`Meeting().Query()`
- 添加参与者
  - test function： `TestMeetingAddParticipator`
  - 所测试的方法：`Meeting().AddParticipator()`
- 检测某用户是否是会议参与者
  - test function： `TestMeetingIsParticipator`
  - 所测试的方法：`Meeting().IsParticipator()`
- 检测某用户是否是会议发起者
  - test function： `TestMeetingIsPresenter`
  - 所测试的方法：`Meeting().IsPresenter()`
- 根据会议名获取会议
  - test function： `TestMeetingGetMeetingByTitle`
  - 所测试的方法：`Meeting().GetMeetingByTitle()`
- 获取作为发起者的所有会议
  - test function： `TestMeetingGetMeetingAsPresenter`
  - 所测试的方法：`Meeting().GetMeetingAsPresenter()`
- 获取作为参与者的所有会议
  - test function： `TestMeetingGetMeetingAsParticipator`
  - 所测试的方法：`Meeting().GetMeetingAsParticipator()`
- 获取作为发起或参与者的所有会议
  - test function： `TestMeetingGetMeetingByName`
  - 所测试的方法：`Meeting().GetMeetingByName()`
- 移除参与者
  - test function： `TestMeetingRemoveParticipator`
  - 所测试的方法：`Meeting().RemoveParticipator()`
- 删除会议
  - test function： `TestMeetingDelete`
  - 所测试的方法：`Meeting().Delete()`







