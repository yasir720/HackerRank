package main

import "fmt"

func peePee(queries [][]int32) []int32 {
	// for _, query := range queries {
	// 	fmt.Println(query)
	// }

	mp := make(map[int32]int32)
	mpv := make(map[int32]int32)
	vs := make([]int32, 0)
	for _, v := range queries {
		switch v[0] {
		case 1:
			cur := mp[v[1]]
			mp[v[1]] = cur + 1
			if cur > 0 {
				mpv[cur] -= 1
			}
			mpv[cur+1] += 1
		case 2:
			cur := mp[v[1]]
			if cur > 0 {
				mp[v[1]] -= 1
				mpv[cur] -= 1
				mpv[cur-1] += 1
			}
		case 3:
			if x, ok := mpv[v[1]]; !ok || x == 0 {
				vs = append(vs, 0)
				fmt.Println("0")
			} else {
				vs = append(vs, 1)
				fmt.Println("1")
			}

		}
	}
	return vs

}

func pooPoo(n int32, s string) int64 {
	m := make(map[rune]int32)
	p := make([]rune, 0)
	var count int32 = 0

	for _, letter := range s {
		if _, ifLetter := m[letter]; !ifLetter {
			p = append(p, letter)
		}
		m[letter]++		
		//count++ // for condition 1
	}

	for i := 0; i< len(p); i++ {
			if m[p[i]] >= 2 {
				fmt.Println("did it")
				// for j := i; j < len(p); j++ {

				// }
				count = count + ((m[p[i]]/2) * int32(len(p) - 1))
			}
	}

	 fmt.Println(m)
	 fmt.Println(p)
	
	return 5
}

func main() {

	// b := [][]int32{}
	// row1 := []int32{1,1}
	// row2 := []int32{2,2}
	// row3 := []int32{3,2}
	// row4 := []int32{1,1}

	// b = append(b, row1)
	// b = append(b, row2)
	// b = append(b, row3)
	// b = append(b, row4)

	// p := peePee(b)
	// fmt.Println(p)

	pooPoo(4, "abbcccdddd")


}
