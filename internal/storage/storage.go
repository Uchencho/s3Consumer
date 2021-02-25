package storage

import (
	"io"

	"github.com/Uchencho/commons/aws/storage"
)

// UploadS3FileFunc returns the functionality of uploading a file to s3
type UploadS3FileFunc func(key string, reader io.Reader, fileName, contentType string) error

// UploadFileToS3 uploads a file to s3
func UploadFileToS3(bucket string, upload storage.UploadFunc) UploadS3FileFunc {
	return func(key string, reader io.Reader, fileName, contentType string) error {
		metaDataOpt := func(r *storage.OptionalUploadRequest) {
			r.MetaData = map[string]*string{
				"clientFilename": &fileName,
			}
		}
		contentTypeOpt := func(r *storage.OptionalUploadRequest) {
			r.ContentType = contentType
		}
		return upload(reader, bucket, key, metaDataOpt, contentTypeOpt)
	}
}
