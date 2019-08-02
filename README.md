# gobcos
Golang Client For FISCO BCOS

FISCO BCOS Go语言版本的SDK，主要的功能有：

- FISCO BCOS RPC API 服务
- `Solidity`合约编译为Go合约文件

# 环境准备

- [Golang](https://golang.org/), 版本需不低于`1.12.6`，本项目采用`go module`进行包管理。
- [FISCO BCOS 2.0.0](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/), **需要提前运行** FISCO BCOS 区块链平台


# 代码运行

### RPC API 测试

此部分只要对项目的RPC API接口调用进行测试，以确定是否能顺利连接FISCO BCOS以获取区块链信息。

首先需要拉取代码：

```shell
git clone https://github.com/KasperLiu/gobcos.git
```

然后切换到远程`dev`分支：

```shell
git checkout -b eth-dev origin/eth-dev
```

进行代码测试前，请先按照实际部署的RPC URL更改`goclient_test.go`中的默认FISCO BCOS RPC连接以及群组ID：
```go
func GetClient(t *testing.T) (*Client) {
	// RPC API
	c, err := Dial("http://localhost:8545", 1) // your RPC API
	if err != nil {
		t.Fatalf("can not dial to the RPC API: %v", err)
	}
	return c
}
```
测试代码默认开启的测试函数为`GetClientVersion, GetBlockNumber, GetPBFTView`，其余函数需去除注释并更改为实际存在的数据后才能执行。

执行测试代码的命令为：

```shell
go test -v -count=1 ./client
```

### Solidity合约编译为Go文件

这部分主要包含了三个流程：

- 准备需要编译的智能合约
- 配置好相应版本的`solc`编译器
- 构建gobcos的合约编译工具`abigen`
- 编译生成go文件
- 调用go文件进行合约测试

1.提供一份简单的智能合约`Store.sol`如下：

```solidity
pragma solidity ^0.4.25;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string _version) public {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}
```

2.安装对应版本的[`solc`编译器](https://github.com/ethereum/solidity/releases/tag/v0.4.25)，目前FISCO BCOS默认的`solc`编译器版本为[0.4.25](https://github.com/ethereum/solidity/releases/tag/v0.4.25)。

```bash
solc --version
# solc, the solidity compiler commandline interface
# Version: 0.4.25+commit.59dbf8f1.Linux.g++
```

3.构建`gobcos`的代码生成工具`abigen`

```bash
git clone https://github.com/KasperLiu/gobcos.git # 下载gobcos代码
git checkout -b eth-dev origin/eth-dev # 切分分支
go build ./cmd/abigen # 编译生成abigen工具
```

执行命令后，检查目录下是否存在`abigen`

4.编译生成go文件，首先需要先利用`solc`将合约文件生成`abi`和`bin`文件，以前面所提供的`Store.sol`为例：

```bash
solc --bin -o ./ Store.sol
sloc --abi -o ./ Store.sol
```

此时`Store.sol`目录下会生成`Store.bin`和`Store.abi`。此时利用`abigen`工具将`Store.bin`和`Store.abi`转换成`Store.go`：

```bash
abigen --bin=Store.bin --abi=Store.abi --pkg=store --out=Store.go
```

