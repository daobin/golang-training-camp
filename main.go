package main

import (
	"trainingcamp/dao"
	"trainingcamp/pkg/setting"
	"trainingcamp/work"
)

func main() {
	setting.Setup()

	dao.Setup()
	defer dao.Close()

	// Go 训练营第 2 周作业
	//work.RunWeek02()

	// Go 训练营第 3 周作业
	//work.RunWeek03()

	// Go 训练营第 4 周作业
	work.RunWeek04()

	// Go 训练营第 xx 周作业

}
