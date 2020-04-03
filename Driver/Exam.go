package main

import (
	"fmt"
	"time"
)


var (
	chNames = make(chan string,100)
	chLanes = make(chan int,5)  //信号量，5人考试
	chFouls = make(chan string,100) //违纪通道

	scoreMap = make(map[string]int) //分数

	examers = make([]string,0)  //考生

)


func Patrol() {
	ticker := time.NewTimer(1*time.Second)
	for  {
		//fmt.Println("考官正在巡考。。")
		select {
		case name := <- chFouls:
			fmt.Println(name,"考试违纪")
		default:
			fmt.Println("考场秩序良好")
		}
		<- ticker.C
	}
}

func TakeExam(name string) {
	chLanes <- 123
	fmt.Println(name,"正在考试...")
	examers = append(examers,name)

	//生成考试成绩
	score := GetRandomInt(0,100)
	scoreMap[name] = score
	if score < 10{
		score = 0
		chFouls <- name
	}

	<- time.After(3*time.Second) //考试持续5秒
	<- chLanes
	wg.Done()
}

func QueryScore(name string) {
	score,err := QueryFromRedis(name)

	if err != nil {
		score,_ = QueryScoreFromMysql(name)
		fmt.Println(name,"+",score)

		//将数据写入Redis
		WriteScoreToRedis(name,score)
	}else {
		fmt.Println(name,":",score)
	}
	wg.Done()
}





