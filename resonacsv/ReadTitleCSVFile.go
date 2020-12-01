package resonacsv

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// ReadTitleCSVFile 直近で入手したCSVファイルを指定する(読み込む)。
func ReadTitleCSVFile(inputFileName string) (inputCSVTitleData CSVResonaTitle) {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	//r.Comma = ',' // 標準

	// CSVファイルのタイトル部分を処理する。
	record, err := r.Read()
	if err == io.EOF {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	for ii, recordWord := range record {
		switch ii {
		case 0:
			// レコード区分(int64)
			inputCSVTitleData.RecordClass = recordWord
		case 1:
			// csv取得-年(int8)
			inputCSVTitleData.CsvGetYear = recordWord
		case 2:
			// csv取得-月(int8)
			inputCSVTitleData.CsvGetMonth = recordWord
		case 3:
			// csv取得-日(int8)
			inputCSVTitleData.CsvGetDay = recordWord
		case 4:
			// csv取得-時(int8)
			inputCSVTitleData.CsvGetTime = recordWord
		case 5:
			// csv取得-分(int8)
			inputCSVTitleData.CsvGetMinutes = recordWord
		case 6:
			inputCSVTitleData.ContactName = recordWord // 連絡先名(string)

		case 7:
			inputCSVTitleData.FinancialName = recordWord // 金融機関名(string)

		case 8:
			inputCSVTitleData.BranchName = recordWord // 支店名(string)

		case 9:
			inputCSVTitleData.AccountClass = recordWord // 口座番号区分(string)

		case 10:
			inputCSVTitleData.AccountType = recordWord // 口座種別(string)
		case 11:
			// 口座番号(int64)
			inputCSVTitleData.AccountNumber = recordWord

		case 12:
			inputCSVTitleData.ResendDisplay = recordWord // 再送表示(string)

		case 13:
			inputCSVTitleData.TransName = recordWord // 取引名(string)
		case 14:
			// 取扱日付　年(int8)
			inputCSVTitleData.HandlingYear = recordWord
		case 15:
			// 取扱日付　月(int8)
			inputCSVTitleData.HandlingMonth = recordWord
		case 16:
			// 取扱日付　日(int8)
			inputCSVTitleData.HandlingDate = recordWord
		case 17:
			// 金額(int64)
			inputCSVTitleData.AmountMoney = recordWord
		case 18:
			// 取引後残高(int64)
			inputCSVTitleData.PostTransMoney = recordWord

		case 19:
			inputCSVTitleData.Description = recordWord // 摘要(string)

		case 20:
			inputCSVTitleData.Comment = recordWord // コメント(string)
		}
	}

	return inputCSVTitleData
}
