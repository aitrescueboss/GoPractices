# LinkedList

int型の値を格納できる単方向リストの実装です．

# 例

```
  var list *IntLinkedList = NewList() //リスト生成
  list = list.Insert( 0 ) // リストへの値追加
  list = list.Insert( 1 )
  list = list.Insert( 2 )

  // リストのイテレーション
  var isTailOfList bool = false
  for l := list; !isTailOfList; l, isTailOfList = l.Next() {
    fmt.Println( l.Head() ) // リストの先頭が指す値の取得
  }
```