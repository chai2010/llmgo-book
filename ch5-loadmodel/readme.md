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

其中最重要是`-model`参数指定大模型文件。通过以下命令运行`deepseek-r1:1.5b`大模型：

```
$ go run cmd/runner/main.go -model=$HOME/.ollama/models/blobs/sha256-aabd4debf0c8f08881923f2c25fc0fdeed24435271c2b3e92c4af36704040dbc
time=2025-03-05T09:58:06.228+08:00 level=INFO source=runner.go:932 msg="starting go runner"
time=2025-03-05T09:58:06.231+08:00 level=INFO source=runner.go:935 msg=system info="CPU : SSE3 = 1 | SSSE3 = 1 | LLAMAFILE = 1 | cgo(clang)" threads=4
time=2025-03-05T09:58:06.232+08:00 level=INFO source=.:0 msg="Server listening on 127.0.0.1:8080"
llama_model_loader: loaded meta data with 26 key-value pairs and 339 tensors from /Users/chai/.ollama/models/blobs/sha256-aabd4debf0c8f08881923f2c25fc0fdeed24435271c2b3e92c4af36704040dbc (version GGUF V3 (latest))
llama_model_loader: Dumping metadata keys/values. Note: KV overrides do not apply in this output.
llama_model_loader: - kv   0:                       general.architecture str              = qwen2
llama_model_loader: - kv   1:                               general.type str              = model
llama_model_loader: - kv   2:                               general.name str              = DeepSeek R1 Distill Qwen 1.5B
llama_model_loader: - kv   3:                           general.basename str              = DeepSeek-R1-Distill-Qwen
llama_model_loader: - kv   4:                         general.size_label str              = 1.5B
llama_model_loader: - kv   5:                          qwen2.block_count u32              = 28
llama_model_loader: - kv   6:                       qwen2.context_length u32              = 131072
llama_model_loader: - kv   7:                     qwen2.embedding_length u32              = 1536
llama_model_loader: - kv   8:                  qwen2.feed_forward_length u32              = 8960
llama_model_loader: - kv   9:                 qwen2.attention.head_count u32              = 12
llama_model_loader: - kv  10:              qwen2.attention.head_count_kv u32              = 2
llama_model_loader: - kv  11:                       qwen2.rope.freq_base f32              = 10000.000000
llama_model_loader: - kv  12:     qwen2.attention.layer_norm_rms_epsilon f32              = 0.000001
llama_model_loader: - kv  13:                          general.file_type u32              = 15
llama_model_loader: - kv  14:                       tokenizer.ggml.model str              = gpt2
llama_model_loader: - kv  15:                         tokenizer.ggml.pre str              = qwen2
llama_model_loader: - kv  16:                      tokenizer.ggml.tokens arr[str,151936]  = ["!", "\"", "#", "$", "%", "&", "'", ...
llama_model_loader: - kv  17:                  tokenizer.ggml.token_type arr[i32,151936]  = [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, ...
llama_model_loader: - kv  18:                      tokenizer.ggml.merges arr[str,151387]  = ["Ġ Ġ", "ĠĠ ĠĠ", "i n", "Ġ t",...
llama_model_loader: - kv  19:                tokenizer.ggml.bos_token_id u32              = 151646
llama_model_loader: - kv  20:                tokenizer.ggml.eos_token_id u32              = 151643
llama_model_loader: - kv  21:            tokenizer.ggml.padding_token_id u32              = 151643
llama_model_loader: - kv  22:               tokenizer.ggml.add_bos_token bool             = true
llama_model_loader: - kv  23:               tokenizer.ggml.add_eos_token bool             = false
llama_model_loader: - kv  24:                    tokenizer.chat_template str              = {% if not add_generation_prompt is de...
llama_model_loader: - kv  25:               general.quantization_version u32              = 2
llama_model_loader: - type  f32:  141 tensors
llama_model_loader: - type q4_K:  169 tensors
llama_model_loader: - type q6_K:   29 tensors
llm_load_vocab: special_eos_id is not in special_eog_ids - the tokenizer config may be incorrect
llm_load_vocab: special tokens cache size = 22
llm_load_vocab: token to piece cache size = 0.9310 MB
llm_load_print_meta: format           = GGUF V3 (latest)
llm_load_print_meta: arch             = qwen2
llm_load_print_meta: vocab type       = BPE
llm_load_print_meta: n_vocab          = 151936
llm_load_print_meta: n_merges         = 151387
llm_load_print_meta: vocab_only       = 0
llm_load_print_meta: n_ctx_train      = 131072
llm_load_print_meta: n_embd           = 1536
llm_load_print_meta: n_layer          = 28
llm_load_print_meta: n_head           = 12
llm_load_print_meta: n_head_kv        = 2
llm_load_print_meta: n_rot            = 128
llm_load_print_meta: n_swa            = 0
llm_load_print_meta: n_embd_head_k    = 128
llm_load_print_meta: n_embd_head_v    = 128
llm_load_print_meta: n_gqa            = 6
llm_load_print_meta: n_embd_k_gqa     = 256
llm_load_print_meta: n_embd_v_gqa     = 256
llm_load_print_meta: f_norm_eps       = 0.0e+00
llm_load_print_meta: f_norm_rms_eps   = 1.0e-06
llm_load_print_meta: f_clamp_kqv      = 0.0e+00
llm_load_print_meta: f_max_alibi_bias = 0.0e+00
llm_load_print_meta: f_logit_scale    = 0.0e+00
llm_load_print_meta: n_ff             = 8960
llm_load_print_meta: n_expert         = 0
llm_load_print_meta: n_expert_used    = 0
llm_load_print_meta: causal attn      = 1
llm_load_print_meta: pooling type     = 0
llm_load_print_meta: rope type        = 2
llm_load_print_meta: rope scaling     = linear
llm_load_print_meta: freq_base_train  = 10000.0
llm_load_print_meta: freq_scale_train = 1
llm_load_print_meta: n_ctx_orig_yarn  = 131072
llm_load_print_meta: rope_finetuned   = unknown
llm_load_print_meta: ssm_d_conv       = 0
llm_load_print_meta: ssm_d_inner      = 0
llm_load_print_meta: ssm_d_state      = 0
llm_load_print_meta: ssm_dt_rank      = 0
llm_load_print_meta: ssm_dt_b_c_rms   = 0
llm_load_print_meta: model type       = 1.5B
llm_load_print_meta: model ftype      = Q4_K - Medium
llm_load_print_meta: model params     = 1.78 B
llm_load_print_meta: model size       = 1.04 GiB (5.00 BPW) 
llm_load_print_meta: general.name     = DeepSeek R1 Distill Qwen 1.5B
llm_load_print_meta: BOS token        = 151646 '<｜begin▁of▁sentence｜>'
llm_load_print_meta: EOS token        = 151643 '<｜end▁of▁sentence｜>'
llm_load_print_meta: EOT token        = 151643 '<｜end▁of▁sentence｜>'
llm_load_print_meta: PAD token        = 151643 '<｜end▁of▁sentence｜>'
llm_load_print_meta: LF token         = 148848 'ÄĬ'
llm_load_print_meta: FIM PRE token    = 151659 '<|fim_prefix|>'
llm_load_print_meta: FIM SUF token    = 151661 '<|fim_suffix|>'
llm_load_print_meta: FIM MID token    = 151660 '<|fim_middle|>'
llm_load_print_meta: FIM PAD token    = 151662 '<|fim_pad|>'
llm_load_print_meta: FIM REP token    = 151663 '<|repo_name|>'
llm_load_print_meta: FIM SEP token    = 151664 '<|file_sep|>'
llm_load_print_meta: EOG token        = 151643 '<｜end▁of▁sentence｜>'
llm_load_print_meta: EOG token        = 151662 '<|fim_pad|>'
llm_load_print_meta: EOG token        = 151663 '<|repo_name|>'
llm_load_print_meta: EOG token        = 151664 '<|file_sep|>'
llm_load_print_meta: max token length = 256
llm_load_tensors:   CPU_Mapped model buffer size =  1059.89 MiB
llama_new_context_with_model: n_seq_max     = 1
llama_new_context_with_model: n_ctx         = 2048
llama_new_context_with_model: n_ctx_per_seq = 2048
llama_new_context_with_model: n_batch       = 512
llama_new_context_with_model: n_ubatch      = 512
llama_new_context_with_model: flash_attn    = 0
llama_new_context_with_model: freq_base     = 10000.0
llama_new_context_with_model: freq_scale    = 1
llama_new_context_with_model: n_ctx_per_seq (2048) < n_ctx_train (131072) -- the full capacity of the model will not be utilized
llama_kv_cache_init: kv_size = 2048, offload = 1, type_k = 'f16', type_v = 'f16', n_layer = 28, can_shift = 1
llama_kv_cache_init:        CPU KV buffer size =    56.00 MiB
llama_new_context_with_model: KV self size  =   56.00 MiB, K (f16):   28.00 MiB, V (f16):   28.00 MiB
llama_new_context_with_model:        CPU  output buffer size =     0.59 MiB
llama_new_context_with_model:        CPU compute buffer size =   299.75 MiB
llama_new_context_with_model: graph nodes  = 986
llama_new_context_with_model: graph splits = 1
```

启动后可以看到一些和命令行参数帮助信息对应的参数，比较重要的是端口信息。现在我们可以通过8080端口来测试服务是否正常：

```
$ curl http://localhost:8080/health
{"status":"ok","progress":1}
```

也可以访问`/completion`补全接口：

```
$ curl -X POST -H "Content-Type: application/json" -d '{"prompt": "hi"}' http://localhost:8080/completion
{"content":",","stop":false,"timings":{"predicted_n":0,"predicted_ms":0,"prompt_n":0,"prompt_ms":0}}
{"content":" I","stop":false,"timings":{"predicted_n":0,"predicted_ms":0,"prompt_n":0,"prompt_ms":0}}
{"content":"'m","stop":false,"timings":{"predicted_n":0,"predicted_ms":0,"prompt_n":0,"prompt_ms":0}}
{"content":" trying","stop":false,"timings":{"predicted_n":0,"predicted_ms":0,"prompt_n":0,"prompt_ms":0}}
...
```

补全会产生很多结果，用户需要根据实际情况决定何时结束推理。

此外还有`/embedding`接口：

```
$ curl -X POST -H "Content-Type: application/json" -d '{"prompt": "turn me into an embedding"}' http://localhost:8080/embedding
{"embedding":[1.6663613,-2.8487935,1.8253331,-1.123977,...,0.95781755]}
```

## 5.2 通过程序加载模型

TODO

