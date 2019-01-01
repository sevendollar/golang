package star

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	date          = ""
	starSignIndex = 0

	q              = Astro{}
	firstDate      = ""
	lastDate       = ""
	dateOutOfRange = false
)

type Astro struct {
	SunSign        string    `json:"sun_sign"`
	PredictionDate time.Time `json:"prediction_date"`
	Prediction     Prediction
}

type Prediction struct {
	ShortWord        string `json:"short_word"`
	LuckyNumber      string `json:"lucky_number"`
	LuckyColor       string `json:"lucky_color"`
	LuckyDirection   string `json:"lucky_direction"`
	LuckyTime        string `json:"lucky_time"`
	LuckyStar        string `json:"lucky_star"`
	OverviewRating   int    `json:"overview_rating"`
	Overview         string `json:"overview"`
	EmotionRating    int    `json:"emotion_rating"`
	Emotion          string `json:"emotion"`
	ProfessionRating int    `json:"profession_rating"`
	Profession       string `json:"profession"`
	FinanceRating    int    `json:"finance_rating"`
	Finance          string `json:"finance"`
}

func GetDateRange() (minDate, MaxDate string, err error) {
	yyyy, mm, dd := time.Now().Date()
	date = fmt.Sprintf("%v-%v-%v", yyyy, timeCorection(int(mm)), timeCorection(dd))
	url := "http://astro.click108.com.tw/daily_0.php?iType=0&iAstro=0&iAcDay=" + date
	res, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("%q", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = fmt.Errorf("%q", err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		err = fmt.Errorf("%q", err)
		return
	}
	dateLength := doc.Find("select[id=iAcDay] option").Length()
	doc.Find("select[id=iAcDay] option").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			firstDate = strings.TrimSpace(s.Text())
		} else if i == dateLength-1 {
			lastDate = strings.TrimSpace(s.Text())
		}
	})
	minDate, MaxDate = firstDate, lastDate
	return
}

func timeCorection(d int) string {
	if d < 10 {
		return "0" + fmt.Sprint(d)
	}
	return fmt.Sprint(d)
}

func GetPrediction(starSignIndex int, date string) (result string, err error) {
	content := map[string]string{}
	contentRange := map[string]int{}

	if date == "" {
		yyyy, mm, dd := time.Now().Date()
		date = fmt.Sprintf("%v-%v-%v", yyyy, timeCorection(int(mm)), timeCorection(dd))
	}
	if starSignIndex < 0 || starSignIndex > 11 {
		err = fmt.Errorf("%v", `starSignIndex of out range, take one of the flowing number...
aries=0
taurus=1
gemini=2
cancer=3
leo=4
virgo=5
libra=6
scorpio=7
sagittarius=8
capricorn=9
aquarius=10
pisces=11`)
		return
	}
	url := "http://astro.click108.com.tw/daily_" + strconv.Itoa(starSignIndex) + ".php?iType=0&iAstro=" + strconv.Itoa(starSignIndex) + "&iAcDay=" + date
	res, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("%q", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = fmt.Errorf("%q", err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		err = fmt.Errorf("%q", err)
		return
	}
	dateLength := doc.Find("select[id=iAcDay] option").Length()
	doc.Find("select[id=iAcDay] option").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			firstDate = strings.TrimSpace(s.Text())
		} else if i == dateLength-1 {
			lastDate = strings.TrimSpace(s.Text())
		}

		if _, ok := s.Attr("selected"); ok {
			if s.Text() == date {
				if q.PredictionDate, err = time.Parse("2006-01-02", s.Text()); err != nil {
					err = fmt.Errorf("%q", err)
					return
				}
				doc.Find("div[class=TODAY_WORD]").Each(func(i int, s *goquery.Selection) {
					q.Prediction.ShortWord = strings.TrimSpace(s.Text())
				})
				doc.Find("div[class=TODAY_LUCKY] div[class=LUCKY]").Each(func(i int, s *goquery.Selection) {
					content["lucky"+fmt.Sprint(i)] = strings.TrimSpace(s.Text())
				})
				doc.Find("div[class=TODAY_CONTENT] h3").Each(func(i int, s *goquery.Selection) {
					star := strings.TrimSpace(s.Text())
					q.SunSign = star[6 : len(star)-6]
				})
				doc.Find("div[class=TODAY_CONTENT] p").Each(func(i int, s *goquery.Selection) {
					content["content"+fmt.Sprint(i)] = strings.TrimSpace(s.Text())
				})
				for i := 0; i < 8; i++ {
					if i%2 == 0 {
						x := content["content"+fmt.Sprint(i)]
						y := 0
						xBytes := []byte(x[12 : len(x)-3])
						for j := 2; j < len(xBytes); j = j + 3 {
							if int(xBytes[j]) == 133 {
								y++
							}
						}
						contentRange["content"+fmt.Sprint(i)] = y
					}
				}

				q.Prediction.LuckyNumber = content["lucky0"]
				q.Prediction.LuckyColor = content["lucky1"]
				q.Prediction.LuckyDirection = content["lucky2"]
				q.Prediction.LuckyTime = content["lucky3"]
				q.Prediction.LuckyStar = content["lucky4"]

				q.Prediction.Overview = content["content1"]
				q.Prediction.Emotion = content["content3"]
				q.Prediction.Profession = content["content5"]
				q.Prediction.Finance = content["content7"]

				q.Prediction.OverviewRating = contentRange["content0"]
				q.Prediction.EmotionRating = contentRange["content2"]
				q.Prediction.ProfessionRating = contentRange["content4"]
				q.Prediction.FinanceRating = contentRange["content6"]

				return
			} else {
				dateOutOfRange = true
			}
		}
	})
	if dateOutOfRange {
		err = fmt.Errorf("date out of range: it should be somewhere between %q and %q", firstDate, lastDate)
		return
	}

	jsonDataByte, err := json.MarshalIndent(q, "", "    ")
	if err != nil {
		err = fmt.Errorf("%q", err)
		return
	}
	result = string(jsonDataByte)
	return
}
