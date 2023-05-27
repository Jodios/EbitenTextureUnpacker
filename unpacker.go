package ebitentextureunpacker

import (
	"encoding/json"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

type TexturePackerJSONArray struct {
	Frames []struct {
		Filename string `json:"filename,omitempty"`
		Rotated  bool   `json:"rotated,omitempty"`
		Trimmed  bool   `json:"trimmed,omitempty"`
		Frame    struct {
			X int `json:"x,omitempty"`
			Y int `json:"y,omitempty"`
			W int `json:"w,omitempty"`
			H int `json:"h,omitempty"`
		} `json:"frame"`
		SpriteSourceSize struct {
			X int `json:"x,omitempty"`
			Y int `json:"y,omitempty"`
			W int `json:"w,omitempty"`
			H int `json:"h,omitempty"`
		} `json:"spriteSourceSize"`
		SourceSize struct {
			W int `json:"w,omitempty"`
			H int `json:"h,omitempty"`
		} `json:"sourceSize,omitempty"`
	} `json:"frames,omitempty"`
	Meta struct {
		App     string `json:"app,omitempty"`
		Version string `json:"version,omitempty"`
		Image   string `json:"image,omitempty"`
		Format  string `json:"format,omitempty"`
		Size    struct {
			W int `json:"w,omitempty"`
			H int `json:"h,omitempty"`
		} `json:"size,omitempty"`
		Scale       string `json:"scale,omitempty"`
		Smartupdate string `json:"smartupdate,omitempty"`
	}
}

type Unpacker struct {
	Filesystem fs.FS
}

func (unpacker *Unpacker) Unpack(path string) (map[string]*ebiten.Image, error) {
	sprites := make(map[string]*ebiten.Image)

	f, err := unpacker.Filesystem.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	texturePackerJSONArray := new(TexturePackerJSONArray)
	json.NewDecoder(f).Decode(texturePackerJSONArray)

	spriteSheet, err := unpacker.Filesystem.Open(texturePackerJSONArray.Meta.Image)
	if err != nil {
		return nil, err
	}
	defer spriteSheet.Close()

	spriteSheetImage, _, err := image.Decode(spriteSheet)
	if err != nil {
		return nil, err
	}

	ebitenSpriteSheetImage := ebiten.NewImageFromImage(spriteSheetImage)
	bounds := ebitenSpriteSheetImage.Bounds()
	for _, s := range texturePackerJSONArray.Frames {
		sprite := ebitenSpriteSheetImage.SubImage(image.Rectangle{
			Min: image.Point{
				X: bounds.Min.X + s.Frame.X,
				Y: bounds.Min.Y + s.Frame.Y,
			},
			Max: image.Point{
				X: bounds.Min.X + s.Frame.X + s.Frame.W,
				Y: bounds.Min.Y + s.Frame.Y + s.Frame.H,
			},
		})
		sprites[s.Filename] = ebiten.NewImageFromImage(sprite)
	}

	return sprites, nil

}
