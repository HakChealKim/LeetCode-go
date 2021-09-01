package main

func compress(chars []byte) int {
	newIndex, leftIndex := 0, 0

	for index, ch := range chars {
		if index == len(chars)-1 || ch != chars[index+1] {
			chars[newIndex] = ch
			newIndex++
			num := index - leftIndex + 1
			if num > 1 {
				anchor := newIndex
				for ; num > 0; num /= 10 {
					chars[newIndex] = '0' + byte(num%10)
					newIndex++
				}
				s := chars[anchor:newIndex]
				for i, n := 0, len(s); i < n/2; i++ {
					s[i], s[n-i-1] = s[n-i-1], s[i]
				}
			}
			leftIndex = index + 1
		}
	}
	return newIndex
}
