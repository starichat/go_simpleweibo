# [练手项目]基于 go 语言简单实现类似 weibo 的网站
本项目基于https://github.com/Away0x/gin_weibo 实现一个基于go的微博网站

主要技术架构：
gin + gorm

## 功能简图
https://cdn.learnku.com/uploads/images/201812/14/1/jn0GCr52Zm.png!large

## 用例划分
用户：
-- 注册
-- 激活
-- 修改密码
-- 删除用户
-- 关注用户
-- 取消关注
-- 登录、注销
-- 发布微博、删除微博
-- 查看用户动态流
-- 查看关注和粉丝列表页面
-- 查看个人中心
## 创建应用
### 配置项 viper 构建

## 静态页面开发
为了简单，静态页面采用Bootstrap框架来进行开发。
public 下存放项目静态文件
在resouirce 中存放前端源码



## 模型构建
web 整体架构基于MVC模式，M（模型），根据项目整体结构在此构建一个基本的用户模型。实现用户数据的存储，对模型实体的增删改查。添加用户注册和登录功能，并对用户身份进行权限认证，让管理员可以对用户进行删除操作。接着我们还会构建一套用户账号激活和密码找回系统，只有成功进行邮箱激活的用户才能在网站上进行登录，激活成功后的用户如果出现密码丢失的情况，可以使用已认证的邮箱进行密码找回。

### 用户表
|字段|描述|
|--|--|
|ID|column:id;primary_key;AUTO_INCREMENT;not null|
|CreatedAt|column:created_at|
|UpdatedAt|column:updated_at|
|DeletedAt|column:deleted_at|
|Name|column:name;type:varchar(255);not null|
|Email|column:email;type:varchar(255);unique;not null|
|Avatar|column:avatar;type:varchar(255);not null|
|Password|column:password;type:varchar(255);not null|
|IsAdmin|column:is_admin;type:tinyint(1)|
|ActivationToken|column:activation_token;type:varchar(255)|
|Activated|column:activated;type:tinyint(1);not null|
|EmailVerifiedAt|column:email_verified_at|
|RememberToken|column:remember_token;type:varchar(100)|


#### 博客内容表
|字段|描述|
|--|--|
|ID|column:id;primary_key;AUTO_INCREMENT;not null|
|CreatedAt|column:created_at|
|UpdatedAt|column:updated_at|
|DeletedAt|column:deleted_at|
|Content|column:context;type:text;not null|
|UserId|column:user_id;not null|

### 关注关系表
|字段|描述|
|--|--|
|ID|column:id;primary_key;AUTO_INCREMENT;not null|
|UserId|column:user_id;not null|
|FollowerID|column:follower_id;not null|

### 密码重置表
|字段|描述|
|--|--|
|Email|column:email;type:varchar(255);not null|
|Token|column:token;type:varchar(255);not null|
|CreatedAt|column:created_at|

数据库采用gorm框架开发

https://jasperxu.github.io/gorm-zh/database.html#m
### 用户模型的增删改查

## 功能设计
### 用户注册、登录
一句mvc模型，注册实现比较简单
路由层get展示注册页面，post提交注册表单，在后台数据库实现注册功能
v->c->m
### 会话管理
### 用户管理
### 邮件发送
### 微博管理
### 关注管理


