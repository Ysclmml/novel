最近在学习go语言, 决定拿一个项目来边学边练, 这样学习起来效率最高,

最终决定参考`github`上的一个`java`小说项目来搭建一个go版本的小说网站. 

书城原网址:

https://github.com/201206030/novel-plus

项目骨架参考了2个方案

1. https://github.com/bullteam/zeus
2. https://gitee.com/daitougege/GinSkeleton

`Todo List`:

+ 搭建项目骨架

  + 搭建目录结构  ok
+ 项目初始化逻辑     ok
  + 配置文件初始化viper配置  ok
  + redis缓存配置  ok
  + 多数据库配置(添加读写分离配置)  ok
+ 路由配置  ok
  + 全局统一响应状态 ok
+ 全局统一校验 (参数错误信息未国际化处理, 是英文的错误信息)
  + 自定义时间格式化信息



**搭建记录:**

go-playground使用起来并没有那么方便. 自定义错误消息太麻烦了, 尝试使用新的校验器做替代.

https://gitee.com/inhere/validate

