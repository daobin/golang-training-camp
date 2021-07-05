package week

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
)

// Go 训练营第 3 周作业
// —————— 基于 errgroup 实现一个 http server 的启动和关闭，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func RunWeek03() {
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