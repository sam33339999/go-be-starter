# air with golang

> 熱刷新(Hot Reload)

- [air_example.toml](https://github.com/cosmtrek/air/blob/master/air_example.toml)
- [參考配置檔案翻譯](https://www.cnblogs.com/wylshkjj/p/16546583.html)

## 初始化
```
air init
# 目錄下會產生對應的 .air.toml
# 隨後調整自己所需要參數來使專案熱刷新(fanotify)


# 可以透過下列指令來顯示更詳細的刷新資訊
air -d 
```

### 根據目前專案使用，運行時需要帶引數 (arguments)
- [Add support for runtime arguments](https://github.com/cosmtrek/air/pull/166/files)

```toml
args_bin = ["app:serve"]
```



### 配置檔案相關
```toml
# [Air](https://github.com/cosmtrek/air) TOML 格式的配置文件

# 工作目录
# 使用. 或绝对路径, 请注意以下目录tmp_dir必须在根目录root下。
root = "."
tmp_dir = "tmp"

[build]
# 只需要写你平常编译使用的shell命令。也可以使用`make`
# Windows平台示例: cmd = "go build -o tmp\main.exe ."
cmd = "go build -o ./tmp/main ."
# 由`cmd`命令得到的二进制文件名
# Windows平台示例：bin = "tmp\main.exe"
bin = "tmp/main"
# 自定义执行程序的命令，可以添加额外的编译标识例如添加 GIN_MODE=release
# Windows平台示例：full_bin = "tmp\main.exe"
full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
# 监听扩展名的文件
include_ext = ["go", "tpl", "tmpl", "html"]
# 忽略（不监听）文件的扩展名或目录
exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]
# 监听指定目录的文件
include_dir = []
# 忽略（不监听）指定文件
exclude_file = []
# 忽略符合通过正则匹配到的文件
exclude_regex = ["_test\\.go"]
# 忽略未进行修改的文件
exclude_unchanged = true
# 按照目录的符号链接
follow_symlink = true
# 这个日志文件放在你的`tmp_dir`中
log = "air.log"
# 如果文件更改过于频繁，则没有必要在每次更改时都触发构建。可以设置触发构建的延迟时间/毫秒
delay = 1000 # ms
# 发生构建错误时，停止运行旧的二进制文件
stop_on_error = true
# 杀死进程前发送中断信号(Windows不支持)
send_interrupt = false
# 发送中断信号后延迟时间/毫秒
kill_delay = 500 # ms
# 在运行二进制文件时添加额外的参数 (bin/full_bin)。将运行“./tmp/main hello world”
args_bin = ["hello", "world"]

[log]
# 显示日志时间
time = false

[color]
# 自定义每个部分的颜色。如果未找到颜色，请使用原始应用程序日志。
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# 退出时删除 tmp 目录
clean_on_exit = true
```