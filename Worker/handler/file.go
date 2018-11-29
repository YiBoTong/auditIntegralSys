package handler

import (
	"auditIntegralSys/Worker/db/file"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"errors"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/os/gfile"
	"gitee.com/johng/gf/g/os/gtime"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

type File struct {
	gmvc.Controller
}

func (f *File) Upload() {
	fileId := 0
	// 年月作为文件的存放文件夹
	folder := gtime.Now().Format("Ym") + "/"
	msg := ""
	name := ""
	fileName := ""
	fl, h, err := f.Request.FormFile("file")
	if err == nil && h.Size > 1024*1024*1024 {
		msg = "文件过大，不允许上传"
		err = errors.New(msg)
	}
	if err == nil {
		defer fl.Close()
		name = gfile.Basename(h.Filename)
		fileName = gconv.String(gtime.Second()) + "_" + name
		buffer := make([]byte, h.Size)
		_, _ = fl.Read(buffer)
		err = gfile.PutBinContents(g.Config().GetString("filePath")+folder+fileName, buffer)
	}
	log.Instance().Infofln("Upload:header:%v,size:%v,folder:%v", h.Header, h.Size, folder)
	if err == nil {
		fileNameArr := strings.Split(fileName, ".")
		// 文件后缀
		fileSuffix := fileNameArr[len(fileNameArr)-1]
		fileId, err = db_file.AddFile(g.Map{
			"name":      name[:len(name)-len(fileSuffix)-1],
			"suffix":    fileSuffix,
			"path":      folder,
			"size":      h.Size,
			"file_name": fileName[:len(fileName)-len(fileSuffix)-1],
			"time":      util.GetLocalNowTimeStr(),
		})
	}
	if err != nil {
		// 失败就在数据库中删除记录
		_, _ = db_file.UpdateFile(fileId, g.Map{"delete": 1})
		fileId = 0
		log.Instance().Errorfln("[File Add]: %v", err)
	}
	if msg == "" {
		msg = config.GetTodoResMsg(config.UploadStr, err != nil)
	}
	f.Response.WriteJson(app.Response{
		Data: fileId,
		Status: app.Status{
			Code:  0,
			Error: err != nil,
			Msg:   msg,
		},
	})
}
