package opencvutils

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

func OpenCVFindCenterCoords(img, search image.Image) (int, int) {
	return OpenCVFindCenterCoordsWithDebug(img, search, false)
}

func OpenCVFindCenterCoordsWithDebug(img, search image.Image, debug bool) (int, int) {
	x, y, w, h, _ := OpenCVFindCoordsWithDebug(img, search, debug)
	return x + (w / 2), y + (h / 2)

}

func OpenCVFindCoordsWithDebug(img, search image.Image, debug bool) (int, int, int, int, float32) {
	mat, _ := gocv.ImageToMatRGB(img)
	tpl, _ := gocv.ImageToMatRGB(search)

	result := gocv.NewMat()

	gocv.MatchTemplate(mat, tpl, &result, gocv.TmCcorrNormed, gocv.NewMat())

	// newResult := gocv.NewMat()
	// gocv.Normalize(result, &newResult, 1, 0, gocv.NormMinMax)

	_, maxConfidence, _, maxLoc := gocv.MinMaxLoc(result)

	size := tpl.Size()

	if debug {
		gocv.Rectangle(&mat, image.Rect(
			maxLoc.X,
			maxLoc.Y,
			maxLoc.X+size[0],
			maxLoc.Y+size[1],
		), color.RGBA{R: 255, G: 0, B: 0, A: 1}, 3)

		gocv.IMWrite("./debug.png", mat)
	}

	return maxLoc.X, maxLoc.Y, size[0], size[1], maxConfidence
}
