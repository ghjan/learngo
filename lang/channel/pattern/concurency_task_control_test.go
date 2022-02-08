package pattern

import (
	"fmt"
	"testing"
	"time"
)

func TestNonBlockingWait(t *testing.T) {
	m1 := MsgGen("service1", 2, nil, 2)
	m2 := MsgGen("service2", 2, nil, 2)
	i := 0
	countSuccess1 := 0
	countSuccess2 := 0
	for i < 20 {
		i++
		msg1 := <-m1
		countSuccess1++
		fmt.Println(msg1)
		if m, ok := NonBlockingWait(m2); ok {
			countSuccess2++
			fmt.Println(m)
		} else {
			fmt.Println("mo message from service2")
		}
	}
	fmt.Printf("countSuccess1:%d,countSuccess2:%d\n", countSuccess1, countSuccess2)
}

func TestTimeoutWait(t *testing.T) {
	done := make(chan struct{})
	m1 := MsgGen("service1", 5, done, 2)
	i := 0
	countSuccess1 := 0
	countTimeout := 0
	for i < 5 {
		i++
		msg1 := <-m1
		countSuccess1++
		fmt.Println(msg1)
		if m, ok := TimeoutWait(m1, time.Second); ok {
			countSuccess1++
			fmt.Println(m)
		} else {
			countTimeout++
			fmt.Println("timeout")
		}

	}
	done <- struct{}{}
	<-done
	time.Sleep(time.Second)
	fmt.Printf("countSuccess1:%d,countTimeout:%d\n", countSuccess1, countTimeout)

}
