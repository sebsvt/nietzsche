package image

type AddWatermarkParameters struct {
	Elements []WatermarkElement
}

type WatermarkElement struct {
	Type                           string
	Text                           string
	Image                          string
	Gravity                        string
	VerticalAdjustmentPercentage   int
	HorizontalAdjustmentPercentage int
	Rotation                       int
	FontFamily                     string
	FontStyle                      string
	FontSize                       int
	FontColor                      string
	Transparency                   int
	Mosaic                         bool
}

// func AddWatermark(params *AddWatermarkParameters) (string, error) {
// 	img, err := imaging.Open(params.InPath)
// 	if err != nil {
// 		return "", ErrCouldNotReadFile
// 	}
// 	overlay, err := imaging.Open(params.Image)
// 	if err != nil {
// 		return "", ErrCouldNotReadFile
// 	}
// 	// add watermark
// 	// reduce transparency to smaller scale
// 	transparency := float64(params.Transparency) / 100
// 	image := imaging.Overlay(img, overlay, image.Pt(50, 50), transparency)
// 	err = imaging.Save(image, params.OutPath)
// 	if err != nil {
// 		return "", ErrCouldNotSaveFile
// 	}

// 	return params.OutPath, nil
// }
