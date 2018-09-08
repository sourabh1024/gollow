//Copyright 2018 Sourabh Suman ( https://github.com/sourabh1024 )
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package storage

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sourabh1024/gollow/logging"
	"io/ioutil"
)

// s3API defines the interface for S3 API.
type s3API interface {
	GetObject(*s3.GetObjectInput) (*s3.GetObjectOutput, error)
	GetObjectRequest(*s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput)

	HeadObject(*s3.HeadObjectInput) (*s3.HeadObjectOutput, error)

	PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error)
	PutObjectRequest(*s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput)
}

// client is the wrapper for the aws s3 client
type S3Storage struct {
	client s3API
	bucket string
	key    string
}

// NewS3Client constructs a new s3 client
func NewS3Storage(s3Config *Config, cfgs ...*aws.Config) (*S3Storage, error) {
	config := &aws.Config{
		Region: aws.String(s3Config.Region),
	}
	config.MergeIn(cfgs...)

	sess, err := session.NewSession(config)
	if err != nil {
		logging.GetLogger().Error("error getting session with err : %v", err)
		return nil, err
	}

	awsClient := s3.New(sess)

	instance := &S3Storage{
		client: awsClient,
		bucket: s3Config.Bucket,
		key:    s3Config.Key,
	}
	return instance, nil
}

// IsObjectExist retrieves metadata from an object to verify whether object exists
func (storage *S3Storage) IsExist() bool {
	headObjectInput := &s3.HeadObjectInput{
		Bucket: &storage.bucket,
		Key:    &storage.key,
	}

	_, err := storage.client.HeadObject(headObjectInput)
	if err != nil {
		logging.GetLogger().Error("error getting head object from s3 for bucket %s key %s", storage.bucket, storage.key)
		return false
	}

	return true
}

// Read retrieves an object from s3
func (storage *S3Storage) Read() ([]byte, error) {
	getObjectInput := &s3.GetObjectInput{
		Bucket: &storage.bucket,
		Key:    &storage.key,
	}

	getObjectOutput, err := storage.client.GetObject(getObjectInput)
	if err != nil {
		logging.GetLogger().Error("error getting object output from s3 for bucket %s key %s", storage.bucket, storage.key)
		return nil, err
	}

	body, err := ioutil.ReadAll(getObjectOutput.Body)
	if err != nil {
		logging.GetLogger().Error("error reading object content from object output of s3 for bucket %s key %s", storage.bucket, storage.key)
	}

	return body, err
}

// Write puts an object into S3.
func (storage *S3Storage) Write(data []byte) (int, error) {
	putObjectInput := &s3.PutObjectInput{
		Bucket: &storage.bucket,
		Key:    &storage.key,
		Body:   bytes.NewReader(data),
	}

	_, err := storage.client.PutObject(putObjectInput)
	if err != nil {
		logging.GetLogger().Error("error putting object to s3 for bucket %s key %s", storage.bucket, storage.key)
		return 0, err
	}

	return 1, nil
}
