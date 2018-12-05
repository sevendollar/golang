package star

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	stars = []string{
		"Aries",
		"Taurus",
		"Gemini",
		"Cancer",
		"Leo",
		"Virgo",
		"Libra",
		"Scorpio",
		"Sagittarius",
		"Capricornus",
		"Aquarius",
		"Pisces",
	}
)

func GetStar(starIndex []int, dayIndex int) (r string, err error) {
	getInfo := func(i int, dayIndex int) (r string, err error) {
		yy, mm, dd := time.Now().Add(time.Duration(dayIndex) * 24 * time.Hour).Date()
		date := strconv.Itoa(yy) + "-" + strconv.Itoa(int(mm)) + "-" + strconv.Itoa(dd)

		ii := strconv.Itoa(i)
		url := "http://astro.click108.com.tw/daily_" + ii + ".php?iAstro=" + ii + "&iType=0&iAcDay=" + date
		resp, err := http.Get(url)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			err = fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
			return
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return
		}

		r += "【" + stars[i] + "】\n"
		tmp1 := []string{}
		tmp2 := ""
		doc.Find("div[class=TODAY_CONTENT] p").Each(func(i int, s *goquery.Selection) {
			tmp1 = append(tmp1, s.Text())
		})
		for i := 0; i < 8; i++ {
			if i%4 == 0 {
				r += tmp1[i][:9] + tmp1[i][12:len(tmp1[i])-3] + "    " + tmp1[i+2][:9] + tmp1[i+2][12:len(tmp1[i])-3] + "\n"
			} else if i%2 != 0 {
				tmp2 += fmt.Sprintf("%v:\n%v\n", tmp1[i-1][:9], tmp1[i])
			}
		}
		doc.Find("div[class=TODAY_WORD]").Each(func(i int, s *goquery.Selection) {
			r += "今日短評: " + strings.TrimSpace(s.Text()) + "\n"
		})
		r += tmp2
		return
	}

	if len(starIndex) == 0 {
		err = errors.New("add some indexes")
		return
	}
	var tmp string
	for i, v := range starIndex {
		tmp, err = getInfo(v, dayIndex)
		if err != nil {
			return
		}
		r += tmp
		if i != len(starIndex)-1 {
			r += "\n"
		}
	}
	return
}
