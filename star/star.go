package star

import (
        "errors"
        "fmt"
        "net/http"
        "strconv"

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

func GetStar(index ...int) (r string, err error) {
        getInfo := func(i int) (t string, err error) {
                ii := strconv.Itoa(i)
                url := "http://astro.click108.com.tw/daily_" + ii + ".php?iAstro=" + ii
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

                t += stars[i] + ":\n"
                doc.Find("div[class=TODAY_CONTENT] p").Each(func(i int, s *goquery.Selection) {
                        t += s.Text() + "\n"
                })
                return
        }

        if len(index) == 0 {
                err = errors.New("add some indexes")
                return
        }
        var tmp string
        for i, v := range index {
                tmp, err = getInfo(v)
                if err != nil {
                        return
                }
                r += tmp
                if i != len(index)-1 {
                        r += "\n"
                }
        }
        return
}

