package resonacsv

import (
	"fmt"
	"log"
	"os"
)

// OutputCSVfile 統合済みのCSVデータをファイルに出力。
func OutputCSVfile(csvFile []CSVResonaData) {
	f, err := os.OpenFile(OutPutCSVFilename, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSVファイルのタイトル部分を削除する。
	csvFile = csvFile[1:]

	for _, line := range csvFile {

		if line.RecordClass[0:3] == "注" {
			_, err := fmt.Fprint(f, line.RecordClass+"\n") // レコード区分(int64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "line.RecordClass(メッセージ版),%s\n", err)
			}
			continue // メッセージの場合は、これ以降の処理をスキップする。
		} else {
			_, err := fmt.Fprint(f, line.RecordClass) // レコード区分(int64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "line.RecordClass,%s\n", err)
			}
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.RecordClass後の読点,%s\n", err)
		}

		_, err = fmt.Fprintf(f, "%04d", line.CsvGetYear) // csv取得-年(int8)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.CsvGetYear,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.CsvGetYear後の読点,%s\n", err)
		}

		_, err = fmt.Fprintf(f, "%02d", line.CsvGetMonth) // csv取得-月(int8)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.CsvGetMonth,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.CsvGetMonth後の読点,%s\n", err)
		}

		_, err = fmt.Fprintf(f, "%02d", line.CsvGetDay) // csv取得-日(int8)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.CsvGetDay,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.CsvGetDay後の読点,%s\n", err)
		}

		_, err = fmt.Fprintf(f, "%02d", line.CsvGetTime) // csv取得-時(int8)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.CsvGetTime,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.CsvGetTime後の読点,%s\n", err)
		}

		_, err = fmt.Fprintf(f, "%02d", line.CsvGetMinutes) // csv取得-分(int8)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.CsvGetMinutes,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.CsvGetMinutes後の読点,%s\n", err)
		}

		_, err = fmt.Fprint(f, line.ContactDoubleNAME) // 連絡先名(全角)(string)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.ContactDoubleNAME,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.ContactDoubleNAME後の読点,%s\n", err)
		}

		_, err = fmt.Fprint(f, line.FinancialName) // 金融機関名(string)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.FinancialName,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.FinancialName後の読点,%s\n", err)
		}

		_, err = fmt.Fprint(f, line.BranchName) // 支店名(string)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.BranchName,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.BranchName後の読点,%s\n", err)
		}

		_, err = fmt.Fprint(f, line.AccountClass) // 口座番号区分(string)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.AccountClass,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.AccountClass後の読点,%s\n", err)
		}

		_, err = fmt.Fprint(f, line.AccountType) // 口座種別(string)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.AccountType,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.AccountType後の読点,%s\n", err)
		}

		_, err = fmt.Fprint(f, line.AccountNumber) // 口座番号(int64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.AccountNumber,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.AccountNumber後の読点,%s\n", err)
		}

		_, err = fmt.Fprint(f, line.ResendDisplay) // 再送表示(string)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.ResendDisplay,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.ResendDisplay後の読点,%s\n", err)
		}

		_, err = fmt.Fprint(f, line.TransName) // 取引名(string)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.TransName,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.TransName後の読点,%s\n", err)
		}

		_, err = fmt.Fprintf(f, "%04d", line.HandlingYear) // 取扱日付　年(int8)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.HandlingYear,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.HandlingYear後の読点,%s\n", err)
		}

		_, err = fmt.Fprintf(f, "%02d", line.HandlingMonth) // 取扱日付　月(int8)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.HandlingMonth,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.HandlingMonth後の読点,%s\n", err)
		}

		_, err = fmt.Fprintf(f, "%02d", line.HandlingDate) // 取扱日付　日(int8)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.HandlingDate,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.HandlingDate後の読点,%s\n", err)
		}

		_, err = fmt.Fprint(f, line.AmountMoney) // 金額(int64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.AmountMoney,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.AmountMoney後の読点,%s\n", err)
		}

		_, err = fmt.Fprint(f, line.PostTransMoney) // 取引後残高(int64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.PostTransMoney,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.PostTransMoney後の読点,%s\n", err)
		}

		_, err = fmt.Fprint(f, line.DescriptDouble) // 摘要(全角)(string)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.DescriptDouble,%s\n", err)
		}
		_, err = fmt.Fprint(f, ",")
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.DescriptDouble後の読点,%s\n", err)
		}

		_, err = fmt.Fprint(f, line.CommentDouble+"\n") // コメント(string)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line.Comment,%s\n", err)
		}
	}
}
