package resonacsv

import (
	"fmt"
	"log"
	"os"
)

// OutputTitleCSV 統合済みのCSVデータをファイルに出力。
func OutputTitleCSV(inputFile FileSysInfo) {
	if f, err := os.Stat(OutPutCSVFilename); os.IsNotExist(err) || f.IsDir() {
		// ファイルがない。
	} else {
		// ファイルの削除
		if err := os.Remove(OutPutCSVFilename); err != nil {
			log.Fatal(err)
		}
	}
	f, err := os.OpenFile(OutPutCSVFilename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// タイトル部分の読み込み。
	titleCSVFile := ReadTitleCSVFile(inputFile.FullPath + inputFile.Name)

	_, err = fmt.Fprint(f, titleCSVFile.RecordClass+",") // レコード区分(int64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.RecordClass,%s\n", err)
	}
	_, err = fmt.Fprint(f, "CSV取得-"+titleCSVFile.CsvGetYear) // csv取得-年(int8)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetYear,%s\n", err)
	}
	_, err = fmt.Fprint(f, ",")
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetYear,%s\n", err)
	}
	_, err = fmt.Fprint(f, titleCSVFile.CsvGetMonth) // csv取得-月(int8)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetMonth,%s\n", err)
	}
	_, err = fmt.Fprint(f, ",")
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetYear,%s\n", err)
	}
	_, err = fmt.Fprint(f, titleCSVFile.CsvGetDay) // csv取得-日(int8)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetDay,%s\n", err)
	}
	_, err = fmt.Fprint(f, ",")
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetYear,%s\n", err)
	}
	_, err = fmt.Fprint(f, titleCSVFile.CsvGetTime) // csv取得-時(int8)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetTime,%s\n", err)
	}
	_, err = fmt.Fprint(f, ",")
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetYear,%s\n", err)
	}
	_, err = fmt.Fprint(f, titleCSVFile.CsvGetMinutes) // csv取得-分(int8)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetMinutes,%s\n", err)
	}
	_, err = fmt.Fprint(f, ",")
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetYear,%s\n", err)
	}
	if titleCSVFile.ContactName == "" {
		titleCSVFile.ContactName = ","
	}
	_, err = fmt.Fprint(f, titleCSVFile.ContactName+",") // 連絡先名(全角)(string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.ContactDoubleNAME,%s\n", err)
	}
	if titleCSVFile.FinancialName == "" {
		titleCSVFile.FinancialName = ","
	}
	_, err = fmt.Fprint(f, titleCSVFile.FinancialName+",") // 金融機関名(string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.FinancialName,%s\n", err)
	}
	if titleCSVFile.BranchName == "" {
		titleCSVFile.BranchName = ","
	}
	_, err = fmt.Fprint(f, titleCSVFile.BranchName+",") // 支店名(string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.BranchName,%s\n", err)
	}
	if titleCSVFile.AccountClass == "" {
		titleCSVFile.AccountClass = ","
	}
	_, err = fmt.Fprint(f, titleCSVFile.AccountClass+",") // 口座番号区分(string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.AccountClass,%s\n", err)
	}
	if titleCSVFile.AccountType == "" {
		titleCSVFile.AccountType = ","
	}
	_, err = fmt.Fprint(f, titleCSVFile.AccountType+",") // 口座種別(string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.AccountType,%s\n", err)
	}
	_, err = fmt.Fprint(f, titleCSVFile.AccountNumber) // 口座番号(int64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.AccountNumber,%s\n", err)
	}
	if titleCSVFile.ResendDisplay == "" {
		titleCSVFile.ResendDisplay = ","
	}
	_, err = fmt.Fprint(f, titleCSVFile.ResendDisplay+",") // 再送表示(string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.ResendDisplay,%s\n", err)
	}
	if titleCSVFile.TransName == "" {
		titleCSVFile.TransName = ","
	}
	_, err = fmt.Fprint(f, titleCSVFile.TransName+",") // 取引名(string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.TransName,%s\n", err)
	}
	_, err = fmt.Fprint(f, titleCSVFile.HandlingYear) // 取扱日付　年(int8)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.HandlingYear,%s\n", err)
	}
	_, err = fmt.Fprint(f, ",")
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetYear,%s\n", err)
	}
	_, err = fmt.Fprint(f, titleCSVFile.HandlingMonth) // 取扱日付　月(int8)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.HandlingMonth,%s\n", err)
	}
	_, err = fmt.Fprint(f, ",")
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetYear,%s\n", err)
	}
	_, err = fmt.Fprint(f, titleCSVFile.HandlingDate) // 取扱日付　日(int8)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.HandlingDate,%s\n", err)
	}
	_, err = fmt.Fprint(f, ",")
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetYear,%s\n", err)
	}
	_, err = fmt.Fprint(f, titleCSVFile.AmountMoney) // 金額(int64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.AmountMoney,%s\n", err)
	}
	_, err = fmt.Fprint(f, ",")
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetYear,%s\n", err)
	}
	_, err = fmt.Fprint(f, titleCSVFile.PostTransMoney) // 取引後残高(int64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.PostTransMoney,%s\n", err)
	}
	_, err = fmt.Fprint(f, ",")
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.CsvGetYear,%s\n", err)
	}
	if titleCSVFile.Description == "" {
		titleCSVFile.Description = ","
	}
	_, err = fmt.Fprint(f, titleCSVFile.Description+",") // 摘要(全角)(string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.DescriptDouble,%s\n", err)
	}
	_, err = fmt.Fprint(f, titleCSVFile.Comment+"\n") // コメント(string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "titleCSVFile.Comment,%s\n", err)
	}
}
