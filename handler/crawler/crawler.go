package crawler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
)

type Role struct {
	Star            string `json:"star"`             // 星數
	Camp            string `json:"camp"`             // 陣營
	Name            string `json:"name"`             // 武將
	Cost            string `json:"cost"`             // 佔位
	Cavalry         string `json:"cavalry"`          // 騎兵
	Shield          string `json:"shield"`           // 盾兵
	Bow             string `json:"bow"`              // 弓兵
	Spear           string `json:"spear"`            // 槍兵
	Instrument      string `json:"instrument"`       // 器械
	Command         string `json:"command"`          // 統率
	Force           string `json:"force"`            // 武力
	Intelligence    string `json:"intelligence"`     // 智力
	Speed           string `json:"speed"`            // 速度
	Politics        string `json:"politics"`         // 政治
	Charm           string `json:"charm"`            // 魅力
	Tactic          string `json:"tactic"`           // 戰法
	InheritedTactic string `json:"inherited_tactic"` // 傳承戰法
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
		keyMapping := genKeyMap()
		roleList := []map[string]string{}

		// 寫入武將內容
		for i, role := range roles {
			contents := htmlquery.Find(role, `./td`)
			singleRoleInfo := map[string]string{}

			for _, content := range contents {
				contentNode := htmlquery.FindOne(content, `.`)
				_, _ = file.Write([]byte(htmlquery.InnerText(contentNode) + " "))
				singleRoleInfo[keyMapping[i]] = strings.TrimSpace(htmlquery.InnerText(contentNode))
			}

			roleList = append(roleList, singleRoleInfo)
			_, _ = file.Write([]byte("\n"))
		}

		res := []Role{}
		resData, _ := json.Marshal(roleList)
		fmt.Println(roleList)

		if err = json.Unmarshal(resData, &res); err != nil {
			log.Fatal(err)
		}
	})
}

func genKeyMap() map[int]string {
	keyMap := map[int]string{
		0:  "Star",
		1:  "Camp",
		2:  "Name",
		3:  "Cost",
		4:  "Cavalry",
		5:  "Shield",
		6:  "Bow",
		7:  "Spear",
		8:  "Instrument",
		9:  "Force",
		10: "Command",
		11: "Intelligence",
		12: "Speed",
		13: "Politics",
		14: "Charm",
		15: "Tactic",
		16: "InheritedTactic",
	}

	return keyMap
}
