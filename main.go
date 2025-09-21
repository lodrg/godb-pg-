package main

import (
	"fmt"
	"godb-pg-/logger"
	"godb-pg-/tree"
)

// @Title        main.go
// @Description
// @Create       david 2025-09-21 16:29
// @Update       david 2025-09-21 16:29

func main() {
	logger.SetLevel(logger.DEBUG)
	logger.Info("Hello World")
	tree := tree.NewBPTree(4)

	// 插入测试数据
	testData := map[uint32]string{
		1:  "一",   // One
		2:  "二",   // Two
		3:  "三",   // Three
		4:  "四",   // Four
		5:  "五",   // Five
		6:  "六",   // Six
		7:  "七",   // Seven
		8:  "八",   // Eight
		9:  "九",   // Nine
		10: "十",   // Ten
		11: "十一", // Eleven
		12: "十二", // Twelve
		13: "十三", // Thirteen
		14: "十四", // Fourteen
		15: "十五", // Fifteen
		16: "十六", // Sixteen
		17: "十七", // Seventeen
		18: "十八", // Eighteen
		19: "十九", // Nineteen
		20: "二十", // Twenty
	}

	// 插入数据并打印树的状态
	for k, v := range testData {
		tree.Insert(int(k), []byte(v))
		fmt.Printf("\n插入 %d:%s 后的树结构:\n", k, v)
		tree.Print()
	}

	// 搜索测试
	fmt.Println("\n搜索测试:")
	for k := 1; k <= 10; k++ {
		if v, found := tree.Search(k); found {
			fmt.Printf("找到键 %d，值为: %s\n", k, v)
		}
	}
}
