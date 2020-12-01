package resonacsv

// GrandmasterCSVFileGet 基準ファイルの全取得
func GrandmasterCSVFileGet(inputFile FileSysInfo) []CSVResonaData {
	var oldestInputFile []CSVResonaData

	// タイトル部分の読み込み。
	//	titleCSVFile := ReadTitleCSVFile(inputFile.Name)

	// 渡されたファイルのCSVファイル内容をすべて取得した。
	oldestInputFile = readFile(inputFile.FullPath+inputFile.Name, inputFile.fileLineCount)
	//	for ii, line := range oldestInputFile {
	//		fmt.Printf("oldestInputFileに追加：oldestInputFile[%d].HandlingYear/Month/Date %d %d %d, 金額(%d), 取引残高(%d)\n", ii, line.HandlingYear, line.HandlingMonth, line.HandlingDate, line.AmountMoney, line.PostTransMoney)
	//	}

	return oldestInputFile
}
