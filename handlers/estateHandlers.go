package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
	d "workspace/data"
	e "workspace/elasticsearch"
	"workspace/models"

	"github.com/tebeka/selenium"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func RunSelenium() {
	fmt.Println("Starting")
	fmt.Println("Waiting for goroutines to finish...")
	var wg sync.WaitGroup
	wg.Add(10)
	for x := 0; x <= 9; x++ {
		url := strconv.Itoa(x)
		urlInt := x
		a := 0
		if x != 0 {
			a = x * 4
		}
		districtData := d.DistrictData()
		go func() {
			townDataList := d.TownData()
			townData := []string{townDataList[a], townDataList[a+1], townDataList[a+2], townDataList[a+3]}
			DataSearch(townData, url, &wg, districtData, urlInt)
		}()
	}
	wg.Wait()
	fmt.Println("Done!")
}

func DataSearch(townDataList []string, url string, wg *sync.WaitGroup, districtData []string, x int) {
	var wd selenium.WebDriver
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	wd, err = selenium.NewRemote(caps, "http://host.docker.internal:981"+url)
	Error(err)
	wd.MaximizeWindow("")
	if err := wd.Get("https://www.zingat.com/istanbul-bolge-raporu"); err != nil {
		Error(err)
	}
	if wd == nil {
		for {
			wd, err = selenium.NewRemote(caps, "http://host.docker.internal:981"+url)
			Error(err)
			if err := wd.Get("https://www.zingat.com/istanbul-bolge-raporu"); err != nil {
				Error(err)
			}
			if wd != nil {
				break
			}
		}
	}
	for i := 0; i < len(townDataList); {
		pageUrl, err := wd.CurrentURL()
		Error(err)
		urlControl := strings.Contains(pageUrl, "?")

		if urlControl {
			for {
				if err := wd.Get("https://www.zingat.com/istanbul-bolge-raporu"); err != nil {
					Error(err)
				}
				pageUrlRetry, err := wd.CurrentURL()
				Error(err)
				urlControlRetry := strings.Contains(pageUrlRetry, "?")
				if !urlControlRetry {
					break
				}
			}
		}
		AnnouncementButtonClose(wd)
		DataVerification(wd, townDataList[i])

		m := &models.EstateIndex{}
		m.Name = townDataList[i]
		currentTime := time.Now().UTC()
		currentTime.Format("2006-01-02")
		m.Date = currentTime
		AnnouncementButtonClose(wd)
		//Satılık - Bölge Ortalaması

		AnnouncementButtonClose(wd)
		regionAverage, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[2]/ul/li[1]/div/p[2]/span")
		Error(err)
		if regionAverage == nil {
			AnnouncementButtonClose(wd)
			regionAverageRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[2]/ul/li[1]/div/p[2]/span")
			Error(err)
			if regionAverageRetry != nil {
				AnnouncementButtonClose(wd)
				regionData, err := regionAverageRetry.Text()
				Error(err)
				m.ForSale.RegionAverage = FloatDataParse(regionData)
				Error(err)
			}
		} else {
			AnnouncementButtonClose(wd)
			regionAverageData, err := regionAverage.Text()
			Error(err)
			m.ForSale.RegionAverage = FloatDataParse(regionAverageData)
			Error(err)
		}

		//Satılık Minimum Fiyat
		AnnouncementButtonClose(wd)
		minPrice, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[2]/ul/li[1]/div/div/div[1]/span[1]")
		Error(err)
		if minPrice == nil {
			AnnouncementButtonClose(wd)
			minPriceRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[2]/ul/li[1]/div/div/div[1]/span[1]")
			Error(err)
			if minPriceRetry != nil {
				AnnouncementButtonClose(wd)
				minPriceRetryData, err := minPriceRetry.Text()
				Error(err)
				m.ForSale.MinPrice = FloatDataParse(minPriceRetryData)
			}
		} else {
			AnnouncementButtonClose(wd)
			minPriceData, err := minPrice.Text()
			Error(err)
			m.ForSale.MinPrice = FloatDataParse(minPriceData)
		}

		//Satılık Max Fiyat
		AnnouncementButtonClose(wd)
		maxPrice, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[2]/ul/li[1]/div/div/div[2]/span[1]")
		Error(err)
		if maxPrice == nil {
			AnnouncementButtonClose(wd)
			maxPriceRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[2]/ul/li[1]/div/div/div[2]/span[1]")
			Error(err)
			if maxPriceRetry != nil {
				AnnouncementButtonClose(wd)
				maxPriceRetryData, err := maxPriceRetry.Text()
				Error(err)
				m.ForSale.MaxPirce = FloatDataParse(maxPriceRetryData)
			}
		} else {
			AnnouncementButtonClose(wd)
			maxPrice, err := maxPrice.Text()
			Error(err)
			m.ForSale.MaxPirce = FloatDataParse(maxPrice)
		}

		//Geri Dönüş Süresi
		AnnouncementButtonClose(wd)
		turnaroundTime, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[2]/ul/li[2]/div[1]/p[2]/span")
		Error(err)
		if turnaroundTime == nil {
			AnnouncementButtonClose(wd)
			turnaroundTimeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[2]/ul/li[2]/div[1]/p[2]/span")
			Error(err)
			if turnaroundTimeRetry != nil {
				AnnouncementButtonClose(wd)
				turnaroundTimeRetryData, err := turnaroundTimeRetry.Text()
				Error(err)
				m.TurnaroundTime = IntDataParse(turnaroundTimeRetryData)
			}
		} else {
			AnnouncementButtonClose(wd)
			turnaroundTimeData, err := turnaroundTime.Text()
			Error(err)
			m.TurnaroundTime = IntDataParse(turnaroundTimeData)
		}
		//1 Aylık Değişim
		oneMonthChange, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[1]/div/div[1]/div[1]/span[1]")
		Error(err)
		if oneMonthChange == nil {
			oneMonthChangeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[1]/div/div[1]/div[1]/span[1]")
			Error(err)
			if oneMonthChangeRetry != nil {
				AnnouncementButtonClose(wd)
				oneMonthChangeData, err := oneMonthChangeRetry.Text()
				Error(err)
				m.ForSale.OneMonthChange = PercentileFloatDataParse(OneMonthChangeStatus(wd) + oneMonthChangeData)
			}

		} else {
			AnnouncementButtonClose(wd)
			oneMonthChangeData, err := oneMonthChange.Text()
			Error(err)
			m.ForSale.OneMonthChange = PercentileFloatDataParse(OneMonthChangeStatus(wd) + oneMonthChangeData)
		}
		//3 Yıllık Değişim
		threeYearChange, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[1]/div/div[2]/div[1]/span[1]")
		Error(err)
		if threeYearChange == nil {
			threeYearChangeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[1]/div/div[2]/div[1]/span[1]")
			Error(err)
			if threeYearChangeRetry != nil {
				AnnouncementButtonClose(wd)
				threeYearChangeData, err := threeYearChangeRetry.Text()
				Error(err)
				m.ForSale.ThreeYearChange = PercentileFloatDataParse(ThreeYearChangeStatus(wd) + threeYearChangeData)
			}
		} else {
			AnnouncementButtonClose(wd)
			threeYearChangeData, err := threeYearChange.Text()
			Error(err)
			m.ForSale.ThreeYearChange = PercentileFloatDataParse(ThreeYearChangeStatus(wd) + threeYearChangeData)
		}

		//5 Yıllık Değişim
		fiveYearChange, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[1]/div/div[3]/div[1]/span[1]")
		Error(err)
		if fiveYearChange == nil {
			fiveYearChangeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[1]/div/div[3]/div[1]/span[1]")
			Error(err)
			if fiveYearChangeRetry != nil {
				AnnouncementButtonClose(wd)
				fiveYearChangeData, err := fiveYearChangeRetry.Text()
				Error(err)
				m.ForSale.FiveYearChange = PercentileFloatDataParse(FiveYearChangeStatus(wd) + fiveYearChangeData)
			}
		} else {
			AnnouncementButtonClose(wd)
			fiveYearChangeData, err := fiveYearChange.Text()
			Error(err)
			m.ForSale.FiveYearChange = PercentileFloatDataParse(FiveYearChangeStatus(wd) + fiveYearChangeData)
		}
		// Kiralık Button Click
		AnnouncementButtonClose(wd)
		element, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[1]/ul/li[2]")
		Error(err)
		if element == nil {
			AnnouncementButtonClose(wd)
			elementRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[1]/ul/li[2]")
			Error(err)
			if elementRetry != nil {
				AnnouncementButtonClose(wd)
				test, err := elementRetry.GetAttribute("class")
				Error(err)
				if test != "active" {
					for {
						AnnouncementButtonClose(wd)
						elementRetry.Click()
						AnnouncementButtonClose(wd)
						activeControl, err := elementRetry.GetAttribute("class")
						Error(err)
						if activeControl == "active" {
							break
						}
					}
				}
			}
		} else {
			AnnouncementButtonClose(wd)
			test, err := element.GetAttribute("class")
			Error(err)
			if test != "active" {

				for {
					AnnouncementButtonClose(wd)
					element.Click()
					AnnouncementButtonClose(wd)
					activeControl, err := element.GetAttribute("class")
					Error(err)
					if activeControl == "active" {
						break
					}
				}
			}
		}
		//Kiralık - Bölge Ortalaması

		AnnouncementButtonClose(wd)
		forRentregionAverage, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[3]/ul/li[1]/div/p[2]/span")
		Error(err)
		if forRentregionAverage == nil {
			AnnouncementButtonClose(wd)
			forRentregionAverageRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[3]/ul/li[1]/div/p[2]/span")
			Error(err)
			if forRentregionAverageRetry != nil {
				AnnouncementButtonClose(wd)
				forRentregionAverageRetryData, err := forRentregionAverageRetry.Text()
				Error(err)
				m.ForRent.RegionAverage = FloatDataParse(forRentregionAverageRetryData)
			}

		} else {
			AnnouncementButtonClose(wd)
			forRentregionAverageData, err := forRentregionAverage.Text()
			m.ForRent.RegionAverage = FloatDataParse(forRentregionAverageData)
			Error(err)
		}

		AnnouncementButtonClose(wd)
		//Kiralık Minimum Fiyat
		forRentminPrice, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[3]/ul/li[1]/div/div/div[1]/span[1]")
		Error(err)
		if forRentminPrice == nil {
			AnnouncementButtonClose(wd)
			forRentminPriceRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[3]/ul/li[1]/div/div/div[1]/span[1]")
			Error(err)
			if forRentminPriceRetry != nil {
				AnnouncementButtonClose(wd)
				forRentminPriceRetryData, err := forRentminPriceRetry.Text()
				m.ForRent.MinPrice = FloatDataParse(forRentminPriceRetryData)
				Error(err)
			}
		} else {
			AnnouncementButtonClose(wd)
			forRentminPriceData, err := forRentminPrice.Text()
			m.ForRent.MinPrice = FloatDataParse(forRentminPriceData)
			Error(err)
		}

		AnnouncementButtonClose(wd)
		//Kiralık Max Fiyat
		forRentmaxPrice, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[3]/ul/li[1]/div/div/div[2]/span[1]")
		Error(err)
		if forRentmaxPrice == nil {
			AnnouncementButtonClose(wd)
			forRentmaxPriceRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[3]/ul/li[1]/div/div/div[2]/span[1]")
			Error(err)
			if forRentmaxPriceRetry != nil {
				AnnouncementButtonClose(wd)
				forRentmaxPriceRetryData, err := forRentmaxPriceRetry.Text()
				Error(err)
				m.ForRent.MaxPirce = FloatDataParse(forRentmaxPriceRetryData)
			}
		} else {
			AnnouncementButtonClose(wd)
			forRentmaxPriceData, err := forRentmaxPrice.Text()
			Error(err)
			m.ForRent.MaxPirce = FloatDataParse(forRentmaxPriceData)
		}

		AnnouncementButtonClose(wd)
		//1 Aylık Değişim
		forRentoneMonthChange, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[2]/div/div[1]/div[1]/span[1]")
		Error(err)
		if forRentoneMonthChange == nil {
			forRentoneMonthChangeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[2]/div/div[1]/div[1]/span[1]")
			Error(err)
			if forRentoneMonthChangeRetry != nil {
				AnnouncementButtonClose(wd)
				forRentoneMonthChangeData, err := forRentoneMonthChangeRetry.Text()
				Error(err)
				m.ForRent.OneMonthChange = PercentileFloatDataParse(OneMonthChangeStatus(wd) + forRentoneMonthChangeData)
			}
		} else {
			AnnouncementButtonClose(wd)
			forRentoneMonthChangeData, err := forRentoneMonthChange.Text()
			Error(err)
			m.ForRent.OneMonthChange = PercentileFloatDataParse(OneMonthChangeStatus(wd) + forRentoneMonthChangeData)

		}

		//3 Yıllık Değişim
		forRentthreeYearChange, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[2]/div/div[2]/div[1]/span[1]")
		Error(err)
		if forRentthreeYearChange == nil {
			forRentthreeYearChangeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[2]/div/div[2]/div[1]/span[1]")
			Error(err)
			if forRentthreeYearChangeRetry != nil {
				AnnouncementButtonClose(wd)
				forRentthreeYearChangeData, err := forRentthreeYearChangeRetry.Text()
				Error(err)
				m.ForRent.ThreeYearChange = PercentileFloatDataParse(ThreeYearChangeStatus(wd) + forRentthreeYearChangeData)
			}
		} else {
			AnnouncementButtonClose(wd)
			forRentthreeYearChangeData, err := forRentthreeYearChange.Text()
			Error(err)
			m.ForRent.ThreeYearChange = PercentileFloatDataParse(ThreeYearChangeStatus(wd) + forRentthreeYearChangeData)
		}

		//5 Yıllık Değişim
		forRentfiveYearChange, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[2]/div/div[3]/div[1]/span[1]")
		Error(err)
		if forRentfiveYearChange == nil {
			forRentfiveYearChangeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[2]/div/div[3]/div[1]/span[1]")
			Error(err)
			if forRentfiveYearChangeRetry != nil {
				AnnouncementButtonClose(wd)
				forRentfiveYearChangeData, err := forRentfiveYearChangeRetry.Text()
				Error(err)
				m.ForRent.FiveYearChange = PercentileFloatDataParse(FiveYearChangeStatus(wd) + forRentfiveYearChangeData)
			}
		} else {
			AnnouncementButtonClose(wd)
			forRentfiveYearChangeData, err := forRentfiveYearChange.Text()
			Error(err)
			m.ForRent.FiveYearChange = PercentileFloatDataParse(FiveYearChangeStatus(wd) + forRentfiveYearChangeData)
		}

		if m.TurnaroundTime == 0 || m.ForRent.MaxPirce == 0 || m.ForRent.MinPrice == 0 || m.ForRent.RegionAverage == 0 ||
			m.ForSale.MaxPirce == 0 || m.ForSale.MinPrice == 0 || m.ForSale.RegionAverage == 0 {
		} else {
			InsertTownEstateindex(*m)
			i++
		}
	}
	districtDataSearchTest(wd, districtData, x)
	wg.Done()
}

func districtDataSearchTest(wd selenium.WebDriver, data []string, x int) {
	StartLen := x * 42
	LastLen := (x + 1) * 42
	if x == 9 {
		LastLen = 422
	}

	for i := StartLen; i < LastLen; {
		AnnouncementButtonClose(wd)
		DataVerification(wd, data[i])

		m := &models.EstateIndex{}
		m.Name = data[i]
		currentTime := time.Now().UTC()
		currentTime.Format("2006-01-02")
		m.Date = currentTime
		AnnouncementButtonClose(wd)
		//Satılık - Bölge Ortalaması

		AnnouncementButtonClose(wd)
		regionAverage, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[2]/ul/li[1]/div/p[2]/span")
		Error(err)
		if regionAverage == nil {
			AnnouncementButtonClose(wd)
			regionAverageRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[2]/ul/li[1]/div/p[2]/span")
			Error(err)
			if regionAverageRetry != nil {
				AnnouncementButtonClose(wd)
				regionData, err := regionAverageRetry.Text()
				Error(err)
				m.ForSale.RegionAverage = FloatDataParse(regionData)
				Error(err)
			}
		} else {
			AnnouncementButtonClose(wd)
			regionAverageData, err := regionAverage.Text()
			Error(err)
			m.ForSale.RegionAverage = FloatDataParse(regionAverageData)
			Error(err)
		}

		//Satılık Minimum Fiyat
		AnnouncementButtonClose(wd)
		minPrice, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[2]/ul/li[1]/div/div/div[1]/span[1]")
		Error(err)
		if minPrice == nil {
			AnnouncementButtonClose(wd)
			minPriceRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[2]/ul/li[1]/div/div/div[1]/span[1]")
			Error(err)
			if minPriceRetry != nil {
				AnnouncementButtonClose(wd)
				minPriceRetryData, err := minPriceRetry.Text()
				Error(err)
				m.ForSale.MinPrice = FloatDataParse(minPriceRetryData)
			}
		} else {
			AnnouncementButtonClose(wd)
			minPriceData, err := minPrice.Text()
			Error(err)
			m.ForSale.MinPrice = FloatDataParse(minPriceData)
		}

		//Satılık Max Fiyat
		AnnouncementButtonClose(wd)
		maxPrice, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[2]/ul/li[1]/div/div/div[2]/span[1]")
		Error(err)
		if maxPrice == nil {
			AnnouncementButtonClose(wd)
			maxPriceRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[2]/ul/li[1]/div/div/div[2]/span[1]")
			Error(err)
			if maxPriceRetry != nil {
				AnnouncementButtonClose(wd)
				maxPriceRetryData, err := maxPriceRetry.Text()
				Error(err)
				m.ForSale.MaxPirce = FloatDataParse(maxPriceRetryData)
			}
		} else {
			AnnouncementButtonClose(wd)
			maxPrice, err := maxPrice.Text()
			Error(err)
			m.ForSale.MaxPirce = FloatDataParse(maxPrice)
		}

		//Geri Dönüş Süresi
		AnnouncementButtonClose(wd)
		turnaroundTime, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[2]/ul/li[2]/div[1]/p[2]/span")
		Error(err)
		if turnaroundTime == nil {
			AnnouncementButtonClose(wd)
			turnaroundTimeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[2]/ul/li[2]/div[1]/p[2]/span")
			Error(err)
			if turnaroundTimeRetry != nil {
				AnnouncementButtonClose(wd)
				turnaroundTimeRetryData, err := turnaroundTimeRetry.Text()
				Error(err)
				m.TurnaroundTime = IntDataParse(turnaroundTimeRetryData)
			}
		} else {
			AnnouncementButtonClose(wd)
			turnaroundTimeData, err := turnaroundTime.Text()
			Error(err)
			m.TurnaroundTime = IntDataParse(turnaroundTimeData)
		}

		// Satılık Fiyatları Değişimi
		//1 Aylık Değişim
		oneMonthChange, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[1]/div/div[1]/div[1]/span[1]")
		Error(err)
		if oneMonthChange == nil {
			oneMonthChangeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[1]/div/div[1]/div[1]/span[1]")
			Error(err)
			if oneMonthChangeRetry != nil {
				AnnouncementButtonClose(wd)
				oneMonthChangeData, err := oneMonthChangeRetry.Text()
				Error(err)
				m.ForSale.OneMonthChange = PercentileFloatDataParse(OneMonthChangeStatus(wd) + oneMonthChangeData)
			}
		} else {
			AnnouncementButtonClose(wd)
			oneMonthChangeData, err := oneMonthChange.Text()
			Error(err)
			m.ForSale.OneMonthChange = PercentileFloatDataParse(OneMonthChangeStatus(wd) + oneMonthChangeData)
		}

		//3 Yıllık Değişim
		threeYearChange, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[1]/div/div[2]/div[1]/span[1]")
		Error(err)
		if threeYearChange == nil {
			threeYearChangeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[1]/div/div[2]/div[1]/span[1]")
			Error(err)
			if threeYearChangeRetry != nil {
				AnnouncementButtonClose(wd)
				threeYearChangeData, err := threeYearChangeRetry.Text()
				Error(err)
				m.ForSale.ThreeYearChange = PercentileFloatDataParse(ThreeYearChangeStatus(wd) + threeYearChangeData)
			}
		} else {
			AnnouncementButtonClose(wd)
			threeYearChangeData, err := threeYearChange.Text()
			Error(err)
			m.ForSale.ThreeYearChange = PercentileFloatDataParse(ThreeYearChangeStatus(wd) + threeYearChangeData)
		}

		//5 Yıllık Değişim
		fiveYearChange, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[1]/div/div[3]/div[1]/span[1]")
		Error(err)
		if fiveYearChange == nil {
			fiveYearChangeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[1]/div/div[3]/div[1]/span[1]")
			Error(err)
			if fiveYearChangeRetry != nil {
				AnnouncementButtonClose(wd)
				fiveYearChangeData, err := fiveYearChangeRetry.Text()
				Error(err)
				m.ForSale.FiveYearChange = PercentileFloatDataParse(FiveYearChangeStatus(wd) + fiveYearChangeData)
			}
		} else {
			AnnouncementButtonClose(wd)
			fiveYearChangeData, err := fiveYearChange.Text()
			Error(err)
			m.ForSale.FiveYearChange = PercentileFloatDataParse(FiveYearChangeStatus(wd) + fiveYearChangeData)
		}

		// Kiralık Button Click
		AnnouncementButtonClose(wd)
		element, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[1]/ul/li[2]")
		Error(err)
		if element == nil {
			AnnouncementButtonClose(wd)
			elementRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[1]/ul/li[2]")
			Error(err)
			if elementRetry != nil {
				AnnouncementButtonClose(wd)
				test, err := elementRetry.GetAttribute("class")
				Error(err)
				if test != "active" {
					for {
						AnnouncementButtonClose(wd)
						elementRetry.Click()
						AnnouncementButtonClose(wd)
						activeControl, err := elementRetry.GetAttribute("class")
						Error(err)
						if activeControl == "active" {
							break
						}
					}
				}
			}
		} else {
			AnnouncementButtonClose(wd)
			test, err := element.GetAttribute("class")
			Error(err)
			if test != "active" {

				for {
					AnnouncementButtonClose(wd)
					element.Click()
					AnnouncementButtonClose(wd)
					activeControl, err := element.GetAttribute("class")
					Error(err)
					if activeControl == "active" {
						break
					}
				}
			}
		}
		//Kiralık - Bölge Ortalaması

		AnnouncementButtonClose(wd)
		forRentregionAverage, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[3]/ul/li[1]/div/p[2]/span")
		Error(err)
		if forRentregionAverage == nil {
			AnnouncementButtonClose(wd)
			forRentregionAverageRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[3]/ul/li[1]/div/p[2]/span")
			Error(err)
			if forRentregionAverageRetry != nil {
				AnnouncementButtonClose(wd)
				forRentregionAverageRetryData, err := forRentregionAverageRetry.Text()
				Error(err)
				m.ForRent.RegionAverage = FloatDataParse(forRentregionAverageRetryData)
			}

		} else {
			AnnouncementButtonClose(wd)
			forRentregionAverageData, err := forRentregionAverage.Text()
			m.ForRent.RegionAverage = FloatDataParse(forRentregionAverageData)
			Error(err)
		}

		AnnouncementButtonClose(wd)
		//Kiralık Minimum Fiyat
		forRentminPrice, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[3]/ul/li[1]/div/div/div[1]/span[1]")
		Error(err)
		if forRentminPrice == nil {
			AnnouncementButtonClose(wd)
			forRentminPriceRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[3]/ul/li[1]/div/div/div[1]/span[1]")
			Error(err)
			if forRentminPriceRetry != nil {
				AnnouncementButtonClose(wd)
				forRentminPriceRetryData, err := forRentminPriceRetry.Text()
				m.ForRent.MinPrice = FloatDataParse(forRentminPriceRetryData)
				Error(err)
			}
		} else {
			AnnouncementButtonClose(wd)
			forRentminPriceData, err := forRentminPrice.Text()
			m.ForRent.MinPrice = FloatDataParse(forRentminPriceData)
			Error(err)
		}

		AnnouncementButtonClose(wd)
		//Kiralık Max Fiyat
		forRentmaxPrice, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[3]/div/div/div/div/div[2]/div[3]/ul/li[1]/div/div/div[2]/span[1]")
		Error(err)
		if forRentmaxPrice == nil {
			AnnouncementButtonClose(wd)
			forRentmaxPriceRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div/div[2]/div[3]/ul/li[1]/div/div/div[2]/span[1]")
			Error(err)
			if forRentmaxPriceRetry != nil {
				AnnouncementButtonClose(wd)
				forRentmaxPriceRetryData, err := forRentmaxPriceRetry.Text()
				Error(err)
				m.ForRent.MaxPirce = FloatDataParse(forRentmaxPriceRetryData)
			}
		} else {
			AnnouncementButtonClose(wd)
			forRentmaxPriceData, err := forRentmaxPrice.Text()
			Error(err)
			m.ForRent.MaxPirce = FloatDataParse(forRentmaxPriceData)
		}

		AnnouncementButtonClose(wd)
		//1 Aylık Değişim
		forRentoneMonthChange, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[2]/div/div[1]/div[1]/span[1]")
		Error(err)
		if forRentoneMonthChange == nil {
			forRentoneMonthChangeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[2]/div/div[1]/div[1]/span[1]")
			Error(err)
			if forRentoneMonthChangeRetry != nil {
				AnnouncementButtonClose(wd)
				forRentoneMonthChangeData, err := forRentoneMonthChangeRetry.Text()
				Error(err)
				m.ForRent.OneMonthChange = PercentileFloatDataParse(OneMonthChangeStatus(wd) + forRentoneMonthChangeData)
			}
		} else {
			AnnouncementButtonClose(wd)
			forRentoneMonthChangeData, err := forRentoneMonthChange.Text()
			Error(err)
			m.ForRent.OneMonthChange = PercentileFloatDataParse(OneMonthChangeStatus(wd) + forRentoneMonthChangeData)
		}

		//3 Yıllık Değişim
		forRentthreeYearChange, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[2]/div/div[2]/div[1]/span[1]")
		Error(err)
		if forRentthreeYearChange == nil {
			forRentthreeYearChangeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[2]/div/div[2]/div[1]/span[1]")
			Error(err)
			if forRentthreeYearChangeRetry != nil {
				AnnouncementButtonClose(wd)
				forRentthreeYearChangeData, err := forRentthreeYearChangeRetry.Text()
				Error(err)
				m.ForRent.ThreeYearChange = PercentileFloatDataParse(ThreeYearChangeStatus(wd) + forRentthreeYearChangeData)
			}
		} else {
			AnnouncementButtonClose(wd)
			forRentthreeYearChangeData, err := forRentthreeYearChange.Text()
			Error(err)
			m.ForRent.ThreeYearChange = PercentileFloatDataParse(ThreeYearChangeStatus(wd) + forRentthreeYearChangeData)
		}

		//5 Yıllık Değişim
		forRentfiveYearChange, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[2]/div/div[3]/div[1]/span[1]")
		Error(err)
		if forRentfiveYearChange == nil {
			forRentfiveYearChangeRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[2]/div/div[3]/div[1]/span[1]")
			Error(err)
			if forRentfiveYearChangeRetry != nil {
				AnnouncementButtonClose(wd)
				forRentfiveYearChangeData, err := forRentfiveYearChangeRetry.Text()
				Error(err)
				m.ForRent.FiveYearChange = PercentileFloatDataParse(FiveYearChangeStatus(wd) + forRentfiveYearChangeData)
			}
		} else {
			AnnouncementButtonClose(wd)
			forRentfiveYearChangeData, err := forRentfiveYearChange.Text()
			Error(err)
			m.ForRent.FiveYearChange = PercentileFloatDataParse(FiveYearChangeStatus(wd) + forRentfiveYearChangeData)
		}

		if m.TurnaroundTime == 0 || m.ForRent.MaxPirce == 0 || m.ForRent.MinPrice == 0 || m.ForRent.RegionAverage == 0 ||
			m.ForSale.MaxPirce == 0 || m.ForSale.MinPrice == 0 || m.ForSale.RegionAverage == 0 {

		} else {
			InsertDistrictEstateindex(*m)
			i++
		}
	}
}

func AnnouncementButtonClose(wd selenium.WebDriver) {
	announcementButton, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div/div/button")
	Error(err)
	if announcementButton != nil {
		announcementButton.Click()
	}

	announcementCloseButton, err := wd.FindElement(selenium.ByClassName, "tingle-modal__close")
	Error(err)
	if announcementCloseButton != nil {
		announcementCloseButton.Click()
	}

	slidedownButton, err := wd.FindElement(selenium.ByID, "onesignal-slidedown-cancel-button")
	Error(err)
	if slidedownButton != nil {
		slidedownButton.Click()
	}

	userButton, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[6]/div[2]/div/span")
	Error(err)
	if userButton != nil {
		userButton.Click()
	}
}

func Error(err error) {
}

func IntDataParse(dataParse string) int {
	dataTrimEndSplit := strings.Split(dataParse, "YIL")
	stringFloatData := ""

	stringFloatData = dataTrimEndSplit[0]
	stringFloatData = strings.TrimSpace(stringFloatData)

	dataParsed, errs := strconv.Atoi(stringFloatData)
	Error(errs)
	return dataParsed
}

func FloatDataParse(dataParse string) float64 {
	dataTrimEndSplit := strings.Split(dataParse, "TL")
	dataTrimEndSplit = strings.Split(dataTrimEndSplit[0], ".")
	stringFloatData := ""
	if len(dataTrimEndSplit) == 3 {
		stringFloatData = dataTrimEndSplit[0] + dataTrimEndSplit[1] + dataTrimEndSplit[2]
		stringFloatData = strings.TrimSpace(stringFloatData)
	} else if len(dataTrimEndSplit) == 2 {
		stringFloatData = dataTrimEndSplit[0] + dataTrimEndSplit[1]
		stringFloatData = strings.TrimSpace(stringFloatData)
	} else if len(dataTrimEndSplit) == 1 {
		stringFloatData = dataTrimEndSplit[0]
		stringFloatData = strings.TrimSpace(stringFloatData)
	}

	dataParsed, errs := strconv.ParseFloat(stringFloatData, 32)
	Error(errs)
	return dataParsed
}

func PercentileFloatDataParse(dataParse string) float64 {
	var dataParsed float64
	var errs error
	dataTrimEndSplit := strings.Split(dataParse, "%")
	stringFloatData := ""
	if len(dataTrimEndSplit) == 3 {
		stringFloatData = dataTrimEndSplit[0] + dataTrimEndSplit[1] + dataTrimEndSplit[2]
		stringFloatData = strings.TrimSpace(stringFloatData)
	} else if len(dataTrimEndSplit) == 2 {
		stringFloatData = dataTrimEndSplit[0] + dataTrimEndSplit[1]
		stringFloatData = strings.TrimSpace(stringFloatData)
	} else if len(dataTrimEndSplit) == 1 {
		stringFloatData = dataTrimEndSplit[0]
		stringFloatData = strings.TrimSpace(stringFloatData)
	}
	lastParse := strings.Contains(stringFloatData, ".")
	if lastParse {
	} else {
		stringFloatData = stringFloatData + ".1"
	}
	dataParsed, errs = strconv.ParseFloat(stringFloatData, 64)
	Error(errs)
	return dataParsed
}

func InsertDistrictEstateindex(model models.EstateIndex) {
	go func() {
		esclient, errors := e.GetESClient()
		Error(errors)
		dataJSON, err := json.Marshal(&model)
		Error(err)
		js := string(dataJSON)
		ind, err := esclient.Index().
			Index("districtestateindex").Type("EstateIndex").
			BodyJson(js).
			Do(context.Background())
		if err != nil {
			fmt.Println("districtestateindex", err)
		}
		if ind.Shards.Successful == 1 {
		}
	}()
}

func InsertTownEstateindex(model models.EstateIndex) {
	go func() {
		esclient, errors := e.GetESClient()
		Error(errors)
		dataJSON, err := json.Marshal(&model)
		Error(err)
		js := string(dataJSON)
		ind, err := esclient.Index().
			Index("townestateindex").Type("EstateIndex").
			BodyJson(js).
			Do(context.Background())
		if err != nil {
			fmt.Println("townestateindex", err)
		}
		if ind.Shards.Successful == 1 {
		}
	}()
}

func DataVerification(wd selenium.WebDriver, data string) {
	searchDistrict := ""
	searchDistrict = strings.TrimSpace(data)
	searchDistrictLower := strings.ToLower(searchDistrict)
	AnnouncementButtonClose(wd)
	var input selenium.WebElement
	var err error
	input, err = wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[2]/div/div/div/div[1]/div/div/form/input")
	Error(err)
	if input == nil {
		for {
			wd.Refresh()
			input, err = wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[2]/div/div/div/div[1]/div/div/form/input")
			Error(err)
			if input != nil {
				break
			}
		}
	} else {
		AnnouncementButtonClose(wd)
		input.Clear()
		AnnouncementButtonClose(wd)
		input.SendKeys(searchDistrict)
		AnnouncementButtonClose(wd)
		time.Sleep(1 * time.Second)
		AnnouncementButtonClose(wd)
		input.SendKeys(selenium.EnterKey)
		pageUrl, err := wd.CurrentURL()
		Error(err)
		pageUrlSplits := strings.Split(pageUrl, "https://www.zingat.com/")
		var pageUrlSplit []string
		if len(pageUrlSplits) == 1 {
			pageUrlSplit = strings.Split(pageUrlSplits[0], "-bolge-raporu")
		} else {
			pageUrlSplit = strings.Split(pageUrlSplits[1], "-bolge-raporu")
		}
		isMn := func(r rune) bool {
			return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
		}
		t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
		r := strings.NewReader(searchDistrictLower)
		x := transform.NewReader(r, t)
		b, err := ioutil.ReadAll(x)
		searchDistrictLowers := strings.ToLower(string(b))
		Error(err)
		actualdataSplit := strings.Split(searchDistrictLowers, " ")
		actualDataCount := len(actualdataSplit)
		actualLastData := ""
		if actualDataCount == 1 {
			actualLastData = actualdataSplit[0]
		} else if actualDataCount == 2 {
			actualLastData = actualdataSplit[1] + actualdataSplit[0]
		} else if actualDataCount == 3 {
			actualLastData = actualdataSplit[2] + actualdataSplit[0] + actualdataSplit[1]
		}
		searchDataSplit := ""
		if len(pageUrlSplit) == 1 {
			searchDataSplit = pageUrlSplit[0]
		} else {
			searchDataSplit = pageUrlSplit[0] + pageUrlSplit[1]
		}
		splitDataUrl := strings.Split(searchDataSplit, "-")
		expectedData := ""
		count := len(splitDataUrl)
		if count == 1 {
			expectedData = splitDataUrl[0]
		} else if count == 2 {
			expectedData = splitDataUrl[0] + splitDataUrl[1]
		} else if count == 3 {
			expectedData = splitDataUrl[0] + splitDataUrl[1] + splitDataUrl[2]
		}
		verification := strings.Contains(actualLastData, expectedData)
		if !verification {
			for {
				AnnouncementButtonClose(wd)
				input.Clear()
				AnnouncementButtonClose(wd)
				input.SendKeys(searchDistrictLower)
				time.Sleep(1 * time.Second)
				AnnouncementButtonClose(wd)
				input.SendKeys(selenium.EnterKey)
				actualpageUrl, err := wd.CurrentURL()
				Error(err)
				actualpageUrlSplits := strings.Split(actualpageUrl, "https://www.zingat.com/")
				var actualpageUrlSplit []string
				if len(pageUrlSplits) == 1 {
					actualpageUrlSplit = strings.Split(actualpageUrlSplits[0], "-bolge-raporu")
				} else {
					actualpageUrlSplit = strings.Split(actualpageUrlSplits[1], "-bolge-raporu")
				}
				expectedDataSplit := strings.Split(searchDistrictLowers, " ")
				expectedDataCount := len(expectedDataSplit)
				expectedDataRetry := ""
				if expectedDataCount == 1 {
					expectedDataRetry = expectedDataSplit[0]
				} else if expectedDataCount == 2 {
					expectedDataRetry = expectedDataSplit[1] + expectedDataSplit[0]
				} else if expectedDataCount == 3 {
					expectedDataRetry = expectedDataSplit[2] + expectedDataSplit[0] + expectedDataSplit[1]
				}

				actualLastSplitData := strings.Split(actualpageUrlSplit[0], "-")
				actualDataCount := len(actualLastSplitData)
				actualDataRetry := ""
				if actualDataCount == 1 {
					actualDataRetry = actualLastSplitData[0]
				} else if actualDataCount == 2 {
					actualDataRetry = actualLastSplitData[0] + actualLastSplitData[1]
				} else if actualDataCount == 3 {
					actualDataRetry = actualLastSplitData[0] + actualLastSplitData[1] + actualLastSplitData[2]
				}

				verificationRetry := strings.Contains(actualDataRetry, expectedDataRetry)
				if verificationRetry {
					break
				}
			}
		}
	}
}

func OneMonthChangeStatus(wd selenium.WebDriver) string {
	Status := ""
	forSaleoneMonthChangeImage, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[1]/div/div[1]/div[1]/span[2]")
	Error(err)
	if forSaleoneMonthChangeImage == nil {
		forSaleoneMonthChangeImageRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[1]/div/div[1]/div[1]/span[2]")
		Error(err)
		if forSaleoneMonthChangeImageRetry != nil {
			forSaleoneMonthChangeImageUrl, err := forSaleoneMonthChangeImageRetry.GetAttribute("style")
			Error(err)
			StateUp := strings.Contains(forSaleoneMonthChangeImageUrl, "price-up")
			StateDown := strings.Contains(forSaleoneMonthChangeImageUrl, "price-down")
			if StateUp {

			} else if StateDown {
				Status = "-"
			} else {
				Status = "0.1"
			}
		}
	} else {
		forSaleoneMonthChangeImageUrl, err := forSaleoneMonthChangeImage.GetAttribute("style")
		Error(err)
		StateUp := strings.Contains(forSaleoneMonthChangeImageUrl, "price-up")
		StateDown := strings.Contains(forSaleoneMonthChangeImageUrl, "price-down")
		if StateUp {

		} else if StateDown {
			Status = "-"
		} else {
			Status = "0.1"
		}
	}
	return Status
}

func ThreeYearChangeStatus(wd selenium.WebDriver) string {
	Status := ""
	forSaleoneMonthChangeImage, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[2]/div/div[2]/div[1]/span[2]")
	Error(err)
	if forSaleoneMonthChangeImage == nil {
		forSaleoneMonthChangeImageRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[2]/div/div[2]/div[1]/span[2]")
		Error(err)
		if forSaleoneMonthChangeImageRetry != nil {
			forSaleoneMonthChangeImageUrl, err := forSaleoneMonthChangeImageRetry.GetAttribute("style")
			Error(err)
			StateUp := strings.Contains(forSaleoneMonthChangeImageUrl, "price-up")
			StateDown := strings.Contains(forSaleoneMonthChangeImageUrl, "price-down")
			if StateUp {

			} else if StateDown {
				Status = "-"
			} else {
				Status = "0.1"
			}
		}
	} else {
		forSaleoneMonthChangeImageUrl, err := forSaleoneMonthChangeImage.GetAttribute("style")
		Error(err)
		StateUp := strings.Contains(forSaleoneMonthChangeImageUrl, "price-up")
		StateDown := strings.Contains(forSaleoneMonthChangeImageUrl, "price-down")
		if StateUp {

		} else if StateDown {
			Status = "-"
		} else {
			Status = "0.1"
		}
	}
	return Status
}

func FiveYearChangeStatus(wd selenium.WebDriver) string {
	Status := ""
	forSaleoneMonthChangeImage, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[4]/div/div/div/div[2]/div/div[3]/div[1]/span[2]")
	Error(err)
	if forSaleoneMonthChangeImage == nil {
		forSaleoneMonthChangeImageRetry, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div/section[5]/div/div/div/div[2]/div/div[3]/div[1]/span[2]")
		Error(err)
		if forSaleoneMonthChangeImageRetry != nil {
			forSaleoneMonthChangeImageUrl, err := forSaleoneMonthChangeImageRetry.GetAttribute("style")
			Error(err)
			StateUp := strings.Contains(forSaleoneMonthChangeImageUrl, "price-up")
			StateDown := strings.Contains(forSaleoneMonthChangeImageUrl, "price-down")
			if StateUp {

			} else if StateDown {
				Status = "-"
			} else {
				Status = "0.1"
			}
		}
	} else {
		forSaleoneMonthChangeImageUrl, err := forSaleoneMonthChangeImage.GetAttribute("style")
		Error(err)
		StateUp := strings.Contains(forSaleoneMonthChangeImageUrl, "price-up")
		StateDown := strings.Contains(forSaleoneMonthChangeImageUrl, "price-down")
		if StateUp {

		} else if StateDown {
			Status = "-"
		} else {
			Status = "0.1"
		}
	}
	return Status
}
