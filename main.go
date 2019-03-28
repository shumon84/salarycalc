package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type 円 uint64

func (y 円) 兆() uint64 {
	return y.億() / uint64(10000)
}

func (y 円) 億() uint64 {
	return y.万() / uint64(10000)
}

func (y 円) 万() uint64 {
	return uint64(y) / uint64(10000)
}

func (y 円) String() string {
	口語表記 := ""
	兆 := y.兆()
	if 兆 != 0 {
		口語表記 += fmt.Sprintf("%d兆", 兆)
	}
	億 := y.億() % 10000
	if 億 != 0 {
		口語表記 += fmt.Sprintf("%d億", 億)
	}
	万 := y.万() % 10000
	if 万 != 0 {
		口語表記 += fmt.Sprintf("%d万", 万)
	}
	円 := y % 10000
	口語表記 += fmt.Sprintf("%d円", 円)
	return 口語表記
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage : $", os.Args[0], "[勤務時間ファイル]", "[時給]")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	hourSalary, err := strconv.Atoi(os.Args[2])
	if err == nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(file)
	sum := time.Duration(0)
	for sc.Scan() {
		line := sc.Text()
		times := strings.Split(line, " ")
		if len(times) != 3 {
			break
		}
		start, err := time.Parse("15:04", times[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := time.Parse("15:04", times[1])
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		working := end.Sub(start)
		sum += working
	}
	fmt.Println(sum, 円(sum.Hours()*float64(hourSalary)))
}
