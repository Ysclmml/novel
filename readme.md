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
+ 权限控制
+ 分页展示封装
+ 事务封装



**搭建记录:**

go-playground使用起来并没有那么方便. 自定义错误消息太麻烦了, 尝试使用新的校验器做替代.

https://gitee.com/inhere/validate

https://github.com/gin-gonic/gin/issues/430#issuecomment-141774133



分页展示封装, 写了好多次,考虑封装一下. 

现在书写事务及其麻烦, 需要封装一下. 

异常错误信息需要再统一下, 现在写的有点乱了. 



前端项目搭建:

前端使用vue来搭建项目

项目路径在cmd/front路径下. 

当前已写完首页和详情页两个页面. 

