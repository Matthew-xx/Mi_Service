package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"sync"
)

var dbMutex sync.RWMutex  //数据库读写锁

func WriteScoreToMysql(scoreMap map[string]int)  {
	dbMutex.Lock()  //锁定为写锁，写入期间不允许读访问

	db, err := sqlx.Open("mysql", "root:666666@tcp(127.0.0.1:3306)/driver?charset=utf8")
	HandlerError(err)
	defer db.Close()

	for name,score := range scoreMap{
		_,err := db.Exec("insert into score(name,score) values(?,?);",name,score)
		HandlerError(err)
		//fmt.Println("插入成功")
	}
	fmt.Println("成绩录入完毕！")

	dbMutex.Unlock()  //解锁数据库，开放查询
	//wg.Done()  //等待组是主函数中的调度，放这里则耦合了。解决：在主函数里写匿名函数
}

func WriteScoreToRedis(name string,score int) error {
	conn,err := redis.Dial("tcp","192.168.99.100:6379")
	HandlerError(err)
	defer conn.Close()

	_,err = conn.Do("set",name,score)
	fmt.Println("redis写入成功")
	return err
}

func QueryScoreFromMysql(name string) (score int,err error) {
	//fmt.Println("QueryFromMysql")
	dbMutex.RLock()  //读锁
	db, err := sqlx.Connect("mysql", "root:666666@tcp(127.0.0.1:3306)/driver?charset=utf8")
	HandlerError(err)
	defer db.Close()

	examscore := make([]ExamScore,0) //创建切片接收成绩
	err = db.Select(&examscore,"select * from score where name=?;",name)
	if err != nil {
		return
	}
	//fmt.Println(examscore)

	dbMutex.RUnlock()
	//	wg.Done()
	return examscore[0].Score,nil
}

func QueryFromRedis(name string) (score int,err error){
	//fmt.Println("QueryFromRedis")
	conn,err := redis.Dial("tcp","192.168.99.100:6379")
	HandlerError(err)
	defer conn.Close()

	reply,err1 := conn.Do("get",name)  //Redis里面没有int得到的是string
	if reply != nil {
		score,err1 = redis.Int(reply,err1)
	}else {
		return 0,errors.New("未能从Redis中查到数据")
	}

	if err1 != nil {
		return 0,err1
	}

	return score,nil
}

