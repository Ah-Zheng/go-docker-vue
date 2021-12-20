package crawler

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
)

type Tactic struct {
	Name              string `json:"name"`                // 戰法名稱
	Quality           string `json:"quality"`             // 品質
	Type              string `json:"type"`                // 類型
	LaunchProbability string `json:"launch_probability"`  // 觸發機率
	Applicable        string `json:"applicable"`          // 適用兵種
	Effect            string `json:"effect"`              // 效果
	Target            string `json:"target"`              // 目標
	BringYourOwnRole  string `json:"brind_your_own_role"` // 自帶武將
	InheritedRole     string `json:"inherited_role"`      // 傳承武將
	Event             string `json:"event"`               // 事件
	Description       string `json:"description"`         // 說明
	ToUs              string `json:"to_us"`               // 對我方
	ToEemy            string `json:"to_enemy"`            // 對敵方
}

func genTacticKeyMap() map[int]string {
	return map[int]string{
		0:  "Name",
		1:  "Quality",
		2:  "Type",
		3:  "LaunchProbability",
		4:  "Applicable",
		5:  "Effect",
		6:  "Target",
		7:  "BringYourOwnRole",
		8:  "InheritedRole",
		9:  "Event",
		10: "Description",
		11: "ToUs",
		12: "ToEemy",
	}
}

func GetTacticInfo() {
	c := colly.NewCollector()
	tacticList := []map[string]string{}

	c.OnResponse(func(r *colly.Response) {
		doc, err := htmlquery.Parse(strings.NewReader(string(r.Body)))

		if err != nil {
			log.Fatal(err)
		}

		tactics := htmlquery.Find(doc, `//*[@class="table-responsive"]/table/tbody/tr`)

		if len(tactics) > 0 {
			for _, tactic := range tactics {
				contents := htmlquery.Find(tactic, `./td`)
				link := htmlquery.FindOne(contents[1], `./a`)

				if err = r.Request.Visit(link.Attr[0].Val); err != nil {
					log.Fatal(err)
				}
			}
		}

		tacInfos := htmlquery.Find(doc, `//*[@class="tabs"]/*[@class="tab-content mt-3"]/*[@role="tabpanel"]`)
		keyMapping := genTacticKeyMap()

		if len(tacInfos) > 0 {
			singleTac := map[string]string{}
			nameNode := htmlquery.FindOne(tacInfos[0], `./span`)
			singleTac[keyMapping[0]] = strings.TrimSpace(htmlquery.InnerText(nameNode))
			trs := htmlquery.Find(tacInfos[0], `./table/tbody/tr`)

			for i, tr := range trs {
				td := htmlquery.Find(tr, `./td`)
				singleTac[keyMapping[i+1]] = strings.TrimSpace(htmlquery.InnerText(td[1]))
			}

			tacticList = append(tacticList, singleTac)
		}
	})

	if err := c.Visit("https://sgzzlbdb.com/skill"); err != nil {
		log.Fatal(err)
	}

	res := []Tactic{}
	resData, _ := json.Marshal(tacticList)

	if err := json.Unmarshal(resData, &res); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
