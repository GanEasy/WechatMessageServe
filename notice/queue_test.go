// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notice

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_T(t *testing.T) {

	//初始化对象池

	maxWorkers := 10
	maxQueue := 200
	//初始化一个调试者,并指定它可以操作的 工人个数
	dispatch := NewDispatcher(maxWorkers)
	JobQueue = make(chan Job, maxQueue) //指定任务的队列长度
	//并让它一直接运行
	dispatch.Run()

	for i := 0; i < 100; i++ {
		p := NoticeDemo{
			fmt.Sprintf("[%s]", strconv.Itoa(i)),
			fmt.Sprintf("玩家-[%s]", strconv.Itoa(i)),
		}
		JobQueue <- Job{
			Notice: &p,
		}
		fmt.Println(i, len(JobQueue))
		// time.Sleep(time.Millisecond)
	}
	time.Sleep(time.Second * 11)
	close(JobQueue)
}
