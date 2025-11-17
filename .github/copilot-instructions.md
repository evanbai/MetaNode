# Copilot / AI agent instructions for MetaNode

Purpose: provide concise, repo-specific guidance so an automated coding agent can be immediately productive.

Top-level overview
- This is a small Go module (module name `MetaNode` in `go.mod`). The binary entrypoint is `main.go`.
- Learning/exercise code is organized into two folders: `stage1/` (algorithmic functions) and `stage2/` (concurrency/Go patterns).
- `stage1` contains many standalone functions (e.g. `TwoSum` in `stage1/two_sum.go`) and a runner `Stage1Valid` in `stage1/stage1_valid.go` that prints example outputs. `stage2` has similar runners (e.g. `Stage2Valid` called by `main.go`).

How this repo is typically used
- Run the project from the repo root (module imports use the module name `MetaNode`):

```sh
# build
go build ./...

# run the example runner (main.go currently calls stage2.Stage2Valid)
go run main.go
```

Guidance for code changes
- Preserve the module path in imports (`MetaNode/...`) — changing module path or packages will break local imports unless `go.mod` is updated.
- Files in `stage1/` are independent algorithm implementations. Tests and examples are executed via `stage1/stage1_valid.go`; prefer adding small functions and keeping them pure (no global state).
- Files in `stage2/` illustrate concurrency patterns (goroutines, sync.WaitGroup). Runners are named `Stage2Valid()` or `ExeGoroutine()` and are invoked from `main.go` for quick manual testing.

Conventions observed
- Package names are the directory name (e.g., `package stage1`, `package stage2`).
- Exported functions use CamelCase (e.g., `TwoSum`, `Stage1Valid`). Keep that style for any new public functions.
- Runner files named `*_valid.go` (e.g., `stage1/stage1_valid.go`, `stage2/stage2_valid.go`) print example outputs — prefer adding small runnable examples there, not in `main.go`.

Build / test / debug workflow
- Standard Go tooling works here. From repo root:

```sh
go fmt ./...       # format
go vet ./...       # static checks
go build ./...      # compile
go test ./...       # run tests (no tests present by default)
go run main.go      # run the example runner
```

- For step-debugging on Windows, standard approach is Delve (if installed): `dlv debug --headless --listen=:2345 --api-version=2` run from the repo root.

Examples agents should follow when editing
- Small change example: to add a new algorithm, create `stage1/new_algo.go` with `package stage1`, add exported function `NewAlgo(...)` and add a short invocation in `stage1/stage1_valid.go` to demonstrate outputs.
- For concurrency examples, add to `stage2/` and extend `stage2/stage2_valid.go` (or `stage2/goroutine.go`) — keep examples self-contained and deterministic where possible.

Integration points and external deps
- This repo currently uses only the Go standard library (see `go.mod` with `module MetaNode`). There are no external services, network calls, or third-party packages to mock.

Edge cases to watch for
- Changing module name or package layout will require updating `go.mod` and all internal imports.
- Runner files may print to stdout — tests or CI relying on output parsing should use pure functions instead.

If you need more context
- I reviewed `main.go`, `go.mod`, `stage1/two_sum.go`, `stage1/stage1_valid.go`, and `stage2/goroutine.go` when composing this guidance. If you want, I can add concrete TODOs or tests for key functions.

-- End of instructions

---

# Copilot / AI 代理说明 - MetaNode（中文版）

目的：提供简洁的仓库特定指导，帮助 AI 代理立即提高生产力。

顶层概览
- 这是一个小型 Go 模块（`go.mod` 中模块名为 `MetaNode`）。二进制入口点是 `main.go`。
- 学习/练习代码组织在两个文件夹中：`stage1/`（算法函数）和 `stage2/`（并发/Go 模式）。
- `stage1` 包含许多独立函数（例如 `stage1/two_sum.go` 中的 `TwoSum`）和运行器 `Stage1Valid`（在 `stage1/stage1_valid.go` 中）用于打印示例输出。`stage2` 有类似的运行器（例如 `main.go` 调用的 `Stage2Valid`）。

本仓库的典型用法
- 从仓库根目录运行项目（模块导入使用模块名 `MetaNode`）：

```sh
# 编译
go build ./...

# 运行示例运行器（main.go 当前调用 stage2.Stage2Valid）
go run main.go
```

代码修改指导
- 在导入中保留模块路径（`MetaNode/...`）— 改变模块路径或包会破坏本地导入，除非更新 `go.mod`。
- `stage1/` 中的文件是独立的算法实现。测试和示例通过 `stage1/stage1_valid.go` 执行；优先选择添加小函数并保持它们纯净（无全局状态）。
- `stage2/` 中的文件展示并发模式（goroutines、sync.WaitGroup）。运行器命名为 `Stage2Valid()` 或 `ExeGoroutine()` 并从 `main.go` 调用以进行快速手动测试。

遵循的约定
- 包名是目录名（例如 `package stage1`、`package stage2`）。
- 导出函数使用 CamelCase（例如 `TwoSum`、`Stage1Valid`）。为任何新的公共函数保持这种风格。
- 命名为 `*_valid.go` 的运行器文件（例如 `stage1/stage1_valid.go`、`stage2/stage2_valid.go`）打印示例输出 — 优先在那里添加小型可运行的示例，而不是在 `main.go` 中。

构建 / 测试 / 调试工作流
- 这里使用标准 Go 工具。从仓库根目录：

```sh
go fmt ./...       # 格式化
go vet ./...       # 静态检查
go build ./...      # 编译
go test ./...       # 运行测试（默认情况下不存在测试）
go run main.go      # 运行示例运行器
```

- 在 Windows 上进行步骤调试，标准方法是 Delve（如果已安装）：从仓库根目录运行 `dlv debug --headless --listen=:2345 --api-version=2`。

代理编辑时应遵循的示例
- 小改动示例：要添加新算法，创建 `stage1/new_algo.go`（`package stage1`），添加导出函数 `NewAlgo(...)` 并在 `stage1/stage1_valid.go` 中添加简短调用以展示输出。
- 对于并发示例，添加到 `stage2/` 并扩展 `stage2/stage2_valid.go`（或 `stage2/goroutine.go`）— 尽可能保持示例自包含且具有确定性。

集成点和外部依赖
- 此仓库当前仅使用 Go 标准库（参见 `go.mod` 中的 `module MetaNode`）。没有外部服务、网络调用或第三方包需要模拟。

需要注意的边界情况
- 改变模块名或包布局将需要更新 `go.mod` 和所有内部导入。
- 运行器文件可能会打印到标准输出 — 依赖输出解析的测试或 CI 应改用纯函数。

需要更多上下文
- 我查阅了 `main.go`、`go.mod`、`stage1/two_sum.go`、`stage1/stage1_valid.go` 和 `stage2/goroutine.go` 来编制这份指导。如果需要，我可以为关键函数添加具体的待办事项或测试。

-- 指导说明结束
