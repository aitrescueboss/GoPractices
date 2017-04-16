package linkedlist //テスト対象パッケージと同じ名前にする

import (
  "fmt"
  "testing"
)

/**
 * 1つのリストを生成し，値の追加と参照を行なう 
 * Print文による確認のみ
 *   Print文(標準出力)の確認を行なうには以下のようにテストする
 *     $ go test -v
 */
func TestLinkedList_AddAndIteration(t *testing.T) {
  var list *IntLinkedList = NewList()
  fmt.Printf("Test(): first list addr = %p \n", list)
  list = list.Insert( 0 ) //listはポインタ型だが，レシーバのポインタ<->は暗黙的に変換される
  list = list.Insert( 1 )
  list = list.Insert( 2 )

  isTailOfList := false
  for l := list; !isTailOfList; l, isTailOfList = l.Next() {
    fmt.Println( l.Head() )
  }
}
