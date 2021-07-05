package week

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Go 训练营第 4 周作业
// —————— 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，
// —————— 以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。
func RunWeek04() {
	// 重构了项目结构，主要将训练营作业当成一个应用，其代码组织在 cmd/work 中，同时将作业代码放在应用内部的 internal 目录中

	// 使用 Wire 构建依赖


	// main 函数对于服务的注册和启动，信号处理
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
}
