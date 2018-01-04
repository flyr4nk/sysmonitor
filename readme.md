# sysmonitor
1. A simple web ui to view current system information, can be turned off
[!webui.png](!webui.png)
2. Configs for us to monitor a specific program and send alarm via an api specified
3. File cleaner for specified directory

```$yaml
basic:
  listen: ":22222"
  enableWeb: true

# 系统报警设置
system:
  cpuLimit: 80
  memLimit: 95
  diskLimit: 90
  systemLoad: 3.0

# 微信API地址，我们用微信报警， 所以这边直接写了微信的地址，以及需要的一个特殊的Header， 默认发Post请求
wechat:
  enable: true
  url: http://10.117.200.107:18081/wechat/send
  header: qh-messenger:wechat

# 文件清理器的配置
cleaner:
  enable: true
  dir:  /opt/jetty/logs
  # 清理两天前没修改过的文件支持的格式为： 1h  2d 这样， 仅支持一种单位要么按天要么按小时
  time: 2d
# 要监控的进程， 请用进程名来区分
processes:
  - name: jetty
    cpuLimit: 80
    memLimit: 98
    connLimit: 1000
    fileLimit: 1
    threadLimit: 1
  - name: monitor
    cpuLimit: 1
    memLimit: 5
    connLimit: 10
    fileLimit: 1
    threadLimit: 1

```

