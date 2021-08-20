package blockies

import (
	"testing"
)

func TestBlockies(t *testing.T) {
	x, _ := RenderIcon(Option{
		Seed:  "testings",
		Size:  8, // width/height of the icon in blocks, default: 8
		Scale: 3, // width/height of each block in pixels, default: 4
	})
	if x != "iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAIAAABvFaqvAAAAjklEQVR4nGJ5WKDHgAM07PqNLuLGiksxEy4JUgHVDGLE9BqmjzABph8HodcStTTRhODOxhNrmFKD0GuYsVYbtQrCYE4IRJP6u2A9hNG8LIxWLqJlrJEHBp/XWOAszMSGmaEwpeAig9Br8DTGAEtjxOQ1OIBrH4Reg+caPJGFCeCKmxnCqOwiqhkECAAA///ZMDIGmukhogAAAABJRU5ErkJggg==" {
		t.Error("not matched")
	}
}

func TestBlockies_FailedMandatoryParameter(t *testing.T) {
	_, err := RenderIcon(Option{
		Size:  8, // width/height of the icon in blocks, default: 8
		Scale: 3, // width/height of each block in pixels, default: 4
	})
	if err == nil {
		t.Error("mandatory parameter")
	}
}

func TestBlockies_WithDefaultScale(t *testing.T) {
	RenderIcon(Option{
		Seed: "0x01DC571BbD43a96c48658c9Fc5c35e7C0e90F9dd",
		Size: 8, // width/height of the icon in blocks, default: 8
	})
	// if x != "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAIAAAD8GO2jAAAApElEQVR4nGJ5WKDHQAg07PqNXdyNlaBeJoIqKARD3wJGXHGAK9xxAVzxMfSDiPZxkKiliVUCOUyJyQe41Az9IBq4fFAbtQrOZk4IxKrm74L1cHbzsjCsaoZ+EA1cPqAWGPpBRHMLWJA5uMoWXGU9LjXI4kM/iGgfB8jlCQNSeUJqfYAMkM0c+kFE+zhALseJSfu4ALLeZgaEmUM/iGhuASAAAP//xHEyFhVeEksAAAAASUVORK5CYII=" {
	// 	t.Error("not matched")
	// }
}
