package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

func main() {
	fmt.Println("@ 統一發票兌獎 @")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"功能", "快捷鍵"})
	data := [][]string{
		[]string{"q", "離開"},
		[]string{"now", "本期"},
		[]string{"up", "上一期"},
	}
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	Init()
}
func Init()  {
	var s string
	for {
		fmt.Print(">> 顯示發票資訊：")
		fmt.Scanf("%s", &s)
		fmt.Println("")
		if s == "up" || s == "now" {
			getTicket(s)
		} else if s == "q" {
			color.Green("（πーπ）依依不捨地離開了。")
			break
		} else {
			break;
		}
	}

}
func getTicket(str string) {
	var choiceMonth string
	if (str == "now") {
		choiceMonth = "#area1"
	} else if (str == "up") {
		choiceMonth = "#area2"
	}
	doc, err := goquery.NewDocument("http://invoice.etax.nat.gov.tw/")
	if err != nil {
		panic(err)
	}
	var title = doc.Find(choiceMonth).Find("h2").Eq(1).Text()
	//var itemList = list.New()
	itemList := make([]string, 0)
	doc.Find(".t18Red").Each(func(i int, s *goquery.Selection) {
		itemList = append(itemList, s.Text())
	})

	fmt.Println("====================")
	fmt.Println(title)
	fmt.Println("====================")
	var i = 0
	tag := []string{"特別獎", "特獎", "頭獎", "特別號"}
	var pointTag string
	for index, val := range itemList {
		if str == "now" {
			if index > 3 {
				continue
			}
		} else if str == "up" {
			if index < 4 {
				continue
			}
		} else {
			break;
		}
		fmt.Printf("=>%s:%s\r\n", tag[i], val)
		if i == 2 {
			pointTag = val
		}
		i++
	}
	pointItem := strings.Split(pointTag, "、")
	tag = []string{"二獎", "三獎", "四獎", "五獎", "六獎"}
	fmt.Println("=============================")
	for _, val := range pointItem {
		i = 1
		if len(val) < 6 {
			continue
		}
		for i := 1; i <= 5; i++ {
			fmt.Printf("%s : %s\r\n", tag[i - 1], val[i:])
		}
		fmt.Println("=============================")
	}
}