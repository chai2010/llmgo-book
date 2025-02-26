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

## 2.2 通过API和大模型聊天

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

> 嗯，用户让我帮他写一封给他老板的邮件，告诉他他明天不能来上班。首先，我得想想用户的具体情况。可能是一个刚入职的人，但因为某些原因无法按时到岗，或者有其他安排，导致需要提前通知老板。
> 
> 接下来，我要考虑邮件的结构。通常来说，正式的商务邮件包括称呼、开头问候、说明理由、表达歉意和感谢、并留下联系方式等部分。这样看起来更专业，也更有礼貌。
> 
> 用户可能没有提到具体的原因，所以我的回应中应该留有 blanks，让他能填写具体情况。比如，是否因为家庭问题无法到岗，还是其他临时原因？这样可以让他知道邮件可以根据实际情况调整内容。
> 
> 另外，语气

OpenAI的REST接口风格已经成为事实上的标准，通过类似的方式不仅仅可以连接OpenAI服务，甚至也可以连接DeepSeek提供的服务。

## 2.3 聊天的上下文

要通过`curl`实现多轮聊天，你需要在每次请求时将之前的对话历史（即消息）传递给模型，以便模型能够基于上下文生成合理的回答。`ollama` 的 API 会将整个对话历史作为请求的一部分，以便生成基于上下文的响应。

在每次与模型对话时，你将之前的用户和助手消息一并发送给 API，这样模型就能理解对话的上下文。假设你已经启动了 `ollama` 的 API，并且模型在 `localhost:11434` 上运行。以下是一个多轮聊天的示例。

### 2.3.1 第一次对话

在第一次发送请求时，你的 `request.json` 文件可能是这样的：

```json
{
  "model": "deepseek-r1:1.5b",
  "messages": [
    {"role": "user", "content": "你好，今天怎么样？"}
  ]
}
```

第一次请求时，我们会发送一个用户的消息：

```bash
$ curl -X POST http://localhost:11434/v1/chat/completions \
    -H "Content-Type: application/json" \
    -d @request.json
```

假设模型的回应如下：

```json
{
  "id":"chatcmpl-507",
  "object":"chat.completion",
  "created":1740572517,
  "model":"deepseek-r1:1.5b",
  "system_fingerprint":"fp_ollama",
  "choices":[
    {
      "index":0,
      "message":{
        "role":"assistant",
        "content":"您好！感觉怎么样呢？如果您有任何问题或需要帮助的地方，请随时告诉我。我会尽力为您提供更好的服务。"
      },
      "finish_reason":"stop"
    }
  ],
  "usage":{
    "prompt_tokens":8,
    "completion_tokens":28,
    "total_tokens":36
  }
}
```

### 2.3.2 进行第二轮对话

在第二轮对话中，你需要将用户的消息和模型的回应一并添加到 `messages` 数组中。例如，用户接着问：“你做了什么？”：

更新后的 `request.json` 文件：

```json
{
  "model": "deepseek-r1:1.5b",
  "messages": [
    {"role": "user", "content": "你好，今天怎么样？"},
    {"role": "assistant", "content": "您好！感觉怎么样呢？如果您有任何问题或需要帮助的地方，请随时告诉我。我会尽力为您提供更好的服务。！"},
    {"role": "user", "content": "你做了什么？"}
  ]
}
```

然后，发送第二轮请求：

```bash
$ curl -X POST http://localhost:11434/v1/chat/completions \
    -H "Content-Type: application/json" \
    -d @request.json
```

模型将基于所有历史消息（包括助手的回应）来生成回答。例如，模型可能会回应：

```json
{
  "id":"chatcmpl-208",
  "object":"chat.completion",
  "created":1740572681,
  "model":"deepseek-r1:1.5b",
  "system_fingerprint":"fp_ollama",
  "choices":[
    {
      "index":0,
      "message":{
        "role":"assistant",
        "content":"好的，现在我们可以先专注于当前的问题和相关回复。如果还有其他问题需要我帮忙解答的，请随时告诉我。我们共同进步，一起成长！"
      },
      "finish_reason":"stop"
    }
  ],
  "usage":{
    "prompt_tokens":39,
    "completion_tokens":37,
    "total_tokens":76
  }
}
```

### 2.3.3 第三轮对话

如果用户继续提问，可以像之前一样更新 `request.json` 文件，逐步将对话历史传递给模型。更新后的 `request.json` 文件（第三轮对话）：

```json
{
  "model": "deepseek-r1:1.5b",
  "messages": [
    {"role": "user", "content": "你好，今天怎么样？"},
    {"role": "assistant", "content": "您好！感觉怎么样呢？如果您有任何问题或需要帮助的地方，请随时告诉我。我会尽力为您提供更好的服务。！"},
    {"role": "user", "content": "你做了什么？"},
    {"role": "assistant", "content": "好的，现在我们可以先专注于当前的问题和相关回复。如果还有其他问题需要我帮忙解答的，请随时告诉我。我们共同进步，一起成长！"},
    {"role": "user", "content": "那我们继续聊吧！"}
  ]
}
```

然后再发送：

```bash
$ curl -X POST http://localhost:11434/v1/chat/completions \
    -H "Content-Type: application/json" \
    -d @request.json
```

返回的结果如下：

```json
{
  "id":"chatcmpl-134",
  "object":"chat.completion",
  "created":1740573171,
  "model":"deepseek-r1:1.5b",
  "system_fingerprint":"fp_ollama",
  "choices":[
    {
      "index":0,
      "message":{
        "role":"assistant",
        "content":"您好！我是由中国的深度求索（DeepSeek）公司开发的智能助手DeepSeek-R1。有关模型和产品的详细内容请参考官方文档。"
      },
      "finish_reason":"stop"
    }
  ],
  "usage":{
    "prompt_tokens":80,
    "completion_tokens":38,
    "total_tokens":118
  }
}
```

## 2.4 上下文校准

因为采用的是DeepSeek小模型测试，回答的效果并不太好，读者可以自行更换为更大的模型测试。不过在在进行多轮对话中，我们也可以调整或构建新的对话上下文。比如我们之间以下面的数据进入另一个版本的对话中：

```json
{
  "model": "deepseek-r1:1.5b",
  "messages": [
    {"role": "user", "content": "你好，今天怎么样？"},
    {"role": "assistant", "content": "你好呀！我今天很好，谢谢！"},
    {"role": "user", "content": "你做了什么？"},
    {"role": "assistant", "content": "我今天一直在和你聊天！"},
    {"role": "user", "content": "那我们继续聊吧！"}
  ],
  "temperature": 0.7,
  "max_tokens": 150
}
```

重新请求后返回的结果如下：

```json
{
  "id":"chatcmpl-891",
  "object":"chat.completion",
  "created":1740573640,
  "model":"deepseek-r1:1.5b",
  "system_fingerprint":"fp_ollama",
  "choices":[
    {
      "index":0,
      "message":{
        "role":"assistant",
        "content":"您好！很高兴能与您保持联系。希望在未来的日子里能够共同进步、共事、共患难，一起克服任何挑战，享受生活乐趣。如果还有其他需求或问题，请随时告诉我哦！"
      },
      "finish_reason":"stop"
    }
  ],
  "usage":{
    "prompt_tokens":40,
    "completion_tokens":50,
    "total_tokens":90
  }
}
```

从这里可以看出，基于大模型的应用可以通过调整真实的用户聊天的上下文来影响最终的回答。有时候可以改进体验，有时候也可能被用于构建信息茧房。

