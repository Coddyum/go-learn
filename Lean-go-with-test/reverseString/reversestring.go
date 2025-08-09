package reversestring


func reversed(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns) 
}


/*
 ceci est une autre option pour revers une chaine de character
 Mais cela est moins efficace prends plus de temps ( benchmark )

func reverse(str string) (result string) {
    for _, v := range str {
        result = string(v) + result
    }
    return
}
 */

 
func Name(name string) string {
    return reversed(name)
}
