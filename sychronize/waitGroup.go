package sychronize
/*WaitGroup 主要用法是，如果有一需求，需要有多个服务拼凑起来返回，那么必须要等所有服务都返回后才把结果返回。
但是如果有一个routine报错，其他routine不会停止,这样会浪费资源，但是waitgroup做不到这点，需要waitgroup与channel共同使用
WaitGroup类似队列，可以向队列中加任务，任务结束后从队列去除，如果有任务没完成会阻塞队列
*/
import (
	"fmt"
	"sync"
	"time"
)

func RunWaitGroup(){
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		//这个地方不是值传递，是地址传递
		go work(i,&wg)
	}
	wg.Wait()
	fmt.Println("All routines done")
}
//不是值传递，是地址传递
func work(i int, wg *sync.WaitGroup) {
	fmt.Printf("Routine %d is working ", i)
	time.Sleep(time.Second)
	wg.Done()
}