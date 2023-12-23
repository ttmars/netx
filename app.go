package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) TestNetwork(urls []string, proxy string, timeout int) [][2]string {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	if proxy != "" {
		client, _ = createHttpProxyClient(proxy, timeout)
	}

	result := make(chan [2]string, len(urls))
	var wg sync.WaitGroup
	for i, u := range urls {
		wg.Add(1)
		go func(iii int, uuu string) {
			defer wg.Done()
			start := time.Now()
			resp, err := client.Get(uuu)
			if err != nil {
				result <- [2]string{fmt.Sprintf("%v", iii), "超时"}
				return
			}
			defer resp.Body.Close()
			_, err = io.ReadAll(resp.Body)
			if err != nil {
				result <- [2]string{fmt.Sprintf("%v", iii), "超时"}
				return
			}
			result <- [2]string{fmt.Sprintf("%v", iii), time.Since(start).String()}
		}(i, u)
	}
	wg.Wait()
	close(result)
	var res [][2]string
	for v := range result {
		res = append(res, v)
	}
	return res
}

// CreateHttpProxyClient 创建一个http代理客户端，代理类型由uri确定，支持http/https/socks5，默认为http
// http://fans007:fans888@45.76.169.156:39000
func createHttpProxyClient(proxy string, timeout int) (*http.Client, error) {
	u, err := url.Parse(proxy)
	if err != nil {
		return nil, err
	}
	return &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(u)},
		Timeout:   time.Duration(timeout) * time.Second,
	}, nil
}
