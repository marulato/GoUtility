package main

import "document"

func main() {
	xlsx := document.NewExcel("D:\\test.xlsx", "Sheet1")
	for i := 0; i < 10; i++ {

		var value = []string  {"大福", "宝宝", "柴犬", "小黄人", "大福", "宝宝", "柴犬", "小黄人", "大福", "宝宝", "柴犬", "小黄人"}
		xlsx.SetRowValue(i,0, &value)
	}
	xlsx.Save()

}
