package week

import (
	"encoding/json"
	"fmt"
	"trainingcamp/biz"
)

// Go 训练营第 2 周作业
// —————— 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error 抛给上层。
// —————— 为什么，应该怎么做请写出代码？
// 答：对于一些底层的数据库操作遇到 sql.ErrNoRows 时可以使用 Wrap 抛给上层处理，让上层调用者知晓较为具体的错误原因，并让其根据需要进行处理，
// 但对于一些已经做好封装处理的数据库操作时则不建议，直接返回结果就好，比如使用 gorm 查询数据为空时
func RunWeek02() {
	// 底层数据库操作遇到 sql.ErrNoRows 使用 Wrap 向上抛的代码如下：
	article, err := biz.GetArticleById(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(res))
}