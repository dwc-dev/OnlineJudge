package mc

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"image"
	"image/jpeg"
	_ "image/png"

	"backend/microservices/user/internal/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	client *minio.Client
	config config.Minio
}

func NewMinioClient(config config.Minio) (*MinioClient, error) {
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return nil, err
	}
	return &MinioClient{client: minioClient, config: config}, nil
}

func (c *MinioClient) UploadAvatar(ctx context.Context, userID uint64, avatar_base64 string) error {
	// 解码 base64 成字节
	avatarBytes, err := base64.StdEncoding.DecodeString(avatar_base64)
	if err != nil {
		return err
	}

	// 解码成 image.Image
	img, _, err := image.Decode(bytes.NewReader(avatarBytes))
	if err != nil {
		return err
	}

	// 转换为 JPEG 格式
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100})
	if err != nil {
		return err
	}

	objectName := fmt.Sprintf("%s/%d.jpg", c.config.Avatar.Prefix, userID)
	contentType := "image/jpeg"

	// 上传到 MinIO
	_, err = c.client.PutObject(ctx, c.config.Bucket, objectName, &buf, int64(buf.Len()), minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

// func (c *MinioClient) GetBucketPolicy(ctx context.Context, bucketName string) (policy string, err error) {
// 	policy, err = c.client.GetBucketPolicy(ctx, bucketName)
// 	if err != nil {
// 		return "", err
// 	}
// 	return policy, nil
// }

func (c *MinioClient) GetAvatarUrl(userID uint64) (url string) {
	objectName := fmt.Sprintf("%s/%d.jpg", c.config.Avatar.Prefix, userID)
	if c.config.UseSSL {
		url = fmt.Sprintf("https://%s/%s/%s", c.config.Address, c.config.Bucket, objectName)
	} else {
		url = fmt.Sprintf("http://%s/%s/%s", c.config.Address, c.config.Bucket, objectName)
	}
	return url
}

func (c *MinioClient) GetDefaultAvatarUrl() (url string) {
	objectName := fmt.Sprintf("%s/%s", c.config.Avatar.Prefix, c.config.Avatar.Default)
	if c.config.UseSSL {
		url = fmt.Sprintf("https://%s/%s/%s", c.config.Address, c.config.Bucket, objectName)
	} else {
		url = fmt.Sprintf("http://%s/%s/%s", c.config.Address, c.config.Bucket, objectName)
	}
	return url
}
