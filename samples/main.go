package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jodios/ebitentextureunpacker"
)

var sprites map[string]ebitentextureunpacker.ImageWithFrameDetails
var spriteNames []string

func init() {
	unpacker := &ebitentextureunpacker.Unpacker{
		Filesystem: os.DirFS("./samples/assets"),
	}
	var err error
	sprites, err = unpacker.UnpackWithFrameDetails("sample_spritesheet.json")
	if err != nil {
		log.Fatal(err)
	}
	spriteNames = make([]string, 0)
	for n := range sprites {
		spriteNames = append(spriteNames, n)
	}
	fmt.Printf("%d assets loaded\n", len(sprites))
}

type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	return nil
}

/*
*
draw function is called once per frame
*
*/
func (g *Game) Draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
	var y float64 = 0
	for i := 0; i < len(spriteNames); i++ {
		opts.GeoM.Translate(29, 0)
		if i%4 == 0 {
			opts.GeoM.Reset()
			opts.GeoM.Translate(0, y)
			y += 16
		}
		screen.DrawImage(sprites[spriteNames[i]].Image, opts)
	}

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 100, 100
}

// This is just a sample 😊
func main() {
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Just a test :)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
