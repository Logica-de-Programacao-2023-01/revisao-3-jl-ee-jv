package q1

//Dado um array de números inteiros "nums" e um número inteiro "target", retorne os índices dos dois números cuja soma
//seja igual a "target".
//
//Você pode assumir que cada entrada terá exatamente uma solução e não poderá usar o mesmo elemento duas vezes.
//
//Você pode retornar a resposta em qualquer ordem.

func TwoSum(nums []int, target int) []int {
	indices := make([]int, 2)
	for y := 0; y < len(nums); y++ {
		for x := y + 1; x < len(nums); x++ {
			if nums[x]+nums[y] == target {
				indices[0] = y
				indices[1] = x
			}
		}
	}
	return indices
}
