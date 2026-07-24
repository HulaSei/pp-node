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

## TLS 证书

`cert_mode: dns` 和 `cert_mode: http` 使用 ACME 自动签发及续期；`self` 生成本地自签名
证书。自动签发与自签均使用 ECDSA P-256；`file` 从
`/etc/PPanel-node/{协议类型}{服务器 ID}.cer` 和 `.key` 读取已有证书对。
自动续期成功后节点会重载以启用新证书。

自动签发时，建议在本地配置的 `Api` 段填写 `ACMEEmail`，以接收 CA 的证书通知。可选的
`ACMECADirURL` 可用于测试 CA（例如 Let's Encrypt staging 或 Pebble）；首次使用某个 CA
前就应设置该值。

当 `cert_mode: self` 时，节点会将证书 DER 的 SHA-256 指纹编码为小写十六进制，并在后续
节点 API 请求中使用 `X-Node-Certificate-SHA256` Header 上报。面板应以大小写不敏感的
方式比较该十六进制值。
