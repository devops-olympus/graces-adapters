# Mercury Adapters

流水线引擎适配服务


## 代理设置

在Go中关闭代理，通常涉及到环境变量的设置。在使用Go模块时，Go通过`GOPROXY`环境变量来决定是否使用代理。你可以通过设置`GOPROXY`环境变量为`direct`来绕过代理，直接从源服务器获取依赖，或者完全关闭代理。

在命令行中关闭代理的方法如下：

对于**临时关闭**，仅对当前的shell会话有效：

```sh
export GOPROXY=direct
```

如果你想在**Windows**系统上关闭它（临时）：

```cmd
set GOPROXY=direct
```

如果你希望**永久关闭**代理，你可以将上述命令添加到你的shell初始化文件中，比如`.bashrc`、`.bash_profile`或`.zshrc`，或者在Windows系统中设置环境变量。

此外，`go env -w`命令也可以用来永久设置环境变量：

```sh
go env -w GOPROXY=direct
```

请注意，将代理设置为`direct`可能会导致访问某些依赖变慢，特别是当这些依赖在被墙的国家或地区时。

另外，如果你想要关闭模块支持，可以通过设置`GO111MODULE=off`来实现。这会让go命令忽略go.mod文件，并且回到go路径的工作方式：

```sh
export GO111MODULE=off
```

或者在Windows上：

```cmd
set GO111MODULE=off
```

或者永久设置：

```sh
go env -w GO111MODULE=off
```

确保在更改这些设置后，在你的开发环境中测试是否一切正常。如果你是在IDE中开发Go代码，你可能还需要在IDE的配置中更新这些设置。