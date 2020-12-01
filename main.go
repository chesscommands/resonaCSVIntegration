package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./resonacsv"
)

func main() {

	var fileinfo []resonacsv.FileSysInfo
	var criteriaCSVFile []resonacsv.CSVResonaData
	var theFirstCSVFile []resonacsv.CSVResonaData
	var padawanCSVFile []resonacsv.CSVResonaData
	var compareDone []resonacsv.CSVResonaData
	if len(os.Args) < 2 {
		log.Fatal("引数にりそなIB用のCSVファイルを渡してください。")
	}

	fmt.Println("りそなIB用CSVファイル読み取り開始。", time.Now().String())
	fileinfo = resonacsv.TimestampGet(os.Args[1:]) // ファイルタイムスタンプの取得
	fileinfo = resonacsv.TimestampSort(fileinfo)   // ファイルタイムスタンプ基準で並び替え。

	theFirstCSVFile = resonacsv.GrandmasterCSVFileGet(fileinfo[0])       // 基準ファイルのデータ全読み込み
	criteriaCSVFile = resonacsv.CriteriaFilePreparation(theFirstCSVFile) // 基準ファイルからの末尾最古行抜き出し。

	for ii := 1; ii < len(fileinfo); ii++ {
		padawanCSVFile = resonacsv.ApprenticeCSVFileGet(fileinfo[ii])                // 統合されるファイル
		compareDone = resonacsv.TimestampComparison(criteriaCSVFile, padawanCSVFile) // ファイル比較
		theFirstCSVFile = resonacsv.IntegrationCSVfile(theFirstCSVFile, compareDone) // ファイル比較後の統合

		criteriaCSVFile = resonacsv.CriteriaFilePreparation(theFirstCSVFile) // 基準ファイルからの末尾最古行抜き出し。
	}
	resonacsv.OutputTitleCSV(fileinfo[0])    // タイトルの出力
	resonacsv.OutputCSVfile(theFirstCSVFile) // CSVファイルへの出力
	//	resonacsv.OutputCleanUp(fileinfo[0], theFirstCSVFile) //タイトルを差し替える。

	fmt.Println("りそなIB用CSVファイル読み取り終了。", time.Now().String())
}
