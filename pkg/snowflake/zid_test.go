/*
   @Time : 2019-01-15 11:43
   @Author : frozenchen
   @File : zid_test
   @Software: 24on
*/
package snowflake

import (
	"fmt"
	"jtHome/commons/util"
	"sync"
	"testing"
)

var list map[int64]int64

var m sync.Mutex
var j int

func TestGen(t *testing.T) {

	println(0 % 1024)
	println(1 % 1024)
	println(2 % 1024)
	println(1024 % 1024)
	println(1025 % 1024)

	return
	list = map[int64]int64{}
	group := sync.WaitGroup{}
	//group.Add(10)
	//go g(group)
	fmt.Println(stepMax)

	//group.Wait()
	group.Add(10)
	for i := 0; i < 10; i++ {
		j++
		fmt.Println("j:", j)
		go g(&group, i)
	}

	group.Wait()

	var s int
	for k, _ := range list {
		s++
		fmt.Println(len(util.ToString(k)),k)
		if s >= 5 {

			return
		}
	}

}

func g(group *sync.WaitGroup, i int) {
	newNode, _ := NewNode(int64(i))
	for i := 0; i < 1000000; i++ {
		add(int64(newNode.Generate()))
	}

	fmt.Println(i, "done")
	group.Done()

}

func add(key int64) {
	//fmt.Println(1)
	m.Lock()
	list[key] = 1
	m.Unlock()
}
