package resonacsv

import (
	"time"
)

// TimestampComparison 一つ目と二つ目のファイルを順次比較する(重複箇所の切り捨て後のデータを返却)。
func TimestampComparison(compareSource []CSVResonaData, inputNewFile []CSVResonaData) []CSVResonaData {
	const layout = "2006-01-02 15:04:05" // 年月のフォーマット
	const month30 = 30                   // 1ヶ月30日と仮定。

	// 引数：compareSource　⇒　統合元(比較元)のデータ(複数行あるが、基本は同年同月同日分)
	// 引数：inputFile　⇒　比較元から1つずつ新しいファイルが渡される。

	// 戻り値：compareDone　⇒　統合元との重複部分を切り捨てて、今回データを返却する。
	var sourceData CSVResonaData
	var compareDone []CSVResonaData

	var memoMessageSpan CSVResonaData
	var memoMessageDuplicate CSVResonaData
	memoMessageSpan.RecordClass = "注）2ヶ月以上取引が発生していない。（★注意★"
	memoMessageDuplicate.RecordClass = "注）重複あり。"

	// TransName(取引名)・Description(摘要)を最初に比較する。
	if len(compareSource) != 0 {
		//		sourceData = compareSource[len(compareSource)-1]
		sourceData = compareSource[0]
	}
	oneMessage := 1
	//for _, sourceData := range compareSource {
	for _, futureData := range inputNewFile {
		sourceDate := time.Date(int(sourceData.HandlingYear), time.Month(sourceData.HandlingMonth), int(sourceData.HandlingDate), 00, 00, 00, 0, time.Local)
		futureDate := time.Date(int(futureData.HandlingYear), time.Month(futureData.HandlingMonth), int(futureData.HandlingDate), 00, 00, 00, 0, time.Local)
		sourceDate.Format(layout)
		futureDate.Format(layout)
		dateDifference := futureDate.Sub(sourceDate) // 時間と分で差分が出る(for文の外に出すべき処理か？)。
		hours := int(dateDifference.Hours())
		hours = hours / 24 // 差分を日に変換
		//fmt.Println("hours(for文突入直後)：", hours)
		//fmt.Printf("比較元(%s)%s：year%d, month%d, day%d,時間%d,メッセージ%d\n", sourceData.TransName, sourceData.Description, sourceData.HandlingYear, sourceData.HandlingMonth, sourceData.HandlingDate, hours, oneMessage)
		//fmt.Printf("比較先(%s)%s：year%d, month%d, day%d\n", futureData.TransName, futureData.Description, futureData.HandlingYear, futureData.HandlingMonth, futureData.HandlingDate)
		if futureData.RecordClass == "レコード区分" || futureData.RecordClass == "合計" {
			//			fmt.Println("TimestampComparison関数：レコード区分と合計行をスキップ")
			continue
		} else if hours < 0 {
			//	fmt.Println("continue", hours)
			continue
		} else if hours >= 0 {
			//if hours == 0 && oneMessage == 1 {
			if hours == 0 {
				//				fmt.Printf("比較元(%s)%s：year%d, month%d, day%d,時間%d,メッセージ%d\n", sourceData.TransName, sourceData.Description, sourceData.HandlingYear, sourceData.HandlingMonth, sourceData.HandlingDate, hours, oneMessage)
				//				fmt.Printf("比較先(%s)%s：year%d, month%d, day%d\n", futureData.TransName, futureData.Description, futureData.HandlingYear, futureData.HandlingMonth, futureData.HandlingDate)
				for _, hoge := range compareSource {
					if hoge.TransName == futureData.TransName && hoge.Description == futureData.Description {
						//						fmt.Printf("比較元(%s)%s：year%d, month%d, day%d,時間%d,メッセージ%d\n", hoge.TransName, hoge.Description, hoge.HandlingYear, hoge.HandlingMonth, hoge.HandlingDate, hours, oneMessage)
						//						fmt.Printf("比較元(%s)%s：year%d, month%d, day%d,時間%d,メッセージ%d\n", sourceData.TransName, sourceData.Description, sourceData.HandlingYear, sourceData.HandlingMonth, sourceData.HandlingDate, hours, oneMessage)
						//						fmt.Printf("比較先(%s)%s：year%d, month%d, day%d\n", futureData.TransName, futureData.Description, futureData.HandlingYear, futureData.HandlingMonth, futureData.HandlingDate)
						oneMessage = 777
						continue
					}
				}
				//oneMessage = 777
				continue
			} else if hours > month30*2 && oneMessage == 1 {
				//fmt.Printf("比較元(%s)%s：year%d, month%d, day%d,時間%d,メッセージ%d\n", sourceData.TransName, sourceData.Description, sourceData.HandlingYear, sourceData.HandlingMonth, sourceData.HandlingDate, hours, oneMessage)
				//fmt.Printf("比較先(%s)%s：year%d, month%d, day%d\n", futureData.TransName, futureData.Description, futureData.HandlingYear, futureData.HandlingMonth, futureData.HandlingDate)
				oneMessage = 666
				compareDone = append(compareDone, memoMessageSpan) // 2ヶ月以上取引がない場合のメッセージ
				compareDone = append(compareDone, futureData)
			} else {
				//				fmt.Printf("無条件に取り込む：比較元(%s)%s：year%d, month%d, day%d,時間%d,メッセージ%d\n", sourceData.TransName, sourceData.Description, sourceData.HandlingYear, sourceData.HandlingMonth, sourceData.HandlingDate, hours, oneMessage)
				//				fmt.Printf("無条件に取り込む：比較先(%s)%s：year%d, month%d, day%d\n", futureData.TransName, futureData.Description, futureData.HandlingYear, futureData.HandlingMonth, futureData.HandlingDate)
				compareDone = append(compareDone, futureData)
				//	fmt.Println("無条件に取り込む")
			}
			//fmt.Printf("比較元(%s)%s：year%d, month%d, day%d,時間%d,メッセージ%d\n", sourceData.TransName, sourceData.Description, sourceData.HandlingYear, sourceData.HandlingMonth, sourceData.HandlingDate, hours, oneMessage)
			//fmt.Printf("比較先(%s)%s：year%d, month%d, day%d\n", futureData.TransName, futureData.Description, futureData.HandlingYear, futureData.HandlingMonth, futureData.HandlingDate)
		}
	}
	//}

	return compareDone
}
