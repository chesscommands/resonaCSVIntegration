package resonacsv

// ApprenticeCSVFileGet 新しめファイルの全取得
func ApprenticeCSVFileGet(inputFile FileSysInfo) []CSVResonaData {
	var padawanInputFile []CSVResonaData

	// 渡されたファイルのCSVファイル内容をすべて取得した。
	padawanInputFile = readFile(inputFile.FullPath+inputFile.Name, inputFile.fileLineCount)

	return padawanInputFile
}
