package resonacsv

import (
	"os"
	"time"
)

// OutPutCSVFilename 出力ファイル名。
const OutPutCSVFilename = "りそなIB統合済みcsvファイル.txt"

// Duration 時間を操作(保存)する。
type Duration int64

// FileSysInfo ファイル情報
type FileSysInfo struct {
	id         int         // 連番
	Name       string      // ファイル名
	FullPath   string      // Pathを含む。
	Size       int64       // ファイルサイズ
	Mode       os.FileMode // ファイル権限
	ModTime    string      // 修正日時(簡易版)
	ModifyTime time.Time   // 修正日時の正式版
	IsDir      bool        // ディレクトリか否か(ディレクトリであればtrue)
	Sys        string      // 基礎情報(取得がめんどくさいので未使用)

	fileLineCount int64 // ファイル行数
}

// CSVResonaData csvファイルの中身を保存
type CSVResonaData struct {
	id                int    // カラム番号(役に立たない。必ずゼロ)
	recordNo          int    // 行数(ファイル内の行数)
	RecordClass       string // レコード区分
	CsvGetYear        int64  // csv取得-年
	CsvGetMonth       int64  // csv取得-月
	CsvGetDay         int64  // csv取得-日
	CsvGetTime        int64  // csv取得-時
	CsvGetMinutes     int64  // csv取得-分
	ContactName       string // 連絡先名
	ContactDoubleNAME string // 連絡先名(全角)
	FinancialName     string // 金融機関名
	BranchName        string // 支店名
	AccountClass      string // 口座番号区分
	AccountType       string // 口座種別
	AccountNumber     int64  // 口座番号
	ResendDisplay     string // 再送表示
	TransName         string // 取引名
	HandlingYear      int64  // 取扱日付　年
	HandlingMonth     int64  // 取扱日付　月
	HandlingDate      int64  // 取扱日付　日
	AmountMoney       int64  // 金額
	PostTransMoney    int64  // 取引後残高
	Description       string // 摘要
	DescriptDouble    string // 摘要(全角)
	Comment           string // コメント
	CommentDouble     string // コメント(全角)
}

// CSVResonaTitle csvファイルのタイトル部分のみ保存
type CSVResonaTitle struct {
	RecordClass    string // レコード区分
	CsvGetYear     string // csv取得-年
	CsvGetMonth    string // csv取得-月
	CsvGetDay      string // csv取得-日
	CsvGetTime     string // csv取得-時
	CsvGetMinutes  string // csv取得-分
	ContactName    string // 連絡先名
	FinancialName  string // 金融機関名
	BranchName     string // 支店名
	AccountClass   string // 口座番号区分
	AccountType    string // 口座種別
	AccountNumber  string // 口座番号
	ResendDisplay  string // 再送表示
	TransName      string // 取引名
	HandlingYear   string // 取扱日付　年
	HandlingMonth  string // 取扱日付　月
	HandlingDate   string // 取扱日付　日
	AmountMoney    string // 金額
	PostTransMoney string // 取引後残高
	Description    string // 摘要
	Comment        string // コメント
}
