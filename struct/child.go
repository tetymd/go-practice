package main

import (
    "fmt"
)

type BattleObject struct {
    name string
    hp int
    mp int
}

// 構造体の埋め込み
type Monster struct {
    BattleObject
    kind string
}

func main() {
    var slime Monster

    slime.name = "ザコスラ"
    slime.hp = 100
    slime.mp = 10
    slime.kind = "スライム"

    fmt.Println(slime.name)
    fmt.Println(slime.hp)
    fmt.Println(slime.mp)
    fmt.Println(slime.kind)

    //ワンライナーで初期化
    var dragon Monster = Monster{BattleObject{"drgn", 200, 100}, "dragon"}

    fmt.Println(dragon.name)
    fmt.Println(dragon.hp)
    fmt.Println(dragon.mp)
    fmt.Println(dragon.kind)
}
