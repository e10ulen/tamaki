# これなに  
玉置に基本情報技術者試験解法を解くためのリポジトリ  

## No1
携帯端末上で稼働する簡易メモ帳の機能のうち、メモの編集処理（メモの追加・削除・変更・移動）  
を行う部分のプログラムである。
図１は、簡易メモ帳に４件のメモ "Aoki", "Imai", "Uno", 及び "Endo" を登録した。  

問題は [no1.go](no1.go) に記載。
解答は [no1answer.go](no1answer.go)


## 設問１  
a, b に関する解答群
ア: DataLen  
イ: DataLen + 1  
ウ: DataLen + textLen  
エ: DataLen + textLen + 1  
オ: textLen  
カ: textLen + 1  

cに関する解答群  
ア: MemoCnt - pos  
イ: pos - 1  
ウ: pos  
エ: pos + 1  

dに関する解答群  
ア: fromPos, i >= toPos - 1, -1  
イ: fromPos, i >= toPos + 1. -1  
ウ: toPos, i <= fromPos - 1, 1  
エ: toPos, i <= fromPos + 1, 1  

### e10ulenなりの答え。  
オカアイ  
### 参考書上の答え。  
アウエイ  

## 解法解説  
この欄は後日説明
