package repository

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/minio/minio-go/v7"
)

type StorageRepository struct {
	Client *minio.Client
}

func NewStorageRepository(cl *minio.Client) *StorageRepository {
	return &StorageRepository{
		Client: cl,
	}
}

func (s *StorageRepository) MakeBucket(user_id string) error {
	err := s.Client.MakeBucket(context.Background(), user_id, minio.MakeBucketOptions{ObjectLocking: false})
	return err
}

func (s *StorageRepository) UploadOne(user_id string, filename string, file io.Reader, size int64) error {
	info, err := s.Client.PutObject(context.Background(), user_id, filename, file, size, minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	fmt.Println(info.Key)
	return nil
}

func (s *StorageRepository) GetFile(user_id, filename string) ([]byte, error) {
	object, err := s.Client.GetObject(context.Background(), user_id, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	imgbyte, err := ioutil.ReadAll(object)
	if err != nil {
		return nil, err
	}
	return imgbyte, nil
}

func (s *StorageRepository) GetAllFilesList(user_id string) []string{
	var data []string
	for objectInfo := range s.Client.ListObjects(context.Background(), user_id, minio.ListObjectsOptions{Recursive: true}) {
		data = append(data, objectInfo.Key)
	}
	return data
}
