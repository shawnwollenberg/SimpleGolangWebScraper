//////////////////////////////////////////////////////////////////////////
//
// This has been added to the scheduled process for posting reports to GRS
//
//////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io/ioutil"
	"logger"
	"net/http"
	"os"
	"os/exec"

	"strings"
	"time"

	mtproxy "MTProxy"

	"github.com/jasonlvhit/gocron"
	/*
		mtproxy "MTProxy"
		"os"
		"os/exec"

		"fmt"
		"io/ioutil"
		"log"
		"net/http"
		"strings"
	*/)

func scrapeSPRussell() {
	var indexLookup = [][]string{}
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/%5EGSPC/history/", "S&PData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/%5ERUT/history/", "Russell2000Data"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/MTD/history/", "MTDData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/TMO/history/", "TMOData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/WAT/history/", "WATData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/ROK/history/", "ROKData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/EMR/history/", "EMRData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/DHR/history/", "DHRData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/PKI/history/", "PKIData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/ITW/history/", "ITWData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/ROP/history/", "ROPData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/ETN/history/", "ETNData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/PH/history/", "PHData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/AME/history/", "AMEData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/A/history/", "AData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/BRKR/history/", "BRKRData"})
	for a := 0; a < len(indexLookup); a++ {
		time.Sleep(10 * time.Second)
		response, err := http.Get(indexLookup[a][0])
		if err != nil {
			logger.Log.Println(err.Error())
		}
		defer response.Body.Close()

		dataInBytes, err := ioutil.ReadAll(response.Body)
		pageContent := string(dataInBytes)

		tblStartIndex := strings.Index(pageContent, "<table")
		if tblStartIndex == -1 {
			fmt.Println("No opening element found")
			os.Exit(0)
		}
		tblEndIndex := strings.Index(pageContent, "</tbody")
		if tblEndIndex == -1 {
			fmt.Println("No closing tag found.")
			os.Exit(0)
		}

		dataTbl := pageContent[tblStartIndex:tblEndIndex]
		sRows := strings.Split(dataTbl, "<tr")
		file, fileErr := os.Create("\\\\us01s-grs02\\reporting$\\MTUpload\\Management\\ExternalData\\" + indexLookup[a][1] + ".csv")
		if fileErr != nil {
			fmt.Println(fileErr)
			return
		}

		for i := 0; i < len(sRows); i++ {
			sCols := strings.Split(sRows[i], "<span")
			holdStr := ""
			for j := 0; j < len(sCols); j++ {
				if strings.Index(sCols[j], "<table") < 0 {
					colStartIndex := strings.Index(sCols[j], ">") + 1
					colEndIndex := strings.Index(sCols[j], "</span")
					if colEndIndex > 0 {
						colOutput := sCols[j][colStartIndex:colEndIndex]
						holdStr += colOutput + "|"
					}
				}
			}
			fmt.Fprintln(file, holdStr)
		}
	}
}
func getMTDIntraDayPrice() {
	currentTime := time.Now()
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	//timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(currentTime.In(location))
	//currentTime.Format("2006-01-02 15:04:05")

	/*
		check that the time is between 9:35 and 4:30
		at 9:30 wipe out old data (separate function)

		split up between different sites?

		url:="https://finance.yahoo.com/quote/MTD/"
		//"currentPrice":{"raw":759.97,
	*/
	//url := "https://finance.yahoo.com/quote/MTD/"
	//"currentPrice":{"raw":
}
func main() {
	sprxystring := mtproxy.Prxystring
	os.Setenv("HTTP_PROXY", sprxystring)
	exec.Command("install  ")
	scrapeSPRussell()
	gocron.Every(1).Day().At("10:32").Do(scrapeSPRussell)
	gocron.Every(1).Day().At("11:32").Do(scrapeSPRussell)
	gocron.Every(1).Day().At("12:32").Do(scrapeSPRussell)
	gocron.Every(1).Day().At("13:32").Do(scrapeSPRussell)
	gocron.Every(1).Day().At("14:32").Do(scrapeSPRussell)
	gocron.Every(1).Day().At("15:32").Do(scrapeSPRussell)
	gocron.Every(1).Day().At("16:32").Do(scrapeSPRussell)
	gocron.Every(1).Day().At("16:47").Do(scrapeSPRussell)
	gocron.Every(1).Day().At("17:02").Do(scrapeSPRussell)
	gocron.Every(1).Day().At("17:17").Do(scrapeSPRussell)
	<-gocron.Start()
	/*
		var indexLookup = [][]string{}
		indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/%5EGSPC/history/", "S&PData"})
		indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/%5ERUT/history/", "Russell2000Data"})
		indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/MTD/history/", "MTDData"})

		for a := 0; a < len(indexLookup); a++ {
			response, err := http.Get(indexLookup[a][0])
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()

			dataInBytes, err := ioutil.ReadAll(response.Body)
			pageContent := string(dataInBytes)

			tblStartIndex := strings.Index(pageContent, "<table")
			if tblStartIndex == -1 {
				fmt.Println("No opening element found")
				os.Exit(0)
			}
			tblEndIndex := strings.Index(pageContent, "</tbody")
			if tblEndIndex == -1 {
				fmt.Println("No closing tag found.")
				os.Exit(0)
			}

			dataTbl := pageContent[tblStartIndex:tblEndIndex]
			sRows := strings.Split(dataTbl, "<tr")
			file, fileErr := os.Create(indexLookup[a][1] + ".csv")
			if fileErr != nil {
				fmt.Println(fileErr)
				return
			}

			for i := 0; i < len(sRows); i++ {
				sCols := strings.Split(sRows[i], "<span")
				holdStr := ""
				for j := 0; j < len(sCols); j++ {
					if strings.Index(sCols[j], "<table") < 0 {
						colStartIndex := strings.Index(sCols[j], ">") + 1
						colEndIndex := strings.Index(sCols[j], "</span")
						if colEndIndex > 0 {
							colOutput := sCols[j][colStartIndex:colEndIndex]
							holdStr += colOutput + "|"
						}
					}
				}
				fmt.Fprintln(file, holdStr)
			}
		}
	*/
}
