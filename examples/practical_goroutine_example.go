package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 实用示例: 并行获取多个区块信息
func demonstrateParallelBlockFetching() {
	fmt.Println("\n=== 实用示例: 并行获取多个区块信息 ===")

	var wg sync.WaitGroup
	results := make(chan BlockResult, 5) // 收集区块结果

	// 要获取的区块号
	blockNumbers := []int64{1000, 1001, 1002, 1003, 1004}

	// 为每个区块启动一个协程
	for _, blockNum := range blockNumbers {
		wg.Add(1)
		go func(blockNum int64) {
			defer wg.Done()

			// 模拟获取区块信息
			result := fetchBlockInfo(blockNum)
			results <- result
		}(blockNum)
	}

	// 启动一个协程等待所有工作完成
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集所有结果
	fmt.Println("收集区块信息...")
	var successCount, errorCount int

	for result := range results {
		if result.Error != nil {
			fmt.Printf("❌ 区块 %d 获取失败: %v\n", result.BlockNumber, result.Error)
			errorCount++
		} else {
			fmt.Printf("✅ 区块 %d: Hash=%s, Transactions=%d\n",
				result.BlockNumber, result.Hash[:10]+"...", result.TxCount)
			successCount++
		}
	}

	fmt.Printf("\n汇总: 成功 %d 个, 失败 %d 个\n", successCount, errorCount)
}

// BlockResult 区块结果结构
type BlockResult struct {
	BlockNumber int64
	Hash        string
	TxCount     int
	Error       error
}

// fetchBlockInfo 模拟获取区块信息
func fetchBlockInfo(blockNum int64) BlockResult {
	// 模拟网络延迟
	time.Sleep(time.Duration(100+blockNum%200) * time.Millisecond)

	// 模拟偶尔的失败
	if blockNum == 1002 {
		return BlockResult{
			BlockNumber: blockNum,
			Error:       fmt.Errorf("网络超时"),
		}
	}

	// 模拟成功获取数据
	return BlockResult{
		BlockNumber: blockNum,
		Hash:        fmt.Sprintf("0x%x1234567890abcdef", blockNum),
		TxCount:     int(blockNum % 50 + 1),
	}
}

// 实用示例: 带超时的协程控制
func demonstrateGoroutineWithTimeout() {
	fmt.Println("\n=== 实用示例: 带超时的协程控制 ===")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	results := make(chan TaskResult, 3)

	// 启动3个任务
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				// 超时处理
				results <- TaskResult{
					TaskID: taskID,
					Error:  ctx.Err(),
				}
				return

			default:
				// 执行任务
				result := performTask(taskID)
				results <- result
			}
		}(i)
	}

	// 等待所有任务完成
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集结果
	for result := range results {
		if result.Error != nil {
			fmt.Printf("任务 %d: %v\n", result.TaskID, result.Error)
		} else {
			fmt.Printf("任务 %d 成功: %s\n", result.TaskID, result.Message)
		}
	}
}

// TaskResult 任务结果
type TaskResult struct {
	TaskID  int
	Message string
	Error   error
}

// performTask 模拟执行任务
func performTask(taskID int) TaskResult {
	// 模拟不同的执行时间
	duration := time.Duration(taskID) * time.Second
	time.Sleep(duration)

	return TaskResult{
		TaskID:  taskID,
		Message: fmt.Sprintf("任务 %d 完成，耗时 %v", taskID, duration),
	}
}

// 实用示例: 协程池模式
func demonstrateWorkerPool() {
	fmt.Println("\n=== 实用示例: 协程池模式 ===")

	const numWorkers = 3
	const numTasks = 8

	jobs := make(chan int, numTasks)
	results := make(chan TaskResult, numTasks)

	// 创建worker池
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for taskID := range jobs {
				result := processTask(workerID, taskID)
				results <- result
			}
		}(i)
	}

	// 发送任务
	go func() {
		for i := 1; i <= numTasks; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// 等待所有worker完成
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集结果
	fmt.Printf("使用 %d 个worker处理 %d 个任务...\n", numWorkers, numTasks)

	for result := range results {
		fmt.Printf("Worker %d -> 任务 %d: %s\n",
			result.TaskID%numWorkers+1, result.TaskID, result.Message)
	}
}

// processTask worker处理任务
func processTask(workerID, taskID int) TaskResult {
	time.Sleep(500 * time.Millisecond) // 模拟处理时间
	return TaskResult{
		TaskID:  taskID,
		Message: fmt.Sprintf("处理完成"),
	}
}

func main() {
	demonstrateParallelBlockFetching()
	demonstrateGoroutineWithTimeout()
	demonstrateWorkerPool()
}