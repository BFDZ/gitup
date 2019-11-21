# GitUp
服务器端的Git仓库自动拉取工具。

# Get Start
## 获取GitUp
```shell
# 下载程序
wget https://raw.githubusercontent.com/zhshch2002/gitup/master/build/gitup-linux-amd64
# 设置为可执行文件
chmod +x gitup-linux-amd64
```
其他平台的可执行文件请看[build](https://github.com/zhshch2002/gitup/tree/master/build)目录。

## 设置Git保存密码
```shell
# 请前往工作仓库
git config credential.helper store # 开启保存密码功能
# 然后执行一次git pull，这一次输入密码后将不会再要求输入密码。
```

## 配置
请在与程序同级目录下创建`config.yml`。
```yaml
listen: 0.0.0.0:8000 # WebHook服务器监听配置
repo:
  - dir: /bulabula1  # 仓库本地地址
    branch: origin/master
    mode: ontime # 触发模式 ontime:根据time的配置轮训。适用于不暴露在公网的情景。
    time: "0 0 */1 * *" # 配置方法参考 https://github.com/robfig/cron

  - dir: /bulabula2
    branch: origin/master
    mode: webhook # 触发模式 webhook:程序运行后会创建一个WebHook链接，当有请求发送到那个地址时会触发更新。
```

## 运行
```shell
./gitup-linux-amd64 # 直接运行

nohup ./gitup-linux-amd64 > gitup.log 2>&1 & # 后台运行
```

## 获取WebHook链接
当运行GitUp后会输出目录和链接的对应关系。例如：
```
2019/11/21 21:02:56 config loaded
2019/11/21 21:02:56 Set OnTime job /bulabula1
2019/11/21 21:02:56 /bulabula2 - /a/ab759829d5f73f2586f1717d47ccd670  这里就是链接
2019/11/21 21:02:56 Listen on 0.0.0.0:8000
```

## 注意！
请注意不要外泄GitUp提供的链接。连接的泄露可能造成安全问题。