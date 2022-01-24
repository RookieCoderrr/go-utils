package sychronize
/*
channel 通常会用来两个线程的协作，一个线程必须在另个线程结束之后执行，或者线程之间的方法协调，这里主线程必须等待其他线程向通道发送信息才可以向下进行，
否则会阻塞在这里。这样就可以协同两个线程
 */
import "fmt"

func RunChannel() {
	c := make(chan int)
	go func() {
		fmt.Println("I am the first task")
		c <- 1
	}()
	<-c
	fmt.Println("I am the second task")
}
/*
注意channel是阻塞通道，这里创建一个大小为3的通道，第四个c <-1 命令会阻塞，因为通道已经占满，直到有接受者接收数据后，才可以有新的数据存入
 */
func RunCacheChannel(){
	c := make(chan int,3)
	c <-1
	c <-1
	c <-1
	c <-1
}
