# vgo

## Next Plan

- [ ] MQTT client
- [ ] Gin

## TCP

`vnet/client` `vnet/server`

### client

> 支持断线重连 心跳检测

```log
ERRO[2021-12-17 20:51:18] TCP连接失败, 等待5s重试...dial tcp 127.0.0.1:9966: connectex: No connection could be made because the target machine actively refused it. 
ERRO[2021-12-17 20:51:25] TCP连接失败, 等待5s重试...dial tcp 127.0.0.1:9966: connectex: No connection could be made because the target machine actively refused it. 
ERRO[2021-12-17 20:51:32] TCP连接失败, 等待5s重试...dial tcp 127.0.0.1:9966: connectex: No connection could be made because the target machine actively refused it. 
INFO[2021-12-17 20:51:37] client连接上server端 : 127.0.0.1:9966            
INFO[2021-12-17 20:51:38] 收到服务端消息: pong2021-12-17 20:51:38             
INFO[2021-12-17 20:51:39] 收到服务端消息: pong2021-12-17 20:51:39             
INFO[2021-12-17 20:51:40] 收到服务端消息: pong2021-12-17 20:51:40             
ERRO[2021-12-17 20:51:41] 向服务端写数据异常...write tcp 127.0.0.1:12218->127.0.0.1:9966: wsasend: An existing connection was forcibly closed by the remote host. 
ERRO[2021-12-17 20:51:41] 接收服务端数据异常...read tcp 127.0.0.1:12218->127.0.0.1:9966: wsarecv: An existing connection was forcibly closed by the remote host. 
ERRO[2021-12-17 20:51:42] 向服务端写数据异常...write tcp 127.0.0.1:12218->127.0.0.1:9966: wsasend: An existing connection was forcibly closed by the remote host. 
ERRO[2021-12-17 20:51:42] 接收服务端数据异常...read tcp 127.0.0.1:12218->127.0.0.1:9966: wsarecv: An existing connection was forcibly closed by the remote host. 
ERRO[2021-12-17 20:51:43] 向服务端写数据异常...write tcp 127.0.0.1:12218->127.0.0.1:9966: wsasend: An existing connection was forcibly closed by the remote host. 
ERRO[2021-12-17 20:51:43] 接收服务端数据异常...read tcp 127.0.0.1:12218->127.0.0.1:9966: wsarecv: An existing connection was forcibly closed by the remote host. 
ERRO[2021-12-17 20:51:44] 向服务端写数据异常...write tcp 127.0.0.1:12218->127.0.0.1:9966: wsasend: An existing connection was forcibly closed by the remote host. 
INFO[2021-12-17 20:51:44] 断开连接....                                         
ERRO[2021-12-17 20:51:46] TCP连接失败, 等待5s重试...dial tcp 127.0.0.1:9966: connectex: No connection could be made because the target machine actively refused it. 
ERRO[2021-12-17 20:51:53] TCP连接失败, 等待5s重试...dial tcp 127.0.0.1:9966: connectex: No connection could be made because the target machine actively refused it. 
```

### server

> 支持心跳检测 主动关闭无效客户端连接 连接数统计

```log
INFO[2021-12-17 20:59:45] server start... listen on port: 9966         
INFO[2021-12-17 20:59:50] 客户端: 127.0.0.1:12616 建立连接...                 
INFO[2021-12-17 20:59:51] 收到客户端: 127.0.0.1:12616 消息: ping2021-12-17 20:59:51 
INFO[2021-12-17 20:59:52] 收到客户端: 127.0.0.1:12616 消息: ping2021-12-17 20:59:52 
INFO[2021-12-17 20:59:53] 收到客户端: 127.0.0.1:12616 消息: ping2021-12-17 20:59:53 
INFO[2021-12-17 20:59:54] 收到客户端: 127.0.0.1:12616 消息: ping2021-12-17 20:59:54 
INFO[2021-12-17 20:59:55] 收到客户端: 127.0.0.1:12616 消息: ping2021-12-17 20:59:55 
INFO[2021-12-17 20:59:56] 收到客户端: 127.0.0.1:12616 消息: ping2021-12-17 20:59:56 
ERRO[2021-12-17 20:59:57] 向客户端: 127.0.0.1:12616 写数据异常... write tcp 127.0.0.1:9966->127.0.0.1:12616: wsasend: Anexisting connection was forcibly closed by the remote host. 
ERRO[2021-12-17 20:59:57] 接收客户端: 127.0.0.1:12616 数据异常... read tcp 127.0.0.1:9966->127.0.0.1:12616: wsarecv: Anexisting connection was forcibly closed by the remote host. 
ERRO[2021-12-17 20:59:58] 向客户端: 127.0.0.1:12616 写数据异常... write tcp 127.0.0.1:9966->127.0.0.1:12616: wsasend: Anexisting connection was forcibly closed by the remote host. 
ERRO[2021-12-17 20:59:58] 接收客户端: 127.0.0.1:12616 数据异常... read tcp 127.0.0.1:9966->127.0.0.1:12616: wsarecv: Anexisting connection was forcibly closed by the remote host. 
ERRO[2021-12-17 20:59:59] 向客户端: 127.0.0.1:12616 写数据异常... write tcp 127.0.0.1:9966->127.0.0.1:12616: wsasend: Anexisting connection was forcibly closed by the remote host. 
ERRO[2021-12-17 20:59:59] 接收客户端: 127.0.0.1:12616 数据异常... read tcp 127.0.0.1:9966->127.0.0.1:12616: wsarecv: Anexisting connection was forcibly closed by the remote host. 
ERRO[2021-12-17 21:00:00] 向客户端: 127.0.0.1:12616 写数据异常... write tcp 127.0.0.1:9966->127.0.0.1:12616: wsasend: Anexisting connection was forcibly closed by the remote host. 
INFO[2021-12-17 21:00:00] 客户端: 127.0.0.1:12616 断开连接....   
```