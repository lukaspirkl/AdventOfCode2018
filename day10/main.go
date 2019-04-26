package main

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"regexp"
	"strconv"
)

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

type light struct {
	position image.Point
	velocity image.Point
}

type lights []light

func (l lights) move(step int) lights {
	new := lights{}
	for _, i := range l {
		new = append(new, light{position: i.position.Add(i.velocity.Mul(step)), velocity: i.velocity})
	}
	return new
}

func (l lights) createImage() *image.RGBA {
	img := image.NewRGBA(l.getSize())
	draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)
	for _, light := range l {
		img.Set(light.position.X, light.position.Y, image.Black)
	}
	return img
}

func (l lights) getSize() image.Rectangle {
	minX := maxInt
	minY := maxInt
	maxX := minInt
	maxY := minInt
	for _, light := range l {
		if minX > light.position.X {
			minX = light.position.X
		}
		if maxX < light.position.X {
			maxX = light.position.X
		}
		if minY > light.position.Y {
			minY = light.position.Y
		}
		if maxY < light.position.Y {
			maxY = light.position.Y
		}
	}
	return image.Rectangle{image.Point{X: minX, Y: minY}, image.Point{X: maxX + 1, Y: maxY + 1}}
}

func main() {
	originalLights := getLightPositions()

	//lights.move(11030)
	min := maxInt
	minStep := 0
	var newLights lights
	for i := 0; i < 20000; i++ {
		newLights = originalLights.move(i)
		size := newLights.getSize().Size()
		s := size.X * size.Y
		if min > s {
			min = s
			minStep = i
		}
	}
	fmt.Println(minStep)

	f, _ := os.Create("image.png")
	png.Encode(f, originalLights.move(minStep).createImage())
}

func getLightPositions() lights {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()
	regex, err := regexp.Compile(`^position=< *(-?\d*), *(-?\d*)> velocity=< *(-?\d*), *(-?\d*)>$`)
	if err != nil {
		panic(err)
	}
	lights := []light{}
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		submatch := regex.FindStringSubmatch(scanner.Text())
		lights = append(lights, light{
			position: image.Point{X: atoi(submatch[1]), Y: atoi(submatch[2])},
			velocity: image.Point{X: atoi(submatch[3]), Y: atoi(submatch[4])},
		})
	}
	return lights
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
