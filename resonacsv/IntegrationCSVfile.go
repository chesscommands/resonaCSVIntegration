package resonacsv

import (
	"time"
)

// IntegrationCSVfile 比較後の結果を統合する。
func IntegrationCSVfile(baseFile []CSVResonaData, newFile []CSVResonaData) []CSVResonaData {
	var lastLine CSVResonaData

	if baseFile[len(baseFile)-1].RecordClass == "合計" {
		// 一つ目と二つ目の合計値を合算するための待避。
		lastLine = baseFile[len(baseFile)-1]
		baseFile = baseFile[:len(baseFile)-1]
		//	fmt.Println("IntegrationCSVfile(lastLine)：", lastLine)
		//	fmt.Println("IntegrationCSVfile(baseFile)：", baseFile)
	}

	for _, lineData := range newFile {
		//		fmt.Printf("baseFileに追加：lineData.HandlingYear/Month/Date %d %d %d, 金額(%d), 取引残高(%d)\n", lineData.HandlingYear, lineData.HandlingMonth, lineData.HandlingDate, lineData.AmountMoney, lineData.PostTransMoney)
		baseFile = append(baseFile, lineData)
	}

	// "合計"行の処理
	todayNow := time.Now()
	//lastLinePostTransMoney := lastLine.PostTransMoney + lineData.PostTransMoney
	lastLine.CsvGetYear = int64(todayNow.Year())
	lastLine.CsvGetMonth = int64(todayNow.Month())
	lastLine.CsvGetDay = int64(todayNow.Day())
	lastLine.CsvGetTime = int64(todayNow.Hour())
	lastLine.CsvGetMinutes = int64(todayNow.Minute())
	lastLine.ContactDoubleNAME = "りそなIBcsv統合Goプログラム"
	lastLine.FinancialName = "-"
	lastLine.BranchName = "-"
	lastLine.AccountClass = "-"
	lastLine.AccountType = "-"
	lastLine.AccountNumber = baseFile[len(baseFile)-1].AccountNumber //口座番号
	lastLine.ResendDisplay = "-"
	lastLine.HandlingYear = 9999
	lastLine.HandlingMonth = 99
	lastLine.HandlingDate = 99
	lastLine.PostTransMoney = baseFile[len(baseFile)-1].PostTransMoney // 取引後残高の変更
	baseFile = append(baseFile, lastLine)

	return baseFile
}
