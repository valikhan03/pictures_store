package repository

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)


type UploadRepository struct{
	Client *minio.Client
}

func NewUploadRespository(cl *minio.Client) *UploadRepository{
	return &UploadRepository{
		Client: cl,
	}
}


func (u *UploadRepository) UploadOne(user_id string, filename string, file io.Reader, size int64) error{
	info, err := u.Client.PutObject(context.Background(), user_id, filename, file, size, minio.PutObjectOptions{})
	if err != nil{
		return err
	}
	fmt.Println(info.Key)
	return nil
}
