# 4. 大模型文件结构

前文已经通过ollama命令下载、执行大模型，并且构建了一个大模型的应用。本章尝试分析大模型文件结构。

## 4.1 大模型下载到哪里了？

可以通过`ollama pull`命令下载大模型，以下还是deepseek的例子：

```
$ ollama pull deepseek-r1:1.5b
pulling manifest 
pulling aabd4debf0c8...   3% ▕                ▏  27 MB/1.1 GB  3.0 MB/s   5m57s
...
```

下载完成后可以通过`ollama show`查看大模型的一些信息：

```
$ ollama show deepseek-r1:1.5b
  Model
    architecture        qwen2
    parameters          1.8B
    context length      131072
    embedding length    1536
    quantization        Q4_K_M

  Parameters
    stop    "<｜begin▁of▁sentence｜>"
    stop    "<｜end▁of▁sentence｜>"
    stop    "<｜User｜>"
    stop    "<｜Assistant｜>"

  License
    MIT License
    Copyright (c) 2023 DeepSeek
```

根据Ollama文档，不同操作系统下载的目录类似：

- macOS: `~/.ollama/models`
- Linux: `/usr/share/ollama/.ollama/models`
- Windows: `C:\Users\%username%\.ollama\models`

下面是macOS系统下载deepseek后的文件：

```
$ tree ~/.ollama/models
/Users/chai/.ollama/models
├── blobs
│   ├── sha256-369ca498f347f710d068cbb38bf0b8692dd3fa30f30ca2ff755e211c94768150
│   ├── sha256-6e4c38e1172f42fdbff13edf9a7a017679fb82b0fde415a3e8b3c31c6ed4a4e4
│   ├── sha256-a85fe2a2e58e2426116d3686dfdc1a6ea58640c1e684069976aa730be6c1fa01
│   ├── sha256-aabd4debf0c8f08881923f2c25fc0fdeed24435271c2b3e92c4af36704040dbc
│   └── sha256-f4d24e9138dd4603380add165d2b0d970bef471fac194b436ebd50e6147c6588
└── manifests
    └── registry.ollama.ai
        └── library
            └── deepseek-r1
                └── 1.5b

6 directories, 6 files
```

其中`blobs`目录下是以SHA256命名保存具体数据，`manifests`是各个模型的主文件。`manifests`目录下的文件以`{host}/{namespace}/{model}/{tag}`形式的路径组织，比如`registry.ollama.ai`是Ollama自己的托管平台地址、`library`对应大模型的名字空间、`deepseek-r1:1.5b`则是具体的模型和Tag版本号

## 4.2 manifest文件结构

通过以下命令查看：

```
$ cat ~/.ollama/models/manifests/registry.ollama.ai/library/deepseek-r1/1.5b
...
```

这是JSON格式的数据，格式化后如下：

```json
{
  "schemaVersion": 2,
  "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
  "config": {
    "mediaType": "application/vnd.docker.container.image.v1+json",
    "digest": "sha256:a85fe2a2e58e2426116d3686dfdc1a6ea58640c1e684069976aa730be6c1fa01",
    "size": 487
  },
  "layers": [
    {
      "mediaType": "application/vnd.ollama.image.model",
      "digest": "sha256:aabd4debf0c8f08881923f2c25fc0fdeed24435271c2b3e92c4af36704040dbc",
      "size": 1117320512
    },
    {
      "mediaType": "application/vnd.ollama.image.template",
      "digest": "sha256:369ca498f347f710d068cbb38bf0b8692dd3fa30f30ca2ff755e211c94768150",
      "size": 387
    },
    {
      "mediaType": "application/vnd.ollama.image.license",
      "digest": "sha256:6e4c38e1172f42fdbff13edf9a7a017679fb82b0fde415a3e8b3c31c6ed4a4e4",
      "size": 1065
    },
    {
      "mediaType": "application/vnd.ollama.image.params",
      "digest": "sha256:f4d24e9138dd4603380add165d2b0d970bef471fac194b436ebd50e6147c6588",
      "size": 148
    }
  ]
}
```

这是类似Dockr的分层组织的文件系统。其中有4个包含sha256信息的文件在blobs目录保存：

```
$ ls ~/.ollama/models/blobs
sha256-369ca498f347f710d068cbb38bf0b8692dd3fa30f30ca2ff755e211c94768150
sha256-6e4c38e1172f42fdbff13edf9a7a017679fb82b0fde415a3e8b3c31c6ed4a4e4
sha256-a85fe2a2e58e2426116d3686dfdc1a6ea58640c1e684069976aa730be6c1fa01
sha256-aabd4debf0c8f08881923f2c25fc0fdeed24435271c2b3e92c4af36704040dbc
sha256-f4d24e9138dd4603380add165d2b0d970bef471fac194b436ebd50e6147c6588
$
```

## 4.3 模板文件

TODO
