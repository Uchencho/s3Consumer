package workflow

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Uchencho/s3Consumer/internal/storage"

	"github.com/Uchencho/commons/ctime"
	"github.com/Uchencho/commons/httputils"
	"github.com/Uchencho/commons/uuid"

	"github.com/pkg/errors"
)

// RawToS3Func provides the functionality of uploading raw data to S3
type RawToS3Func func(s []httputils.FileDetails) error

// UploadFileToS3 uploads a number of files to an s3 bucket
func UploadFileToS3(UUIDGenerator uuid.GenV4Func,
	provideTime ctime.EpochProviderFunc,
	upload storage.UploadS3FileFunc) RawToS3Func {
	return func(s []httputils.FileDetails) error {

		t := provideTime().ToISO8601().DateString()

		for _, fd := range s {

			id := UUIDGenerator()

			y := strings.Split(t, "-")[0]
			m := strings.Split(t, "-")[1]

			key := fmt.Sprintf("%s/%s/%s", y, m, id)
			reader := bytes.NewReader(fd.Data)
			if err := upload(key, reader, fd.FileName, fd.ContentType); err != nil {
				return errors.Wrapf(err, "workflow - Unable to upload file %s", fd.FileName)
			}
		}
		return nil
	}
}
