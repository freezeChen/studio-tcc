/*
   @Time : 2019-07-12 14:51
   @Author : frozenChen
   @File : task
   @Software: studio-tcc
*/
package task

import "time"

type Task struct {
}

func New() *Task {
	return &Task{}
}

func Run() {
	go func() {
		for {
			timer := time.NewTimer(1 * time.Minute)
			<-timer.C

		}

	}()
}
