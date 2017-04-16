/**
 * int型の片方向線形リスト
 *
 * [設計方針]
 *    データの型は構造体で表現する
 *      構造体のメンバはパッケージ利用者から不可視とする
 *        - 構造体の中身(リスト内のデータ，ポインタ)を直接
 *          書き換えられることを防ぐ
 *        - メンバ名の戦闘を小文字にすることで実現
 *    データ型の操作は，データ型のインスタンスをレシーバにする
 *    関数で実現する
 */

package linkedlist

import (
  "fmt"
)

/**
 * リストのデータ型定義
 * 内部メンバは非公開とし，リストの操作は
 * このパッケージが公開する関数を通じて行なう
 */
type IntLinkedList struct {
  head int
  next *IntLinkedList
}

/**
 * リストの末尾を表すNilの値.
 */
var intNil IntLinkedList = IntLinkedList { 
  //構造体の名前付き初期化書式 = <member_name>: <value>,
  head: 0,   // 構造体の初期化時にはカンマを忘れずに
  next: nil, // 構造体の初期化時にはカンマを忘れずに(末尾のメンバも同様)
}

/**
 * 新しいリストを返す
 */
func NewList() *IntLinkedList {
  return &intNil
}

/**
 * リストの先頭の値を返す
 */
func ( list IntLinkedList ) Head() int {
  return list.head
}

/**
 * 引数にとったリストの参照から先頭に値を追加し，新たなリストへの参照を返す
 *
 * 他の公開関数はポインタ経由でも操作できる*1 が，この関数はIntLinkedList
 * へのポインタを明示的に受け取る必要がある. その理由を以下の例で説明する.
 * (例) 引数listの指す先頭がintNilである場合
 *       listの型が IntLinkedListである時，この関数が呼ばれたタイミングでは，
 *       listのポインタ値が呼び出し元レシーバを指している
 *         -> intNilを指さなくなる
 *         -> リストの末尾を表現できない
 *       listの型が *IntLinkedListである時，この関数が呼ばれたタイミングでは，
 *       listのポインタ値はこのソースで定義しているintNilを指している
 *         -> intNilを指しているため，リストの末尾を表現できている
 */
func ( list *IntLinkedList ) Insert( value int ) *IntLinkedList {
  var newHead *IntLinkedList = new( IntLinkedList )
  newHead.head = value
  newHead.next = list
  return newHead
}

/**
 * 引数に取ったリストの先頭を除いた残りの要素からなるリストへの参照と，
 * 引数に取ったリストの先頭が "リストの末尾" であるかどうかを返す
 */
func ( list IntLinkedList ) Next() ( next *IntLinkedList, isTail bool ) {
  if (&list == &intNil) || (list.next == &intNil) || (list.next == nil) {
    next = &list
    isTail = true
  } else {
    next = list.next
    isTail = false
  }
  return
}
