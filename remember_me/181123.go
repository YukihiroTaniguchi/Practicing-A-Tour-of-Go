// http://127.0.0.1:3999/moretypes/
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for y := range img {
		img[y] = make([]uint8, dx)
	}

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			img[y][x] = uint8((x + y) / 2)
		}
	}

	return img

}

func main() {
	pic.Show(Pic)
}
