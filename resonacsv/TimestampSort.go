package resonacsv

import (
	"sort"
)

// TimestampSort タイムスタンプを基準に、古い順に並び替える。
func TimestampSort(inputSliceFile []FileSysInfo) []FileSysInfo {
	//	fmt.Println("TimestampComparison 関数 開始。")
	sort.Slice(inputSliceFile, func(i, j int) bool { return inputSliceFile[i].ModTime < inputSliceFile[j].ModTime })

	//	fmt.Println("TimestampComparison 関数 終了。")
	return inputSliceFile
}
