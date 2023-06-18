package q5

import "unicode"

//Uma frase é um palíndromo se, após converter todas as letras maiúsculas em letras minúsculas e remover todos os
//caracteres não alfanuméricos, ela for lida da mesma forma da esquerda para a direita e vice-versa. Caracteres
//alfanuméricos incluem letras e números.
//
//Dada uma string "s", retorne verdadeiro se for um palíndromo e falso caso contrário.

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	
	var cleaned []rune
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned = append(cleaned, char)
		}
	}

	length := len(cleaned)
	for i := 0; i < length/2; i++ {
		if cleaned[i] != cleaned[length-1-i] {
			return false
		}
	}

	return true
}







