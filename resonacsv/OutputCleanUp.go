package resonacsv

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// OutputCleanUp 1行目のタイトルが文字列と数値で異なりエラーになるため、ここで後始末(辻褄合わせ)をする。
func OutputCleanUp(inputFile FileSysInfo, csvFile []CSVResonaData) {
	fRead, errRead := os.Open(inputFile.Name)
	if errRead != nil {
		log.Fatal(errRead)
	}
	defer fRead.Close()

	// 書き込み用のファイルを開く(今回作成したファイル)。
	fWrite, errWrite := os.OpenFile(OutPutCSVFilename, os.O_RDWR, os.ModePerm)
	if errWrite != nil {
		log.Fatal("Error:", errWrite)
	}
	defer fWrite.Close()

	ioReadFile, errIoutil := ioutil.ReadAll(fWrite)
	if errIoutil != nil {
		log.Fatal("Error:", errIoutil)
	}

	fmt.Println("書き込み用のファイルを開く。")
	regBool := regexp.MustCompile(`.*(?m)?`)
	test := regBool.FindAllStringSubmatch(string(ioReadFile), -1)[0]
	fmt.Println(len(test), cap(test))
	//	fmt.Println(regBool.FindAllStringSubmatch(string(ioReadFile), -1)[0])
	fmt.Println("書き込み用のファイルを開き終わる。")

	// 1行目を読み取るためだけの処理
	lineRead := transform.NewReader(fRead, japanese.ShiftJIS.NewDecoder())
	readLineFile, errIoutil := ioutil.ReadAll(lineRead)
	if errIoutil != nil {
		log.Fatal("Error:", errIoutil)
	}
	fmt.Println("lineRead", string(readLineFile)[:246])

	// 1行目の入れ替え。
	//fWrite.Seek(0, 0)
	//	_, err := fmt.Fprint(fWrite, "テキスト書き込み") // レコード区分(int64)
	//titleDecode := string(readLineFile)[:247]
	//_, err := fmt.Fprint(fWrite, titleDecode)
	//if err != nil {
	//fmt.Fprintf(os.Stderr, "line.RecordClass,%s\n", err)
	//}

}
