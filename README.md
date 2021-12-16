# vgo

## TCP 

### client

> 支持断线重连 心跳检测

```log
time="2021-12-16 22:30:52" level=error msg="客户端连接失败...       等待重试...dial tcp 127.0.0.1:9966: connectex: No connection could be made becau    se the target machine actively refused it."
time="2021-12-16 22:30:59" level=error msg="客户端连接失败...       等待重试...dial tcp 127.0.0.1:9966: connectex: No connection could be made becau    se the target machine actively refused it."
time="2021-12-16 22:31:04" level=info msg="client 连接上server端...127.0.0.1:9966"
time="2021-12-16 22:31:05" level=info msg="收到服务端消息: pong2021-12-16 22:31:05"
time="2021-12-16 22:31:06" level=info msg="收到服务端消息: pong2021-12-16 22:31:06"
time="2021-12-16 22:31:07" level=info msg="收到服务端消息: pong2021-12-16 22:31:07"
time="2021-12-16 22:31:08" level=error msg="向服务端写数据异常127.0.0.1:9966 write tcp 127.0.0.1:4529->127.0.0.1:9966: wsasend: An exis ting connection was forcibly closed by the remote host."
time="2021-12-16 22:31:08" level=error msg="接收服务端数据异常...read tcp 127.0.0.1:4529->127.0.0.1:9966: wsarecv: An existing connecti on was forcibly closed by the remote host."
time="2021-12-16 22:31:09" level=error msg="向服务端写数据异常127.0.0.1:9966 write tcp 127.0.0.1:4529->127.0.0.1:9966: wsasend: An exis ting connection was forcibly closed by the remote host."
time="2021-12-16 22:31:09" level=error msg="接收服务端数据异常...read tcp 127.0.0.1:4529->127.0.0.1:9966: wsarecv: An existing connecti on was forcibly closed by the remote host."
time="2021-12-16 22:31:10" level=error msg="向服务端写数据异常127.0.0.1:9966 write tcp 127.0.0.1:4529->127.0.0.1:9966: wsasend: An exis ting connection was forcibly closed by the remote host."
time="2021-12-16 22:31:10" level=error msg="接收服务端数据异常...read tcp 127.0.0.1:4529->127.0.0.1:9966: wsarecv: An existing connecti on was forcibly closed by the remote host."
time="2021-12-16 22:31:11" level=error msg="向服务端写数据异常127.0.0.1:9966 write tcp 127.0.0.1:4529->127.0.0.1:9966: wsasend: An exis ting connection was forcibly closed by the remote host."
time="2021-12-16 22:31:11" level=info msg="断开连接...."

time="2021-12-16 22:31:13" level=error msg="客户端连接失败...       等待重试...dial tcp 127.0.0.1:9966: connectex: No connection could be made because the target machine actively refused it."
```

### server

> 支持心跳检测 主动关闭无效客户端连接

```log
time="2021-12-16 22:33:23" level=info msg="server start... listen on port: 9966"
time="2021-12-16 22:33:29" level=info msg="客户端建立连接...127.0.0.1:4599"
time="2021-12-16 22:33:30" level=info msg="收到客户端消息: ping2021-12-16 22:33:30"
time="2021-12-16 22:33:31" level=info msg="收到客户端消息: ping2021-12-16 22:33:31"
time="2021-12-16 22:33:32" level=info msg="收到客户端消息: ping2021-12-16 22:33:32"
time="2021-12-16 22:33:33" level=info msg="收到客户端消息: ping2021-12-16 22:33:33"
time="2021-12-16 22:33:34" level=error msg="向客户端写数据异常127.0.0.1:4599 write tcp 127.0.0.1:9966->127.0.0.1:4599: wsasend: An exis ting connection was forcibly closed by the remote host."
time="2021-12-16 22:33:34" level=error msg="接收客户端数据异常...read tcp 127.0.0.1:9966->127.0.0.1:4599: wsarecv: An existing connecti on was forcibly closed by the remote host."
time="2021-12-16 22:33:35" level=error msg="向客户端写数据异常127.0.0.1:4599 write tcp 127.0.0.1:9966->127.0.0.1:4599: wsasend: An existing connection was forcibly closed by the remote host."
time="2021-12-16 22:33:35" level=error msg="接收客户端数据异常...read tcp 127.0.0.1:9966->127.0.0.1:4599: wsarecv: An existing connecti on was forcibly closed by the remote host."
time="2021-12-16 22:33:36" level=error msg="向客户端写数据异常127.0.0.1:4599 write tcp 127.0.0.1:9966->127.0.0.1:4599: wsasend: An existing connection was forcibly closed by the remote host."
time="2021-12-16 22:33:36" level=error msg="接收客户端数据异常...read tcp 127.0.0.1:9966->127.0.0.1:4599: wsarecv: An existing connecti on was forcibly closed by the remote host."
time="2021-12-16 22:33:37" level=error msg="向客户端写数据异常127.0.0.1:4599 write tcp 127.0.0.1:9966->127.0.0.1:4599: wsasend: An exis ting connection was forcibly closed by the remote host."
time="2021-12-16 22:33:37" level=info msg="客户端断开连接....127.0.0.1:4599"
```