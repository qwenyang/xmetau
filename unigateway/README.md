grpc-gateway v2 demo简单实现，详细可以看 [教程](grpc-gateway教程.md) 亦或者 [博客链接](https://www.cnblogs.com/cxt618/p/15647316.html)


统一路由规范：
目标：实现一个平台无关的，业务无关统一接口服务
业务和应用相关的接口， 只有登录接口，其他接口均是业务无关的

业务平台无关 统一路由接口规范 /game/{场景}/{接口名}

只有登录接口时业务平台相关的接口
登录接口规范  /game/{平台名称}/{游戏名称}/login
备注:
平台名称 如 微信小游戏 wx  安卓平台android  苹果ios 等等
游戏名称 如斗子象棋 chess 桌球 billiard 等等

本服务支持了多个小游戏，可以根据main包中config.go，修改成自己小游戏的appId 和 appSecret
如果unidao服务和本服务不是单机部署，可以修改server/config.go 中寻址地址
      



