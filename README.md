# EbitenTextureUnpacker

Purpose of this library is to unpack sprite sheets packed using 
[Texture Packer](https://www.codeandweb.com/texturepacker).  
  
This implementation is for Texture Packer's JSON Array format. 

# Usage

First generate the packed PNG file and place the json and PNG file wherever you keep your assets. 

Add this package to your project.

```bash
go get github.com/Jodios/EbitenTextureUnpacker
```

Create an Unpacker with the path to your assets folder and 
call the Unpack function.
```golang
func init() {
	unpacker := &Unpacker{
		Filesystem: os.DirFS("./samples/assets"),
	}
	var err error
	sprites, err = unpacker.Unpack("sample_spritesheet.json")
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
