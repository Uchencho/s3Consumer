package app

import (
	"net/http"

	"github.com/Uchencho/s3Consumer/internal/storage"
	"github.com/Uchencho/s3Consumer/internal/workflow"

	"github.com/Uchencho/commons/ctime"
	"github.com/Uchencho/commons/httputils"
	"github.com/Uchencho/commons/uuid"
)

const (
	fileUploadKey = "masterFiles"
)

// RawUploadHandler is the handler in charge of handling raw file uploads to s3
func RawUploadHandler(UUIDGenerator uuid.GenV4Func,
	provideTime ctime.EpochProviderFunc,
	uploadFiles storage.UploadS3FileFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		files, err := httputils.ExtractMultipleFileUploads(r, fileUploadKey)
		if err != nil {
			httputils.ServeError(err, w)
			return
		}

		handleUpload := workflow.UploadFileToS3(UUIDGenerator, provideTime, uploadFiles)
		if err := handleUpload(files); err != nil {
			httputils.ServeError(err, w)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
