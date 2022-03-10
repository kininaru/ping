package main

import (
	"fmt"
	"github.com/go-ping/ping"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("参数不足。请至少指定一个目的地址。")
	}
	var err error
	count := 4
	index := 1
	for i, arg := range os.Args {
		if arg[0] == '-' {
			switch arg[1] {
			case 'c':
				count, err = strconv.Atoi(os.Args[i+1])
			}
			index = i + 2
		}
		if err != nil {
			panic(err)
		}
	}

	for _, dest := range os.Args[index:] {
		PingAndPrint(dest, count)
		fmt.Println()
	}
}

func PingAndPrint(host string, count int) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		panic(err)
	}

	pinger.SetPrivileged(true)
	pinger.Count = count
	err = pinger.Run()
	if err != nil {
		panic(err)
	}

	stat := pinger.Statistics()
	fmt.Printf("向地址 %s（IP 为 %s) 发送了 %d 个请求，结果如下：\n", stat.Addr, stat.IPAddr, count)
	fmt.Println(stat.Rtts)
	fmt.Printf("平均 RTT 为 %.3fms。其中最大 RTT 为 %.3fms，最小 RTT 为：%.3fms。\n", float64(stat.AvgRtt)/1000000, float64(stat.MaxRtt)/1000000, float64(stat.MinRtt)/1000000)
}
