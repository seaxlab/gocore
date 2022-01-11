package qrcode

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	qc "github.com/tuotoo/qrcode"
	"image/png"
	"os"
)

// spy 2020/1/27

// EncodeSimple 将内容编码成图片文件png
func EncodeSimple(content, filePath string) {
	qrCode, _ := qr.Encode(content, qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	file, _ := os.Create(filePath)
	defer file.Close()

	png.Encode(file, qrCode)
}

// Encode TODO 再扩展
// 高级API
func Encode(dto EncodeDTO) {
	qrCode, _ := qr.Encode(dto.Content, qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, dto.With, dto.Height)

	file, _ := os.Create(dto.FilePath)
	defer file.Close()

	png.Encode(file, qrCode)
}

// Decode 将二维码文件解码
func Decode(qrCodeFilePath string) (string, error) {
	file, err := os.Open(qrCodeFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	qrmatrix, err := qc.Decode(file)
	if err != nil {
		return "", err
	}

	return qrmatrix.Content, nil
}

type EncodeDTO struct {
	Content  string
	FilePath string
	With     int
	Height   int
}
