//go:build unix

package main

import (
	"fmt"
	"log"
	"net"
	"syscall"
)

func main() {
	// 创建监听 socket
	listenFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		log.Fatal(err)
	}

	// 绑定地址和端口
	addr := syscall.SockaddrInet4{Port: 8088}
	copy(addr.Addr[:], net.ParseIP("0.0.0.0").To4())
	err = syscall.Bind(listenFd, &addr)
	if err != nil {
		log.Fatal(err)
	}

	// 开始监听
	err = syscall.Listen(listenFd, 10)
	if err != nil {
		log.Fatal(err)
	}

	// 创建 epoll 实例
	epfd, err := syscall.EpollCreate1(0)
	if err != nil {
		log.Fatal(err)
	}
	// 创建 EPOLLIN 信号事件
	event := syscall.EpollEvent{Events: syscall.EPOLLIN, Fd: int32(listenFd)}
	// 将监听 EPOLLIN 信号的事件 添加到 epoll 实例中
	err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, listenFd, &event)
	if err != nil {
		log.Fatal(err)
	}

	// 处理 epoll 事件
	events := make([]syscall.EpollEvent, 100)
	// for循环
	for {
		// 等待可用事件通知
		n, err := syscall.EpollWait(epfd, events, -1)
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < n; i++ {
			if events[i].Fd == int32(listenFd) {
				// 接受新的连接
				connFd, _, err := syscall.Accept(listenFd)
				if err != nil {
					log.Fatal(err)
				}

				// 将新的连接添加到 epoll 实例中
				event := syscall.EpollEvent{Events: syscall.EPOLLIN, Fd: int32(connFd)}
				err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, connFd, &event)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				// 处理已连接的客户端数据
				buf := make([]byte, 1024)
				n, err := syscall.Read(int(events[i].Fd), buf)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Received: %s\n", string(buf[:n]))
			}
		}
	}
}
