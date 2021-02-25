package app

import (
	"net/http"
	"os"

	"github.com/Uchencho/commons/aws/storage"
	"github.com/Uchencho/commons/ctime"
	"github.com/Uchencho/commons/pubsub"
	"github.com/Uchencho/commons/uuid"
	"github.com/julienschmidt/httprouter"

	commonAWS "github.com/Uchencho/commons/aws"
	"github.com/Uchencho/s3Consumer/internal/pkg"
	internalStorage "github.com/Uchencho/s3Consumer/internal/storage"
)

// App is a representation of the application
type App struct {
	RawUploadHandler http.HandlerFunc
	ZipFileConsumer  pubsub.ConsumerFunc
}

// OptionalArgs is a representation of all the optional arguments for this application
type OptionalArgs struct {
	UUIDGenerator uuid.GenV4Func
	TimeProvider  ctime.EpochProviderFunc
	S3Uploader    storage.UploadFunc
	S3Bucket      string
}

// Option is a representation of a function that modifies optional arguments
type Option func(oa *OptionalArgs)

// New instantiates a new app
func New(opts ...Option) App {

	c := commonAWS.ConfigFromEnvVars()

	oa := OptionalArgs{
		UUIDGenerator: uuid.GenV4,
		TimeProvider:  ctime.CurrentEpoch,
		S3Uploader:    storage.UploadToS3(c),
		S3Bucket:      os.Getenv("S3_BUCKET"),
	}

	for _, opt := range opts {
		opt(&oa)
	}

	uploadFiles := internalStorage.UploadFileToS3(oa.S3Bucket, oa.S3Uploader)
	uploadHandler := RawUploadHandler(oa.UUIDGenerator, oa.TimeProvider, uploadFiles)
	handleZip := HandleUploadConsumer()

	return App{
		RawUploadHandler: uploadHandler,
		ZipFileConsumer:  handleZip,
	}
}

// Handler returns an http handler for the application
func (a *App) Handler() http.HandlerFunc {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/upload", a.RawUploadHandler)

	h := http.HandlerFunc(router.ServeHTTP)
	return h
}

func (a *App) Consumer() pubsub.ConsumerFunc {
	router := pubsub.NewConsumerRouter()
	router.Register(pkg.HandleUploadMessageType(), a.ZipFileConsumer)
	return pubsub.DefaultConsumerWrapper(os.Stdout)(router.Consume)
}
