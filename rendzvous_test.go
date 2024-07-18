package boomer

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewRendezvousPoint(t *testing.T) {
	rendezvousPointA := NewRendezvousPoint("rendezvousPointA", 10)
	if rendezvousPointA == nil {
		t.Errorf("TestNewRendezvousPoint Error.")
	}
	rendezvousPointB := NewRendezvousPoint("rendezvousPointA", 20)
	if rendezvousPointB == nil {
		t.Errorf("TestNewRendezvousPoint Error.")
	}
}

func TestRendezvousPoint(t *testing.T) {
	rendezvousPointA := NewRendezvousPoint("rendezvousPointC", 1000)
	if rendezvousPointA == nil {
		t.Errorf("TestNewRendezvousPoint Error.")
	}
	for i := 0; i < 1000; i++ {
		go func(i int) {
			rand.Seed(time.Now().UnixNano()) // 使用当前时间的纳秒数作为随机种子

			minSleep := 1000  // 最小睡眠时间，单位为毫秒
			maxSleep := 20000 // 最大睡眠时间，单位为毫秒

			// 生成介于min和max之间的随机睡眠时间
			randomSleep := rand.Intn(maxSleep-minSleep+1) + minSleep
			fmt.Printf("Sleeping for %d milliseconds...\n", randomSleep)

			time.Sleep(time.Duration(randomSleep) * time.Millisecond)

			rendezvousPointA.Done(-1)
			rendezvousPointA.Wait()
			fmt.Println("rendezvous:", i)
		}(i)
	}
	time.Sleep(100 * time.Second)
}

//func TestWaitGroup(t *testing.T) {
//	var wg sync.WaitGroup
//	wg.Add(1000)
//	for i := 0; i < 1000; i++ {
//		go func(i int) {
//			rand.Seed(time.Now().UnixNano()) // 使用当前时间的纳秒数作为随机种子
//
//			min := 1000  // 最小睡眠时间，单位为毫秒
//			max := 10000 // 最大睡眠时间，单位为毫秒
//
//			// 生成介于min和max之间的随机睡眠时间
//			randomSleep := rand.Intn(max-min+1) + min
//			fmt.Printf("Sleeping for %d milliseconds...\n", randomSleep)
//
//			time.Sleep(time.Duration(randomSleep) * time.Millisecond)
//			wg.Done()
//
//		}(i)
//
//	}
//	wg.Wait()
//	fmt.Println("rendezvous")
//	time.Sleep(100 * time.Second)
//}
