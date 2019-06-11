# chattalk
IM,客服IM
## 功能
* 对接各种平台  
    1.web端样例[web-h5](/templates/chat-page.html)  
    2.[百度小程序](examples/baidu_small_program/baidu-chat.js)  
    3.其它小程序【微信...】(实现逻辑同百度小程序案例)  
    4.App(待添加)
    
#### 运行
-  go run server.go
- http://127.0.0.1:9090/    默认账户  admin 12345
- 聊天url http://127.0.0.1:9090/chat
- 客服地址 http://127.0.0.1:9090/admin  

####  部署
- nginx配置[nginx.conf](wss-nginx.conf)
#### go实现的客服IM系统，功能如下：
- 客服和客户一对多聊天
![动图](/static/images/Untitled.gif)
#### 待加功能(TODO)
- 客服管理(增删) 目前只有一个客服
- ~~过期客户列表清除规则~~
- 各种平台接入demo