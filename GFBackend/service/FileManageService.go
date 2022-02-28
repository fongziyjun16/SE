package service

import (
	"GFBackend/logger"
	"GFBackend/model"
	"GFBackend/model/dao"
	"GFBackend/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"mime/multipart"
	"sync"
)

var fileManageServiceLock sync.Mutex
var fileManageService *FileManageService

type IFileManageService interface {
	GetSpaceInfo(username string) (model.Space, error)
	GetUserFiles(username string) ([]string, error)
	RegisterSpaceInfo(username string) error
	UpdateRemaining(username string) error
	ExpandSize(username string, newSize float64) error
	FreeSpace(username string) error
	Upload(context *gin.Context, username string, file *multipart.FileHeader) error
	Download(context *gin.Context, username string, filename string) error
}

type FileManageService struct {
	spaceDAO dao.ISpaceDAO
}

func NewFileManageService(spaceDAO dao.ISpaceDAO) *FileManageService {
	if fileManageService == nil {
		fileManageServiceLock.Lock()
		if fileManageService == nil {
			fileManageService = &FileManageService{
				spaceDAO: spaceDAO,
			}
		}
		fileManageServiceLock.Unlock()
	}
	return fileManageService
}

var FileManageServiceSet = wire.NewSet(
	dao.NewSpaceDAO,
	wire.Bind(new(dao.ISpaceDAO), new(*dao.SpaceDAO)),
	NewFileManageService,
)

func (fileManageService FileManageService) GetSpaceInfo(username string) (model.Space, error) {
	spaceInfo, err1 := fileManageService.spaceDAO.GetSpaceInfo(username)
	if err1 != nil {
		logger.AppLogger.Error(err1.Error())
		return model.Space{}, errors.New("500")
	}
	if spaceInfo.Username == "" {
		return model.Space{}, errors.New("400")
	}
	return spaceInfo, nil
}

func (fileManageService FileManageService) GetUserFiles(username string) ([]string, error) {
	filenames, err := utils.GetFilesInDir(username)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return nil, errors.New("500")
	}
	return filenames, nil
}

func (fileManageService FileManageService) RegisterSpaceInfo(username string) error {
	err := fileManageService.spaceDAO.CreateSpaceInfo(username, nil)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return errors.New("500")
	}
	return nil
}

func (fileManageService FileManageService) UpdateRemaining(username string) error {
	usedSize, err2 := utils.DirSize(username)
	if err2 != nil {
		logger.AppLogger.Error(err2.Error())
		return errors.New("500")
	}
	err3 := fileManageService.spaceDAO.UpdateRemaining(username, usedSize, nil)
	if err3 != nil {
		logger.AppLogger.Error(err3.Error())
		return errors.New("500")
	}
	return nil
}

func (fileManageService FileManageService) ExpandSize(username string, newSize float64) error {
	err := fileManageService.spaceDAO.UpdateCapacity(username, newSize, nil)
	if err != nil {
		return errors.New("500")
	}
	return nil
}

func (fileManageService FileManageService) FreeSpace(username string) error {
	err := fileManageService.spaceDAO.DeleteSpaceInfo(username, nil)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return errors.New("500")
	}
	utils.DeleteDir(username)
	return nil
}

func (fileManageService FileManageService) Upload(context *gin.Context, username string, file *multipart.FileHeader) error {
	filename := utils.DirBasePath + username + "/" + file.Filename
	if err1 := context.SaveUploadedFile(file, filename); err1 != nil {
		logger.AppLogger.Error(err1.Error())
		return errors.New("500")
	}
	return nil
}

func (fileManageService FileManageService) Download(context *gin.Context, username string, filename string) error {
	context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	context.Writer.Header().Add("Content-Type", "application/octet-stream")
	context.File(utils.DirBasePath + username + "/" + filename)
	return nil
}