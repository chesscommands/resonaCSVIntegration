package resonacsv

import (
	"log"
	"os"
	"path/filepath"
)

// TimestampGet 引数で渡された複数のファイル情報を取得する。
func TimestampGet(inputSliceFile []string) []FileSysInfo {
	const timeLayout = "20060102150405" // YYYYMMDDhhmmss
	inputFilesInfo := make([]FileSysInfo, len(inputSliceFile))
	//	fmt.Println("FileTimestampComparison関数開始。")
	for ii, fileName := range inputSliceFile {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if fi, err := file.Stat(); err == nil {
			inputFilesInfo[ii].id = ii // 外部からの参照不可だが、連番格納
			inputFilesInfo[ii].Name = fi.Name()
			fullPathOnly := filepath.Dir(inputSliceFile[ii])
			fullPahtName, err := filepath.Abs(fullPathOnly)
			if err != nil {
				log.Fatal(err)
			}
			inputFilesInfo[ii].FullPath = filepath.FromSlash(fullPahtName + "/")
			inputFilesInfo[ii].Size = fi.Size()
			inputFilesInfo[ii].Mode = fi.Mode()
			inputFilesInfo[ii].ModTime = fi.ModTime().Format(timeLayout)
			inputFilesInfo[ii].ModifyTime = fi.ModTime()
			inputFilesInfo[ii].IsDir = fi.IsDir()

			fileCount, err := lineCounter(fi.Size(), file) // ファイル行取得
			if err != nil {
				log.Fatal(err)
			}
			inputFilesInfo[ii].fileLineCount = fileCount
		}
	}

	// 拡張子がcsv以外のファイルをはじく(順序を気にしていたが、この後並び替える)。
	for ii, filename := range inputFilesInfo {
		extCSV := filepath.Ext(filename.Name)
		if extCSV != ".csv" {
			inputFilesInfo = append(inputFilesInfo[:ii], inputFilesInfo[ii+1:]...)
		}
	}

	//	fmt.Println("FileTimestampComparison関数終了。")
	return inputFilesInfo
}
