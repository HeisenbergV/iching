package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

var gua8 = []byte{0b111, 0b110, 0b101, 0b100, 0b011, 0b010, 0b001, 0b000}

func Iching(num1, num2, num3 int) (gua64 byte, yao int) {
	yao = num3 % 6

	num1 = num1 % 8
	if num1 == 0 {
		num1 = 7
	} else {
		num1 = num1 - 1
	}

	num2 = num2 % 8
	if num2 == 0 {
		num2 = 7
	} else {
		num2 = num2 - 1
	}

	xiagua := gua8[num1]
	shanggua := gua8[num2]

	gua64 = xiagua<<3 | shanggua
	return
}

type Gua struct {
	Name   string   `json:"gua-name"`
	Xiang  string   `json:"gua-xiang"`
	Detail string   `json:"gua-detail"`
	Yao    []string `json:"yao-detail"`
}

func Str2DEC(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) - 48) << uint8(i)
	}
	return
}

var yinyaoName = []string{"初六", "六二", "六三", "六四", "六五", "上六"}
var yangyaoName = []string{"初九", "九二", "九三", "九四", "九五", "上九"}

func RenderGua(g Gua, yaoNum int) {
	yao := g.Xiang[:]
	fmt.Printf("%s卦:%s\n", g.Name, g.Detail)
	for i := 5; i >= 0; i-- {
		if string(yao[i]) == "0" {
			if yaoNum == i {
				fmt.Printf("====  ====   <---- %s: %s\n", yinyaoName[i], g.Yao[i])
			} else {
				fmt.Println("====  ====")
			}
		} else {
			if yaoNum == i {
				fmt.Printf("==========   <---- %s: %s\n", yangyaoName[i], g.Yao[i])
			} else {
				fmt.Println("==========")
			}
		}
	}
}

func main() {
	fmt.Println("易经占卦")
	fmt.Println("1. 不义不占：你如果问 明天抢银行能成功吗，这种不义的事情无法占卜")
	fmt.Println("2. 不疑不占：有真正的疑惑，理性可以解决的就不要问了，理性实在是不能解决的，左右为难的，你再来问")
	fmt.Println("3. 不诚不占：不真诚不占")
	fmt.Println("输入三组数字，每组3位数 例：（123 456 789)")

	var num1, num2, num3 int
	fmt.Scan(&num1, &num2, &num3)
	num1len := len(strconv.Itoa(num1))
	num2len := len(strconv.Itoa(num2))
	num3len := len(strconv.Itoa(num3))

	if num1len != 3 || num2len != 3 || num3len != 3 {
		fmt.Println("三组数字需都是3位数")
		return
	}
	fmt.Println()
	fmt.Println()
	file, err := os.ReadFile("gua.json")
	if err != nil {
		log.Fatalf("Some error occured while reading file. Error: %s", err)
	}

	var gualist []Gua
	err = json.Unmarshal(file, &gualist)
	if err != nil {
		return
	}
	gua64, yao := Iching(num1, num2, num3)
	for _, gua := range gualist {
		if int(gua64) != Str2DEC(gua.Xiang) {
			continue
		}
		RenderGua(gua, yao)
		return
	}
}
