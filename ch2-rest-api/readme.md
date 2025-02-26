# 2. 访问REST服务

因为执行完整版的大语言模型通常需要大量的硬件资源，因此大语言模型一般会以服务的形式运行在特殊的硬件上，然后通过REST协议提高服务。本书的目的不是为了学习REST协议，但是通过他可以了解大语言模型对外提供服务的方式，达到一定的去魅的目的。

## 2.1 查看模型列表

Ollama不仅仅是命令，他一般是以一个独立的应用安装，安装后打开就默认在11434端口启动了服务。比如`/api/tags`接口可以返回当前安装的模型（和`ollama list`命令类似）：

```
$ curl http://localhost:11434/api/tags
{"models":[{"name":"deepseek-r1:1.5b","model":"deepseek-r1:1.5b", ...
```

REST服务返回的是JSON格式的数据，格式化后完整的数据如下： 

```json
{
  "models": [
    {
      "name":"deepseek-r1:1.5b",
      "model":"deepseek-r1:1.5b",
      "modified_at":"2025-02-09T16:23:27.221768541+08:00",
      "size":1117322599,
      "digest":"a42b25d8c10a841bd24724309898ae851466696a7d7f3a0a408b895538ccbc96",
      "details": {
        "parent_model":"",
        "format":"gguf",
        "family":"qwen2",
        "families":["qwen2"],
        "parameter_size":"1.8B",
        "quantization_level":"Q4_K_M"
      }
    }
  ]
}
```

## 2.2 聊天API


TODO

