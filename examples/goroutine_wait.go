package main

import (
	"fmt"
	"sync"
	"time"
)

// 方法1: 使用sync.WaitGroup (最常用)
func demonstrateWaitGroup() {
	fmt.Println("\n=== 方法1: 使用sync.WaitGroup ===")

	var wg sync.WaitGroup

	// 启动3个协程
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 每启动一个协程，计数器+1
		go func(id int) {
			defer wg.Done() // 协程结束时，计数器-1

			fmt.Printf("协程 %d 开始执行\n", id)
			time.Sleep(time.Duration(id) * time.Second) // 模拟耗时操作
			fmt.Printf("协程 %d 执行完成\n", id)
		}(i)
	}

	fmt.Println("主协程等待所有子协程完成...")
	wg.Wait() // 阻塞直到计数器为0
	fmt.Println("所有协程执行完成！")
}

// 方法2: 使用channel (手动实现等待机制)
func demonstrateChannel() {
	fmt.Println("\n=== 方法2: 使用channel ===")

	done := make(chan bool) // 创建一个channel用于通知完成
	workerCount := 3

	// 启动3个协程
	for i := 1; i <= workerCount; i++ {
		go func(id int) {
			fmt.Printf("协程 %d 开始执行\n", id)
			time.Sleep(time.Duration(id) * time.Second)
			fmt.Printf("协程 %d 执行完成\n", id)
			done <- true // 发送完成信号
		}(i)
	}

	fmt.Println("主协程等待所有子协程完成...")

	// 等待所有协程发送完成信号
	for i := 0; i < workerCount; i++ {
		<-done
	}

	fmt.Println("所有协程执行完成！")
}

// 方法3: 使用channel收集结果
func demonstrateResultCollection() {
	fmt.Println("\n=== 方法3: 使用channel收集结果 ===")

	type Result struct {
		id      int
		message string
	}

	results := make(chan Result, 3) // 带缓冲的channel收集结果

	// 启动3个协程
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("协程 %d 开始执行\n", id)
			time.Sleep(time.Duration(id) * time.Second)

			result := Result{
				id:      id,
				message: fmt.Sprintf("协程 %d 的结果", id),
			}

			results <- result // 发送结果
		}(i)
	}

	fmt.Println("主协程收集所有结果...")

	// 收集所有结果
	allResults := make([]Result, 0, 3)
	for i := 0; i < 3; i++ {
		result := <-results
		fmt.Printf("收到结果: %+v\n", result)
		allResults = append(allResults, result)
	}

	fmt.Printf("所有协程完成，共收集到 %d 个结果\n", len(allResults))
}

// 方法4: 错误处理的WaitGroup使用
func demonstrateWaitGroupWithError() {
	fmt.Println("\n=== 方法4: 带错误处理的WaitGroup ===")

	var wg sync.WaitGroup
	errors := make(chan error, 3) // 用于收集错误的channel

	// 启动3个协程，其中一个会返回错误
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Printf("协程 %d 开始执行\n", id)
			time.Sleep(time.Duration(id) * time.Second)

			// 模拟协程2出错
			if id == 2 {
				errors <- fmt.Errorf("协程 %d 执行失败", id)
				return
			}

			fmt.Printf("协程 %d 执行完成\n", id)
		}(i)
	}

	// 启动一个额外的协程来等待WaitGroup
	go func() {
		wg.Wait()
		close(errors) // 所有协程完成后关闭errors channel
	}()

	fmt.Println("主协程等待并检查错误...")

	// 检查是否有错误
	var hasError bool
	for err := range errors {
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			hasError = true
		}
	}

	if !hasError {
		fmt.Println("所有协程成功完成！")
	} else {
		fmt.Println("部分协程执行失败！")
	}
}

func main() {
	demonstrateWaitGroup()
	demonstrateChannel()
	demonstrateResultCollection()
	demonstrateWaitGroupWithError()
}