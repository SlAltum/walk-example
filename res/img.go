package res

import (
	_ "embed"

	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/lxn/walk"
)

const (
	DPI int = 100
)

var (
	//go:embed img/favicon.png
	favicon     EmbedImage
	IconFavicon *walk.Icon
	//go:embed img/open.png
	open     EmbedImage
	IconOpen *walk.Icon
	//go:embed img/document-new.png
	documentNew     EmbedImage
	IconDocumentNew *walk.Icon
	//go:embed img/document-properties.png
	documentProperties     EmbedImage
	IconDocumentProperties *walk.Icon
	//go:embed img/system-shutdown.png
	systemShutdown     EmbedImage
	IconSystemShutdown *walk.Icon
)

func init() {
	if icon, err := favicon.ToIcon(); err == nil {
		IconFavicon = icon
	}
	if icon, err := open.ToIcon(); err == nil {
		IconOpen = icon
	}
	if icon, err := documentNew.ToIcon(); err == nil {
		IconDocumentNew = icon
	}
	if icon, err := documentProperties.ToIcon(); err == nil {
		IconDocumentProperties = icon
	}
	if icon, err := systemShutdown.ToIcon(); err == nil {
		IconSystemShutdown = icon
	}
}

type EmbedImage []byte

// ToBitmap
// 将byte数组转换为Bitmap对象，尽可能只调用一次，避免反复生成对象，造成不必要的内存资源消耗
//
//	@receiver b
//	@return *walk.Bitmap
//	@return error
func (b EmbedImage) ToBitmap() (*walk.Bitmap, error) {
	var img image.Image
	var err error
	img, _, err = image.Decode(bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	return walk.NewBitmapFromImageForDPI(img, DPI)
}

// ToIcon
// 将byte数组转换为Icon对象，尽可能只调用一次，避免反复生成对象，造成不必要的内存资源消耗
//
//	@receiver b
//	@return *walk.Icon
//	@return error
func (b EmbedImage) ToIcon() (*walk.Icon, error) {
	var img image.Image
	var err error
	img, _, err = image.Decode(bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	return walk.NewIconFromImageForDPI(img, DPI)
}
