package main

import (
	//"database/sql"
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup //等待组，直接声明类型，不用赋值
	)

func main()  {
	for i:=0;i<20;i++ {
		chNames <- GetRandomName()
	}
	close(chNames)

	//巡考
	go Patrol()

	//考生并发考试
	for name := range chNames{
		wg.Add(1)  //添加
		go TakeExam(name)
	}
	wg.Wait() //等待所有主协程
	fmt.Println("考试完毕！")

	//录入成绩
	wg.Add(1)
	go func() {
		WriteScoreToMysql(scoreMap)
		wg.Done()
	}()
	<- time.After(1*time.Second)  //给一个时间间隔，确保WriteScore先抢到数据库读写锁

	//考生查询成绩
	for _,name := range examers{
		wg.Add(1)
		go QueryScore(name)
	}
	<- time.After(1*time.Second)
	for _,name := range examers{
		wg.Add(1)
		go QueryScore(name)
	}

	wg.Wait()
	fmt.Println("成绩录入完毕")
}
//使用二级缓存查成绩
