# go-identicon
This is a simple port of `npm i react-blockies` to a go-identicon.

```
ethereum_address := "0x76d46eeb21a7999883246c1ac162a4ea1447e7aa"
base64Blockie, _ := blockies.RenderIcon(blockies.Option{
			Seed:  strings.ToLower(ethereum_address),
			Size:  8, // width/height of the icon in blocks, default: 8
			Scale: 3, // width/height of each block in pixels, default: 4
		})
```
