package types

// ExecDetail 保存单个测试用例的执行结果
type ExecDetail struct {
	Output           string `json:"output"`            // 程序输出
	TimeMilliseconds int64  `json:"time_milliseconds"` // 运行时间，单位为毫秒（Milliseconds）
	MemoryUsage      int64  `json:"memory_usage_kb"`   // 占用内存，单位为KB
	Timeout          bool   `json:"timeout"`           // 是否超时
	MemoryOut        bool   `json:"memory_out"`        // 是否内存超限
	RuntimeError     bool   `json:"runtime_error"`     // 是否运行时错误
	StackOverflow    bool   `json:"stack_overflow"`    // 是否栈溢出
}

// Result 保存所有测试用例的执行结果
type Result struct {
	ExecDetails        []*ExecDetail `json:"exec_details"`         // 每个测试用例的执行结果
	CompileError       bool          `json:"compile_error"`        // 是否编译错误
	CompileErrorOutput string        `json:"compile_error_output"` // 错误输出
}

// RunConfig 运行配置
type RunConfig struct {
	Language           string
	Code               string
	InputList          []string
	MemoryLimitMiB     int64
	StackLimitMiB      int64
	TimeoutMillisecond int
}
