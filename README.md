[![Go Reference](https://pkg.go.dev/badge/github.com/jodios/ebitentextureunpacker.svg)](https://pkg.go.dev/github.com/jodios/ebitentextureunpacker)
# EbitenTextureUnpacker

Purpose of this library is to unpack sprite sheets packed using 
[Texture Packer](https://www.codeandweb.com/texturepacker).  
  
This implementation is for Texture Packer's JSON Array format. 

# Usage

First generate the packed PNG file and place the json and PNG file wherever you keep your assets. 

Add this package to your project.

```bash
go get github.com/jodios/ebitentextureunpacker
```

Create an embed file to embed all of your assets into the actual  executable
of your game. 
```go
package assets
import (
	_ "embed"
)
var (
	//go:embed sample_spritesheet.json
	Spritesheet_JSON []byte
	//go:embed sample_spritesheet.png
	Spritesheet_PNG []byte
)
```  
  

Create an Unpacker with the path to your assets folder and 
call the Unpack function.
```go
func init() {
	unpacker := &Unpacker{}
	var err error
	sprites, err = unpacker.Unpack(assets.Spritesheet_JSON, assets.Spritesheet_PNG)
	if err != nil {
		log.Fatal(err)
	}
}
```
Unpack returns a map of image names to images so you can access the sprite you need by the original file name.   
Example...
```golang
downIdle := sprites["orange_down_idle.png"]
```