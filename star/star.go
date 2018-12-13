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

        q                     = Astro{}
        firstDate             = ""
        lastDate              = ""
        dateOutOfRange        = false
        todayLuckyTmp         = []string{}
        todayContentTmp       = []string{}
        todayContentRatingTmp = []int{}
        todayContent          = []string{}
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
        yyyy, mm, hh := time.Now().Date()
        date = fmt.Sprintf("%v-%v-%v", yyyy, int(mm), hh)
        url := "http://astro.click108.com.tw/daily_10.php?iType=0&iAstro=0&iAcDay=" + date
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

func GetStar(starSignIndex int, date string) (result string, err error) {
        if date == "" {
                yyyy, mm, hh := time.Now().Date()
                date = fmt.Sprintf("%v-%v-%v", yyyy, int(mm), hh)
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
        url := "http://astro.click108.com.tw/daily_10.php?iType=0&iAstro=" + strconv.Itoa(starSignIndex) + "&iAcDay=" + date
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
                                if q.PredictionDate, err = time.Parse("2006-1-2", s.Text()); err != nil {
                                        err = fmt.Errorf("%q", err)
                                        return
                                }
                                doc.Find("div[class=TODAY_WORD]").Each(func(i int, s *goquery.Selection) {
                                        q.Prediction.ShortWord = strings.TrimSpace(s.Text())
                                })
                                doc.Find("div[class=TODAY_LUCKY] div[class=LUCKY]").Each(func(i int, s *goquery.Selection) {
                                        todayLuckyTmp = append(todayLuckyTmp, strings.TrimSpace(s.Text()))
                                })
                                doc.Find("div[class=TODAY_CONTENT] h3").Each(func(i int, s *goquery.Selection) {
                                        star := strings.TrimSpace(s.Text())
                                        q.SunSign = star[6 : len(star)-6]
                                })
                                todayContentLength := doc.Find("div[class=TODAY_CONTENT] p").Length()
                                doc.Find("div[class=TODAY_CONTENT] p").Each(func(i int, s *goquery.Selection) {
                                        todayContent = append(todayContent, strings.TrimSpace(s.Text()))
                                })
                                for i := 0; i < todayContentLength; i++ {
                                        if i%2 == 0 {
                                                x := []byte(todayContent[i][12 : len(todayContent[i])-3])
                                                xTmp := 0
                                                for xi := 2; xi < len(x); xi = xi + 3 {
                                                        if int(x[xi]) == 133 {
                                                                xTmp++
                                                        }
                                                }
                                                todayContentRatingTmp = append(todayContentRatingTmp, xTmp)
                                        } else {
                                                todayContentTmp = append(todayContentTmp, todayContent[i])
                                        }

                                }
                                q.Prediction.LuckyNumber = todayLuckyTmp[0]
                                q.Prediction.LuckyColor = todayLuckyTmp[1]
                                q.Prediction.LuckyDirection = todayLuckyTmp[2]
                                q.Prediction.LuckyTime = todayLuckyTmp[3]
                                q.Prediction.LuckyStar = todayLuckyTmp[4]

                                q.Prediction.Overview = todayContentTmp[0]
                                q.Prediction.Emotion = todayContentTmp[1]
                                q.Prediction.Profession = todayContentTmp[2]
                                q.Prediction.Finance = todayContentTmp[3]

                                q.Prediction.OverviewRating = todayContentRatingTmp[0]
                                q.Prediction.EmotionRating = todayContentRatingTmp[1]
                                q.Prediction.ProfessionRating = todayContentRatingTmp[2]
                                q.Prediction.FinanceRating = todayContentRatingTmp[3]

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
