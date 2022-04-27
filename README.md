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


## 分布式Miner方案
----
  本方案改自[moran](https://github.com/moran666666)开源方案

### Instructions
  - 1.分布式miner需要将原来的miner目录文件复制多份。（若分布式miner在同一机器，监听端口不能相同）
  - 2.分布式miner都需要挂载落盘目录，需要保持所有miner都能访问落盘目录，且保持挂盘路径一致。
  - 3.扇区分配由alloce-miner控制，第一次启动会在alloce-miner目录下会创建sectorid文件，初始值为0，里面记录矿工最新扇区编号。
    - 3.1 如果miner之前没有密封过扇区，不用修改sectorid文件初始值0；
    - 3.2 如果miner之前密封过扇区现改造成分布式，需要将sectorid文件初始值0修改成(使用过扇区ID最大值+1)；
  - 4.密封任务在seal-miner上执行。

### Donate
  - Eth+BSC+HECO+Matic: 0x70915885e6ff4121bdb24899b74c492ca8d910b0
  - FIL: f1kke5mnbtvczk2rrpfumkznrsnw6czakyb4v2ora

### Single deployment
  - 密封和post不冲突;
  - wdpost和wnpost不冲突:用一半核心和所有显卡参与wdpost;用另一半核心参与wnpost;
```shell
  ## wdminer - 用一半核心和所有显卡参与wdpost
  mkdir -p /lotus_data/tmp/tmpwd 
  export TMPDIR=/lotus_data/tmp/tmpwd 
  export LOTUS_MINER_PATH=/lotus_data/lotuswdminer
  nohup taskset -c 0-31 lotus-miner run --wdpost=true --wnpost=false --p2p=false >> /var/log/lotus/wdminer.log 2>&1 & 

  ## wnminer - 用另一半核心参与wnpost
  mkdir -p /lotus_data/tmp/tmpwn 
  export TMPDIR=/lotus_data/tmp/tmpwn 
  export LOTUS_MINER_PATH=/lotus_data/lotuswnminer 
  nohup taskset -c 32-63 lotus-miner run --wdpost=false --wnpost=true --p2p=false --enable-gpu-proving=false >> /var/log/lotus/wnminer.log 2>&1 & 

  ## alloceminer - 扇区分配服务
  export LOTUS_MINER_PATH=/lotus_data/lotusalloceminer 
  nohup lotus-miner run --wdpost=false --wnpost=false --p2p=false --enable-gpu-proving=false --sctype=alloce --sclisten=172.12.1.3:1357 >> /var/log/lotus/alloceminer.log 2>&1 & 

  ## sealminer - 密封扇区服务
  export LOTUS_MINER_PATH=/lotus_data/lotussealminer 
  nohup lotus-miner run --wdpost=false --wnpost=false --p2p=false --enable-gpu-proving=false --sctype=get --sclisten=172.12.1.3:1357 >> /var/log/lotus/sealminer.log 2>&1 & 
```

### Multiple deployments
  - 密封和post不冲突;
  - wdpost和wnpost不冲突;
  - 多台sealminer参与密封,日增算力无上限;
```shell
  ## wdminer
  export LOTUS_MINER_PATH=/lotus_data/lotuswdminer
  nohup lotus-miner run --wdpost=true --wnpost=false --p2p=false >> /var/log/lotus/wdminer.log 2>&1 & 

  ## wnminer
  export LOTUS_MINER_PATH=/lotus_data/lotuswnminer 
  nohup lotus-miner run --wdpost=false --wnpost=true --p2p=false >> /var/log/lotus/wnminer.log 2>&1 & 

  ## alloceminer - 扇区分配服务
  export LOTUS_MINER_PATH=/lotus_data/lotusalloceminer 
  nohup lotus-miner run --wdpost=false --wnpost=false --p2p=false --sctype=alloce --sclisten=172.12.1.3:1357 >> /var/log/lotus/alloceminer.log 2>&1 & 

  ## sealminer - 密封扇区服务
  export LOTUS_MINER_PATH=/lotus_data/lotussealminer 
  nohup lotus-miner run --wdpost=false --wnpost=false --p2p=false --sctype=get --sclisten=172.12.1.3:1357 >> /var/log/lotus/sealminer.log 2>&1 & 
```

### Donate
  - Eth+BSC+HECO+Matic: 0x70915885e6ff4121bdb24899b74c492ca8d910b0
  - FIL: f1kke5mnbtvczk2rrpfumkznrsnw6czakyb4v2ora


## License

Dual-licensed under [MIT](https://github.com/filecoin-project/lotus/blob/master/LICENSE-MIT) + [Apache 2.0](https://github.com/filecoin-project/lotus/blob/master/LICENSE-APACHE)
