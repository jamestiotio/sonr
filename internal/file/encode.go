package file

// func GetBase64Image(path string) (string, error) {
// 	imgBuffer := new(bytes.Buffer)
// 	// New File for ThumbNail
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	// Convert to Image Object
// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Encode as Jpeg into buffer
// 	err = jpeg.Encode(imgBuffer, img, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	b64 := base64.StdEncoding.EncodeToString(imgBuffer.Bytes()), nil
// }