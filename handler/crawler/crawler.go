package crawler

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
)

type Role struct {
	Star            string // 星數
	Camp            string // 陣營
	Name            string // 武將
	Cost            string // 佔位
	Cavalry         string // 騎兵
	Shield          string // 盾兵
	Bow             string // 弓兵
	Spear           string // 槍兵
	instrument      string // 器械
	Command         string // 統率
	Intelligence    string // 智力
	Speed           string // 速度
	Politics        string // 政治
	Charm           string // 魅力
	Tactic          string // 戰法
	inheritedTactic string // 傳承戰法
}

func GetRoleInfo() {
	c := colly.NewCollector()

	file, err := os.Create("roleInfo.txt")

	if err != nil {
		fmt.Print(err)
		return
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("error =>", e)
	})

	c.OnResponse(func(r *colly.Response) {
		doc, err := htmlquery.Parse(strings.NewReader(string(r.Body)))

		if err != nil {
			log.Fatal(err)
		}

		titles := htmlquery.Find(doc, `//*[@class="table-responsive"]/table/thead/tr/th`)

		// 寫入標題
		for _, title := range titles {
			textNode := htmlquery.FindOne(title, `.`)
			text := strings.Split(htmlquery.InnerText(textNode), " ")
			_, _ = file.Write([]byte(text[0] + " "))
		}

		_, _ = file.Write([]byte("\n"))

		roles := htmlquery.Find(doc, `//*[@class="table-responsive"]/table/tbody/tr`)
		// roleList = []Role{}

		// 寫入武將內容
		for _, role := range roles {
			contents := htmlquery.Find(role, `./td`)
			// singleRole := []string{}
			value := Role{}
			v := reflect.ValueOf(value)

			for i, content := range contents {
				contentNode := htmlquery.FindOne(content, `.`)
				// singleRole = append(singleRole, htmlquery.InnerText(contentNode))
				v.FieldByIndex()
				_, _ = file.Write([]byte(htmlquery.InnerText(contentNode) + " "))
			}

			// value := Role{}
			// v := reflect.ValueOf(value)

			// for i := 0; i < v.NumField(); i++ {
			// 	v.Elem().FieldByName("Star").SetString(singleRole[i])
			// 	Star = reflect.Indirect(v).FieldByName("Star")
			// 	Camp := reflect.Indirect(v).FieldByName("Camp")
			// 	Name := reflect.Indirect(v).FieldByName("Name")
			// 	Cost := reflect.Indirect(v).FieldByName("Cost")
			// 	Cavalry := reflect.Indirect(v).FieldByName("Cavalry")
			// 	Shield := reflect.Indirect(v).FieldByName("Shield")
			// 	Bow := reflect.Indirect(v).FieldByName("Bow")
			// 	Spear := reflect.Indirect(v).FieldByName("Spear")
			// 	Command := reflect.Indirect(v).FieldByName("Command")
			// 	Intelligence := reflect.Indirect(v).FieldByName("Intelligence")
			// 	instrument := reflect.Indirect(v).FieldByName("instrument")
			// 	Speed := reflect.Indirect(v).FieldByName("Speed")
			// 	Politics := reflect.Indirect(v).FieldByName("Politics")
			// 	Charm := reflect.Indirect(v).FieldByName("Charm")
			// 	Tactic := reflect.Indirect(v).FieldByName("Tactic")
			// 	InheritedTactic := reflect.Indirect(v).FieldByName("InheritedTactic")
			// }

			_, _ = file.Write([]byte("\n"))
		}
	})
}
