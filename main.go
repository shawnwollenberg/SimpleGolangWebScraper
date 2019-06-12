package main

import (
	mtproxy "MTProxy"
	"os"
	"os/exec"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	sprxystring := mtproxy.Prxystring
	os.Setenv("HTTP_PROXY", sprxystring)
	exec.Command("install  ")
	
	var indexLookup = [][]string{}
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/%5EGSPC/history/", "S&PData"})
	indexLookup = append(indexLookup, []string{"https://finance.yahoo.com/quote/%5ERUT/history/", "Russell2000Data"})
	for a:=0; a<len(indexLookup); a++ {
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
	
}
