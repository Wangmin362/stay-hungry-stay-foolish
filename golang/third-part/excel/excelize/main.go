package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	statis := make(map[string]int64)

	f, err := excelize.OpenFile("C:\\Users\\wangmin\\Desktop\\2019年以来就业创业培训数据统计表.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("花名册")
	if err != nil {
		fmt.Println(err)
		return
	}
	for idx, row := range rows {
		if idx == 0 || idx == 1 {
			continue
		}
		cnt, ok := statis[row[9]]
		if !ok {
			statis[row[9]] = 1
		} else {
			statis[row[9]] = cnt + 1
		}
	}

	sfz_cnt := 0
	for sfz, cnt := range statis {
		if cnt >= 6 {
			fmt.Printf("身份证：%s, 次数：%d\n", sfz, cnt)
			sfz_cnt++
		}
	}

	fmt.Println(sfz_cnt)
}
