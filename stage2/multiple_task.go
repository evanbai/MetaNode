package stage2

import (
	"fmt"
	"sync"
	"time"
)

// 1. 任务结构体：封装任务标识和执行逻辑
type Task struct {
	ID   string       // 任务唯一标识（用于区分不同任务）
	Func func() error // 任务执行函数：无参数，返回错误（表示执行结果）
}

// 2. 任务执行结果结构体：记录执行统计信息
type TaskResult struct {
	TaskID   string        // 关联的任务ID
	Duration time.Duration // 执行耗时
	Err      error         // 执行错误（nil表示成功）
}

// 3. 任务调度器：接收任务列表，并发执行并返回所有结果
func ScheduleTasks(tasks []Task) []TaskResult {
	var (
		wg       sync.WaitGroup                      // 等待所有协程完成
		resultCh = make(chan TaskResult, len(tasks)) // 缓冲通道：存储任务结果
	)

	// 遍历任务，启动协程执行
	for _, task := range tasks {
		// 注意：循环变量复用问题，需通过参数传递task（避免所有协程共享同一个循环变量）
		wg.Add(1)
		go func(t Task) {
			defer wg.Done() // 任务完成后，通知WaitGroup

			// 记录任务开始时间
			startTime := time.Now()

			// 执行任务
			err := t.Func()

			// 计算执行耗时，构造结果
			duration := time.Since(startTime)
			resultCh <- TaskResult{
				TaskID:   t.ID,
				Duration: duration,
				Err:      err,
			}
		}(task) // 传入当前循环的task副本
	}

	// 启动独立协程：等待所有任务完成后关闭通道（避免主程序接收时阻塞）
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// 收集所有任务结果
	var results []TaskResult
	for res := range resultCh {
		results = append(results, res)
	}

	return results
}

// ------------------------------
// 测试代码：模拟不同类型的任务
// ------------------------------
func callFunc() {
	// 构造测试任务列表（包含成功、失败、不同耗时的任务）
	tasks := []Task{
		{
			ID: "task-1",
			Func: func() error {
				time.Sleep(150 * time.Millisecond) // 模拟耗时150ms的任务
				return nil                         // 执行成功
			},
		},
		{
			ID: "task-2",
			Func: func() error {
				time.Sleep(200 * time.Millisecond) // 模拟耗时200ms的任务
				return fmt.Errorf("连接数据库失败")       // 执行失败
			},
		},
		{
			ID: "task-3",
			Func: func() error {
				time.Sleep(100 * time.Millisecond) // 模拟耗时100ms的任务
				return nil                         // 执行成功
			},
		},
		{
			ID: "task-4",
			Func: func() error {
				time.Sleep(50 * time.Millisecond) // 模拟耗时50ms的任务
				return nil                        // 执行成功
			},
		},
	}

	// 启动调度器执行任务
	fmt.Println("任务调度器启动，开始执行任务...")
	startTime := time.Now()
	results := ScheduleTasks(tasks)
	totalDuration := time.Since(startTime)

	// 打印执行结果统计
	fmt.Printf("\n所有任务执行完成，总耗时：%v\n", totalDuration)
	fmt.Println("------------------------------")
	for _, res := range results {
		status := "成功"
		if res.Err != nil {
			status = "失败"
		}
		fmt.Printf("任务ID：%s | 执行状态：%s | 耗时：%v | 错误：%v\n",
			res.TaskID, status, res.Duration, res.Err)
	}
}
