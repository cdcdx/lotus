<p align="center">
  <a href="https://docs.filecoin.io/" title="Filecoin Docs">
    <img src="documentation/images/lotus_logo_h.png" alt="Project Lotus Logo" width="244" />
  </a>
</p>

<h1 align="center">Project Lotus - 莲</h1>

<p align="center">
  <a href="https://circleci.com/gh/filecoin-project/lotus"><img src="https://circleci.com/gh/filecoin-project/lotus.svg?style=svg"></a>
  <a href="https://codecov.io/gh/filecoin-project/lotus"><img src="https://codecov.io/gh/filecoin-project/lotus/branch/master/graph/badge.svg"></a>
  <a href="https://goreportcard.com/report/github.com/filecoin-project/lotus"><img src="https://goreportcard.com/badge/github.com/filecoin-project/lotus" /></a>  
  <a href=""><img src="https://img.shields.io/badge/golang-%3E%3D1.16-blue.svg" /></a>
  <br>
</p>


## 钱包加密方案

### Instructions
  - 前段时间有矿工私钥被盗,owner被人篡改,现把钱包加密部分整理出来,源码开源,欢迎大家整合;
  - 提交issue一起完善: [https://github.com/cdcdx/lotus](https://github.com/cdcdx/lotus)
  - 加入tg一起讨论: [https://t.me/fil_wallet_security](https://t.me/fil_wallet_security)

### Suggest
  - 采用加密方案保障安全的同时,还需加强内部管理,控制设备访问权限;
  - 本方案使用了AES加密,一定程度保障了资金和私钥安全,但还是有泄露的风险,请知悉;
  - 源码编译前记得修改walletSaltPwd;

### Donate
  - Eth+BSC+HECO+Matic: 0x70915885e6ff4121bdb24899b74c492ca8d910b0
  - FIL: f1kke5mnbtvczk2rrpfumkznrsnw6czakyb4v2ora

### Do
  - [x] 兼容适配:
    - [x] 兼容官方lotus节点api;
    - [x] 适配lotus和lotus-wallet两种方案;
  - [x] 钱包密码管理:
    - [x] 增加、修改、清理密码
    - [x] 普通钱包与加密钱包互换;
  - [x] 钱包消息管理:
    - [x] 保障资金安全:加密钱包转账操作需验证密码;
    - [x] 保护私钥安全:加密钱包导出私钥和删除需验证密码;
    - [x] 密封消息不受影响:加密钱包密封消息不需验证密码;

### ToDo
  - [ ] 3次失败锁定24小时:操作连续3次验证密码错误,锁定钱包24小时;
  - [ ] 钱包加密增强;

### Examples
  - 钱包密码管理:
```shell
    lotus wallet addpasswd     ##增加密码
    lotus wallet resetpasswd   ##修改密码
    lotus wallet clearpasswd   ##清理密码
    lotus wallet encrypt <f1xxx/f3xxx>  ##普通钱包->加密钱包
    lotus wallet decrypt <f1xxx/f3xxx>  ##加密钱包->普通钱包
```
  - 钱包消息管理:
```shell
    lotus send f1kke5mnbtvczk2rrpfumkznrsnw6czakyb4v2ora 10  ##转账需要输入密码
    lotus wallet export f1kke5mnbtvczk2rrpfumkznrsnw6czakyb4v2ora  ##导出私钥需要输入密码
    lotus wallet delete f1kke5mnbtvczk2rrpfumkznrsnw6czakyb4v2ora  ##删除钱包需要输入密码
```

### Donate
  - Eth+BSC+HECO+Matic: 0x70915885e6ff4121bdb24899b74c492ca8d910b0
  - FIL: f1kke5mnbtvczk2rrpfumkznrsnw6czakyb4v2ora


## License

Dual-licensed under [MIT](https://github.com/filecoin-project/lotus/blob/master/LICENSE-MIT) + [Apache 2.0](https://github.com/filecoin-project/lotus/blob/master/LICENSE-APACHE)
