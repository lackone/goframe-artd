package service

import (
	"backend/internal/consts"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
)

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

// Upload 上传文件
// @param file 上传文件
// @return 上传结果
// @error 上传错误
func (f *FileService) Upload(file *ghttp.UploadFile) (map[string]any, error) {
	rootDir := gfile.Pwd()
	yearMonth := gtime.Now().Format("Ym")
	day := gtime.Now().Format("d")
	
	// 构建保存目录
	saveDir := gfile.Join(rootDir, "resource", "public", consts.UploadDir, yearMonth, day)
	if !gfile.Exists(saveDir) {
		err := gfile.Mkdir(saveDir)
		if err != nil {
			return nil, err
		}
	}
	
	// 保存文件
	filename, err := file.Save(saveDir, true)
	if err != nil {
		return nil, err
	}

	filePath := "/" + consts.UploadDir + "/" + yearMonth + "/" + day + "/";
	
	// 构建返回结果
	result := map[string]any{
		"ext":  gfile.Ext(filename),
		"path": filePath + filename,
		"url":  consts.Domain + filePath + filename,
		"name": filename,
		"size": file.Size,
	}
	
	return result, nil
}
