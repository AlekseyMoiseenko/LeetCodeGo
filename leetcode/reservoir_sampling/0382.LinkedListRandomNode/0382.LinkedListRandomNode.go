package main

import (
	"math/rand"

	"github.com/AlekseyMoiseenko/LeetCodeGo/util"
)

type Solution struct {
	head *util.ListNode
}

func Constructor(head *util.ListNode) Solution {
	return Solution{
		head: head,
	}
}

// Алгоритм резервуарной выборки позволяет решить эту задачу за O(N) шагов и O(K) памяти. N - кол-во элементов, K - размер выборки (в задаче один элемент возвращаем)
// При этом не требуется знать N заранее, а условие случайности выборки ровно K элементов будет чётко соблюдено.
func (this *Solution) GetRandom() int {
	res := 0 // will be replaced with root.val since rand.Intn(1) always 0
	node := this.head
	i := 0
	for node != nil {
		i++
		// С вероятностью 1/N элемент номер N перезаписываем в качестве выбранного.
		// У каждого из предыдущих пришедших элементов остаётся тот же шанс 1/N остаться выбранным
		if rand.Intn(i) == 0 {
			res = node.Val
		}
		node = node.Next
	}
	return res
}
