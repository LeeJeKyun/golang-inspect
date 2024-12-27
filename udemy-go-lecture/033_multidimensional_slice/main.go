package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnet"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
	for i := 0; i < len(xp); i++ {
		for j := 0; j < len(xp[i]); j++ {
			fmt.Printf("%s\t", xp[i][j])
			if j == len(xp[i])-1 {
				fmt.Print("\n")
			}
		}
	}

}
