package canvas

import (
	"os"
	"path/filepath"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

func CharacterAnalyse(character string) string{
	const fontSize = 30.0
	const hFontToPixelRatio = 0.3
	const wFontToPixelRatio = 0.225
	fileName := "tmp/nf-icon-nf.png"
	// Create new canvas of dimension 100x100 mm
	c := canvas.New(fontSize * wFontToPixelRatio, fontSize * hFontToPixelRatio)

	// Create a canvas context used to keep drawing state
	ctx := canvas.NewContext(c)

	nerdFont := canvas.NewFontFamily("nf")
	if err := nerdFont.LoadFontFile("0xProtoNerdFontMono-Regular.ttf", canvas.FontRegular); err != nil {
			panic(err)
	}

	face := nerdFont.Face(fontSize, canvas.White, canvas.FontNormal, canvas.FontNormal)


	background := canvas.Rectangle(ctx.Width() + 1, ctx.Height() + 1);
	ctx.SetFillColor(canvas.Black)
	ctx.DrawPath(0, 0, background)
	ctx.DrawText(0, 1, canvas.NewTextLine(face, character, canvas.Left))

	// Rasterize the canvas and write to a PNG file with 3.2 dots-per-mm (320x320 px)
	_ = os.Mkdir(filepath.Join(".", "tmp"), os.ModePerm)
	if err := renderers.Write(fileName, c, canvas.DPMM(3.2)); err != nil {
		panic(err)
	}
	return fileName
}
