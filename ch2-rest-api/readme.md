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

如果想和大模型聊天，可以采用OpenAI风格的REST接口。创建request.json文件如下：

```json
{
  "model": "deepseek-r1:1.5b",
  "messages": [
    {"role": "user", "content": "请帮我写一封邮件给老板，告诉他我明天不能来上班。"}
  ],
  "temperature": 0.7,
  "max_tokens": 150
}
```

指定`deepseek-r1:1.5b`大模型，然后以`user`角色发出对话内容，其他参数可以忽略。然后通过curl发出请求：

```
$ curl -X POST http://localhost:11434/v1/chat/completions \
       -H "Content-Type: application/json" \
       -d @request.json
```

POST请求的API地址为`/v1/chat/completions`，通过`-H`指定请求文件为JSON格式，通过`-d`参数指定JSON文件数据。执行的结果如下：

```json
{
  "id":"chatcmpl-201",
  "object":"chat.completion",
  "created":1740558137,
  "model":"deepseek-r1:1.5b",
  "system_fingerprint":"fp_ollama",
  "choices":[
    {
      "index":0,
      "message":{
        "role":"assistant",
        "content":"...回答的内容..."
      },
      "finish_reason":"length"
    }
  ],
  "usage":{
    "prompt_tokens":18,
    "completion_tokens":150,
    "total_tokens":168
  }
}
```

完整的回答内容如下：

> `<think>`
>
> `</think>`
> 嗯，用户让我帮他写一封给他老板的邮件，告诉他他明天不能来上班。首先，我得想想用户的具体情况。可能是一个刚入职的人，但因为某些原因无法按时到岗，或者有其他安排，导致需要提前通知老板。
> 
> 接下来，我要考虑邮件的结构。通常来说，正式的商务邮件包括称呼、开头问候、说明理由、表达歉意和感谢、并留下联系方式等部分。这样看起来更专业，也更有礼貌。
> 
> 用户可能没有提到具体的原因，所以我的回应中应该留有 blanks，让他能填写具体情况。比如，是否因为家庭问题无法到岗，还是其他临时原因？这样可以让他知道邮件可以根据实际情况调整内容。
> 
> 另外，语气

OpenAI的REST接口风格已经成为事实上的标准，通过类似的方式不仅仅可以连接OpenAI服务，甚至也可以连接DeepSeek提供的服务。

