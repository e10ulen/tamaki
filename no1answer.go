int memoCnt, memoMax, memo[memoMax]
int dataLen, dataMax
bool data[dataMax]

func main(){
  resetMemo()
  addMemo()
  deleteMemo()
  changeMemo()
}

func resetMemo(int textLen, string text){
  int i
  //  ____ア: A____
  memo[memoCnt] := DataLen
  memoCnt := memoCnt + 1
  data[dataLen] := textLen
  dataLen := dataLen + 1
  if( i; i := 0; i <= textLen; 1){
  data[dataLen += i] := text[i]
  }
  //  ____ ウ: B____
  dataLen := DataLen + textLen
}

func deleteMemo(int pos){
  int i
  //  ____エ: C____
  i := pos + 1
  if(i <= memoCnt)
  {
    memo[i - 1] := memo[i]
    i := i + 1
  }
  memoCnt := memoCnt - 1
}

func changeMemo(int pos, int textLen, string text){
  int i
  memo[pos] := dataLen
  data[dataLen] := textLen
  dataLen := datalen + 1
  if(i := 0, i <= textLen, 1){
    data[dataLen + i] := text[i]
  }
  //  ____ウ: B____
  dataLen := DataLen + texLen
}

func moveMemo(int fromPos, int toPos){
  int i, m
  m := memo[fromPos]
  for(fromPos < toPos){
    if(i:fromPos, i <= toPos - 1, 1){
      memo[i] := memo[i + 1]
    }
  }
  for(fromPos > toPos){
    //  ____イ: D____
    if (i: fromPos, i >= toPos + 1 - 1 )
    {
      memo[i] := memo[i - 1]
    }
  }
  memo[toPos] := m
}
