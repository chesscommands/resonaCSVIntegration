package resonacsv

// CriteriaFilePreparation 基準ファイルでの末尾行から同年月日での最古行を抜き出す(ファイル読み込み含む)。
func CriteriaFilePreparation(inputFile []CSVResonaData) []CSVResonaData {
	var minimumInputFile []CSVResonaData

	maxCSVfileCount := len(inputFile)
	maxCSVfileCount-- // 最終行にするための調整デクリメント
	maxCSVfileCount-- // 合計行を除く。
	for ii := maxCSVfileCount; ii > 1; ii-- {
		if inputFile[ii-1].HandlingDate == inputFile[ii].HandlingDate {
			if inputFile[ii-1].HandlingMonth == inputFile[ii].HandlingMonth {
				if inputFile[ii-1].HandlingYear == inputFile[ii].HandlingYear {
					//			fmt.Println("1つ古いデータを取得した。")
					//			fmt.Println("inputFile[ii-1].HandlingYear⇒", inputFile[ii-1].HandlingYear)
					//			fmt.Println("inputFile[ii].HandlingYear⇒", inputFile[ii].HandlingYear)
					//			fmt.Println("inputFile[ii-1].HandlingYear/Month/Date", inputFile[ii-1].HandlingYear, inputFile[ii-1].HandlingMonth, inputFile[ii-1].HandlingDate)
					//					fmt.Printf("年月日完全一致：%d,inputFile[ii].HandlingYear/Month/Date %d %d %d, 金額(%d), 取引残高(%d)\n", ii, inputFile[ii].HandlingYear, inputFile[ii].HandlingMonth, inputFile[ii].HandlingDate, inputFile[ii].AmountMoney, inputFile[ii].PostTransMoney)
					minimumInputFile = append(minimumInputFile, inputFile[ii])
				} else {
					//					fmt.Printf("月日一致：%d,inputFile[ii].HandlingYear/Month/Date %d %d %d, 金額(%d), 取引残高(%d)\n", ii, inputFile[ii].HandlingYear, inputFile[ii].HandlingMonth, inputFile[ii].HandlingDate, inputFile[ii].AmountMoney, inputFile[ii].PostTransMoney)
					//minimumInputFile = append(minimumInputFile, inputFile[ii])
					break
				}
			} else {
				//				fmt.Printf("日のみ一致：%d,inputFile[ii].HandlingYear/Month/Date %d %d %d, 金額(%d), 取引残高(%d)\n", ii, inputFile[ii].HandlingYear, inputFile[ii].HandlingMonth, inputFile[ii].HandlingDate, inputFile[ii].AmountMoney, inputFile[ii].PostTransMoney)
				//minimumInputFile = append(minimumInputFile, inputFile[ii])
				break
			}
		} else {
			//			fmt.Printf("何者にも一致せず：%d,inputFile[ii].HandlingYear/Month/Date %d %d %d, 金額(%d), 取引残高(%d)\n", ii, inputFile[ii].HandlingYear, inputFile[ii].HandlingMonth, inputFile[ii].HandlingDate, inputFile[ii].AmountMoney, inputFile[ii].PostTransMoney)
			minimumInputFile = append(minimumInputFile, inputFile[ii])
			break
		}

	}

	return minimumInputFile
}
