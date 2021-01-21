package ftc

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

// 关于 errgroup的第一个demo，将一个大的任务，分为三个小任务。
func errgroup1() {
	// 声明一个 errgroup
	errSlice := make([]error ,3)
	eg := new(errgroup.Group)
	// 任务1
	eg.Go(func() error {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("task1")
		errSlice[0] = nil
		return nil
	})
	// 任务2
	eg.Go(func() error {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("task2")
		errSlice[1] = fmt.Errorf("error test")
		return fmt.Errorf("error test")
	})
	// 任务3
	eg.Go(func() error {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("task3")
		errSlice[3] = nil
		return nil
	})
	if err := eg.Wait(); err != nil {
		fmt.Println(errSlice)
		fmt.Println(err)
	}
}
