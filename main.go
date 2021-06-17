package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"trainingcamp/dao"
	"trainingcamp/pkg/setting"
)

func main() {
	setting.Setup()

	dao.Setup()
	defer dao.Close()

	// Go 训练营第 2 周作业
	// —————— 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error 抛给上层。
	// —————— 为什么，应该怎么做请写出代码？

	// 答：对于一些底层的数据库操作遇到 sql.ErrNoRows 时可以使用 Wrap 抛给上层处理，让上层调用者知晓较为具体的错误原因，并让其根据需要进行处理，
	// 但对于一些已经做好封装处理的数据库操作时则不建议，直接返回结果就好，比如使用 gorm 查询数据为空时

	// 底层数据库操作遇到 sql.ErrNoRows 使用 Wrap 向上抛的代码如下：
	//article, err := biz.GetArticleById(1)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//res, err := json.Marshal(article)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(string(res))

	// Go 训练营第 3 周作业
	// —————— 基于 errgroup 实现一个 http server 的启动和关闭，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

	eg, ctx := errgroup.WithContext(context.Background())

	webHandler := http.NewServeMux()
	webHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello HaHa..."))
	})

	// 手动模拟服务退出
	serOutChan := make(chan struct{})
	webHandler.HandleFunc("/ser_out", func(w http.ResponseWriter, r *http.Request) {
		serOutChan <- struct{}{}
	})

	ser := &http.Server{
		Addr:           ":8090",
		Handler:        webHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 监听 Http Server
	eg.Go(func() error {
		return ser.ListenAndServe()
	})

	// 监听手动模拟服务退出
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			log.Printf("CTX Done ...%+v", ctx.Err())
		case <-serOutChan:
			log.Println("Http Server out ...")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
		defer cancel()
		log.Println("shutdown server ...")
		return ser.Shutdown(timeoutCtx)
	})

	// Linux Signal 信号的注册与处理
	eg.Go(func() error {
		signalChan := make(chan os.Signal)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			log.Printf("CTX Done2 ...%+v", ctx.Err())
			return ctx.Err()
		case sig := <-signalChan:
			log.Printf("Get Signal ...%+v", sig)
			return errors.Errorf("Get Signal: %v", sig)
		}
	})

	// errgroup 等待阻塞
	fmt.Printf("ErrGroup Wait Completed: %+v\r\n", eg.Wait())
	fmt.Println("Main Over")

	// Go 训练营第 4 周作业
	// —————— 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，
	// —————— 以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

}
