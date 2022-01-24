package sychronize
/*
通过协程池限制goroutine数是一个有效避免泄漏的手段, 使用waitgroup和ants可以有效的控制线程数量
 */
import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
)

func RunAnts(){
	var wg sync.WaitGroup
	//这里第一个参数是线程池大小，一次最多有十个线程同时工作，第二个参数是一个函数，他的函数参数是由p.Invoke()来传入的
	p,_ := ants.NewPoolWithFunc(10, func(i interface{}) {
		subTask(i)
		wg.Done()
	})
	//如果这里使用subTask 或者 go subTask ，要么只有单线程，要么有1000个线程，都不是一个好选择，所以选择用线程池。
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		_ = p.Invoke(i)
	}
	wg.Wait()
	fmt.Println("All tasks done !")
}

func subTask(i interface{}) {
	fmt.Println(fmt.Printf("task %d is working !",i))
}