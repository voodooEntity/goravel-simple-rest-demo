package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"path/filepath"
	"strconv"
	"time"
)

type FileController struct {
	//Dependent services
}

func NewFileController() *FileController {
	return &FileController{
		//Inject services
	}
}

func (r *FileController) Upload(ctx http.Context) http.Response {
	file, err := ctx.Request().File("file")

	fileIdentificator := buildFileIdentificator(file.GetClientOriginalName())
	_, err = facades.Storage().PutFileAs("", file, fileIdentificator)

	if nil != err {
		return ctx.Response().Status(422).Json(map[string]string{"error": "error writing file" + err.Error()})
	}

	ret := map[string]map[string]string{}
	ret["data"] = map[string]string{"url": "http://127.0.0.1:3000/api/file/" + fileIdentificator}

	return ctx.Response().Json(http.StatusCreated, ret)
}

func (r *FileController) Get(ctx http.Context) http.Response {
	ident := ctx.Request().Route("ident")
	if facades.Storage().Exists(ident) {
		ext := filepath.Ext(ident)
		imageData, err := facades.Storage().Get(ident)
		if err != nil {
			return ctx.Response().Status(500).Json(map[string]string{"error": "error reading file " + err.Error()})
		}
		return ctx.Response().Data(200, "image/"+ext[1:], []byte(imageData))
	}
	return ctx.Response().Status(404).Json(map[string]string{"error": "unknown file requested " + ident})
}

func buildFileIdentificator(origFileName string) string {
	ext := filepath.Ext(origFileName)
	unixTime := time.Now().Unix()
	str := origFileName + strconv.Itoa(int(unixTime))
	data := []byte(str)
	hasher := md5.New()
	hasher.Write(data)
	hash := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hash)
	finalName := hashedString + ext
	return finalName
}
