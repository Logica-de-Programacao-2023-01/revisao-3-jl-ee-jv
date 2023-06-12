package q1

//Dado um array de números inteiros "nums" e um número inteiro "target", retorne os índices dos dois números cuja soma
//seja igual a "target".
//
//Você pode assumir que cada entrada terá exatamente uma solução e não poderá usar o mesmo elemento duas vezes.
//
//Você pode retornar a resposta em qualquer ordem.

func TwoSum(nums []int, target int) []int {
	var indice []int
	for y:=0;y<len(nums);y++{
		for x:=len(nums);x<=0;x--{
			if nums[y] + nums[x] == target && x != y{
				indice = append(indice, x, y)
			}
		}
	}
	return indice
}
