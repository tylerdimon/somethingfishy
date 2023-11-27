package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tylerdimon/somethingfishy/assets"
	"image"
	"io/fs"
	"log"

	_ "image/png"
)

var PlayerSprite = mustLoadImage("fishTile_081.png")

var OppSprites = mustLoadImages("opps/*.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadImages(pattern string) []*ebiten.Image {
	var images []*ebiten.Image

	matches, err := fs.Glob(assets.Assets, pattern)
	if err != nil {
		log.Fatalf("failed to glob pattern: %v", err)
	}

	for _, match := range matches {
		img := mustLoadImage(match)
		images = append(images, img)
	}

	return images
}
