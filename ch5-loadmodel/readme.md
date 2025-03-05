# 5. 手动加载模型文件

在上一章我们已经找到了下载到本地的大模型文件，现在我们尝试以手工命令行方式加载。

## 5.1 runner命令行加载

之前的例子我们是通过Ollama应用来使用大模型，其底层是通过一个"runner"程序启动的推理引擎。“runner”代码在Ollama主仓库的`cmd/runner`目录下，是Go语言和C/C++混合编写的程序。

现在我们尝试从命令行启动，进入Ollama仓库的`cmd/runner`子目录，执行以下命令：

```
$ go run cmd/runner/main.go -h
Runner usage
  -batch-size int
        Batch size (default 512)
  -ctx-size int
        Context (or KV cache) size (default 2048)
  -flash-attn
        Enable flash attention
  -kv-cache-type string
        quantization type for KV cache (default: f16)
  -lora value
        Path to lora layer file (can be specified multiple times)
  -main-gpu int
        Main GPU
  -mlock
        force system to keep model in RAM rather than swapping or compressing
  -mmproj string
        Path to projector binary file
  -model string
        Path to model binary file
  -multiuser-cache
        optimize input cache algorithm for multiple users
  -n-gpu-layers int
        Number of layers to offload to GPU
  -no-mmap
        do not memory-map model (slower load but may reduce pageouts if not using mlock)
  -parallel int
        Number of sequences to handle simultaneously (default 1)
  -port int
        Port to expose the server on (default 8080)
  -tensor-split string
        fraction of the model to offload to each GPU, comma-separated list of proportions
  -threads int
        Number of threads to use during generation (default 4)
  -verbose
        verbose output (default: disabled)
$
```

其中最重要有2个参数：`-model`指定大模型五年级，`-port`指定大模型服务的端口。


TODO

