# PPanel-node

A PPanel node server based on xray-core, modified from v2node.  
一个基于xray内核的PPanel节点服务端，修改自v2node

## 软件安装

### 一键安装

```
wget -N https://raw.githubusercontent.com/perfect-panel/PPanel-node/master/scripts/install.sh && bash install.sh
```

## 构建
``` bash
GOEXPERIMENT=jsonv2 go build -v -o ./node -trimpath -ldflags "-s -w -buildid="
```

## Protobuf 面板接口

节点会在拉取配置时通过 `Accept: application/protobuf` 自动协商传输格式。面板返回
Protobuf 时，后续用户列表、在线用户、流量和状态上报都会使用
`application/protobuf`；面板返回 JSON 时，节点继续使用 JSON，无需配置开关。

面板可能提供当前节点版本尚未支持的入站协议。节点会跳过这类协议及其配置，不会因
它们导致其他已支持的协议无法启动。
