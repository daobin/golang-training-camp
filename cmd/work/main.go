package main

import (
	"trainingcamp/cmd/work/internal/week"
	"trainingcamp/dao"
	"trainingcamp/pkg/setting"
)

func main() {
	setting.Setup()

	dao.Setup()
	defer dao.Close()

	// Go 训练营第 2 周作业
	//week.RunWeek02()

	// Go 训练营第 3 周作业
	//week.RunWeek03()

	// Go 训练营第 4 周作业
	//week.RunWeek04()

	// Go 训练营第 8 周作业
	week.RunWeek08()

	// Go 训练营第 9 周作业
	week.RunWeek09()

	// Go 训练营第 13 周作业（毕业项目）
	week.RunWeek13()
}
