package resonacsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/text/width"
)

func readFile(inputFileName string, lineCount int64) []CSVResonaData {
	inputCSVFileData := make([]CSVResonaData, lineCount)
	//	fmt.Printf("Read file(%s)関数開始。\n", inputFileName)
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(inputFileName, "ファイルを開けない", err)
	}

	r := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	//r.Comma = ',' // 標準

	recordNo := 0 // ファイル行数
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for ii, csvRecord := range record {
			//			fmt.Printf("for文内のcsvRecord：%s\n", csvRecord)
			switch ii {
			case 0: // レコード区分(int64)
				inputCSVFileData[recordNo].id, _ = strconv.Atoi(csvRecord) // 連番(int)
				inputCSVFileData[recordNo].recordNo = recordNo             // ファイルの何行目かを保存する。
				_, err = fmt.Sscanf(csvRecord, "%s", &inputCSVFileData[recordNo].RecordClass)
				if err != nil {
					fmt.Fprintf(os.Stderr, "&inputCSVFileData[%d].RecordClass：%s,%s\n", recordNo, inputCSVFileData[recordNo].RecordClass, err)
				}
			case 1: // csv取得-年(int8)
				inputCSVFileData[recordNo].CsvGetYear, err = strconv.ParseInt(csvRecord, 10, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "readFile関数-%s\n", inputFileName)
					fmt.Fprintf(os.Stderr, "inputCSVFileData[%d].CsvGetYear：%d,err：%s\n", recordNo, inputCSVFileData[recordNo].CsvGetYear, err)
				}
			case 2: // csv取得-月(int8)
				inputCSVFileData[recordNo].CsvGetMonth, err = strconv.ParseInt(csvRecord, 10, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "inputCSVFileData[%d].CsvGetMonth,err[%s]\n", recordNo, err)
				}
			case 3: // csv取得-日(int8)
				inputCSVFileData[recordNo].CsvGetDay, err = strconv.ParseInt(csvRecord, 10, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "inputCSVFileData[%d].CsvGetDay,err[%s]\n", recordNo, err)
				}
			case 4: // csv取得-時(int8)
				inputCSVFileData[recordNo].CsvGetTime, err = strconv.ParseInt(csvRecord, 10, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "inputCSVFileData[%d].CsvGetTime,err[%s]\n", recordNo, err)
				}
			case 5: // csv取得-分(int8)
				inputCSVFileData[recordNo].CsvGetMinutes, err = strconv.ParseInt(csvRecord, 10, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "inputCSVFileData[%d].CsvGetMinutes,err[%s]\n", recordNo, err)
				}
			case 6:
				inputCSVFileData[recordNo].ContactName = csvRecord                          // 連絡先名(string)
				inputCSVFileData[recordNo].ContactDoubleNAME = width.Fold.String(csvRecord) // 連絡先名(全角)(string)
			case 7:
				inputCSVFileData[recordNo].FinancialName = csvRecord // 金融機関名(string)
			case 8:
				inputCSVFileData[recordNo].BranchName = csvRecord // 支店名(string)
			case 9:
				inputCSVFileData[recordNo].AccountClass = csvRecord // 口座番号区分(string)
			case 10:
				inputCSVFileData[recordNo].AccountType = csvRecord // 口座種別(string)
			case 11: // 口座番号(int64)
				inputCSVFileData[recordNo].AccountNumber, err = strconv.ParseInt(csvRecord, 10, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "inputCSVFileData[%d].AccountNumber,err[%s]\n", recordNo, err)
				}
			case 12:
				inputCSVFileData[recordNo].ResendDisplay = csvRecord // 再送表示(string)
			case 13:
				inputCSVFileData[recordNo].TransName = csvRecord // 取引名(string)
			case 14: // 取扱日付　年(int8)
				inputCSVFileData[recordNo].HandlingYear, err = strconv.ParseInt(csvRecord, 10, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "inputCSVFileData[%d].HandlingYear：%d,err[%s]\n", recordNo, inputCSVFileData[recordNo].HandlingYear, err)
				}
			case 15: // 取扱日付　月(int8)
				inputCSVFileData[recordNo].HandlingMonth, err = strconv.ParseInt(csvRecord, 10, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "inputCSVFileData[%d].HandlingMonth：%d,err[%s]\n", recordNo, inputCSVFileData[recordNo].HandlingMonth, err)
				}
			case 16: // 取扱日付　日(int8)
				inputCSVFileData[recordNo].HandlingDate, err = strconv.ParseInt(csvRecord, 10, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "inputCSVFileData[%d].HandlingDate：%d,err[%s]\n", recordNo, inputCSVFileData[recordNo].HandlingDate, err)
				}
			case 17: // 金額(int64)
				inputCSVFileData[recordNo].AmountMoney, err = strconv.ParseInt(csvRecord, 10, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "inputCSVFileData[%d].AmountMoney,err[%s]\n", recordNo, err)
				}
			case 18: // 取引後残高(int64)
				inputCSVFileData[recordNo].PostTransMoney, err = strconv.ParseInt(csvRecord, 10, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "inputCSVFileData[%d].PostTransMoney,err[%s]\n", recordNo, err)
				}
			case 19:
				inputCSVFileData[recordNo].Description = csvRecord                       // 摘要(string)
				inputCSVFileData[recordNo].DescriptDouble = width.Fold.String(csvRecord) // 摘要(全角)(string)
			case 20:
				inputCSVFileData[recordNo].Comment = csvRecord                          // コメント(string)
				inputCSVFileData[recordNo].CommentDouble = width.Fold.String(csvRecord) // コメント(string)
			}
		}
		recordNo++
	}

	//	fmt.Println("Read file関数終了。")
	return inputCSVFileData
}
