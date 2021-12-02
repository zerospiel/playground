package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/trace"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/minio/minio-go/v7"
	mcredentials "github.com/minio/minio-go/v7/pkg/credentials"
	generatedv1 "github.com/zerospiel/playground/gen/go/pg/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	hc "moul.io/http2curl/v2"
)

//go:generate go run google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0 --version
func noop() {}

type foobar struct {
	generatedv1.UnimplementedStringsServiceServer
}

func (*foobar) ToUpper(ctx context.Context, req *generatedv1.ToUpperRequest) (*generatedv1.ToUpperResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	fmt.Println(metadata.FromIncomingContext(ctx))
	fmt.Println(metadata.FromOutgoingContext(ctx))
	return &generatedv1.ToUpperResponse{
		S: strings.ToUpper(req.S),
	}, nil
}

type sss struct{}

func (*sss) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bb, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(bb))
	w.WriteHeader(200)
	return
}

func main() {

	// minioCall()
	// awsCall()

	return

	req, err := http.NewRequest(http.MethodPost, "http://ru-central3.s3.s.o3.ru:7480/o2-models/prod/xgboost/dm7971_model_wo_clicks_1623311229/features_map.json", nil)
	// req.Header.Add("Authorization", "AWS4-HMAC-SHA256 Credential=DQ05ARP5A76243ATNXN3/20211008/ozon-central1/s3/aws4_request, SignedHeaders=host;x-amz-content-sha256;x-amz-date, Signature=91310c35e96c7c854809f86e0b5539113b684aea3f2c07b91e700fd3f2f8be86")
	// req.Header.Add("X-Amz-Content-Sha256", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	// req.Header.Add("X-Amz-Date", "20211008T092136Z")
	// req.Header.Add("Accept-Encoding", "application/json")
	req.PostForm.Add("Authorization", "AWS4-HMAC-SHA256 Credential=DQ05ARP5A76243ATNXN3/20211008/ozon-central1/s3/aws4_request, SignedHeaders=host;x-amz-content-sha256;x-amz-date, Signature=91310c35e96c7c854809f86e0b5539113b684aea3f2c07b91e700fd3f2f8be86")
	req.PostForm.Add("X-Amz-Content-Sha256", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	req.PostForm.Add("X-Amz-Date", "20211008T092136Z")
	req.PostForm.Add("Accept-Encoding", "application/json")
	if err != nil {
		panic(err)
	}

	bb, _ := httputil.DumpRequest(req, false)
	println(string(bb))
	cmd, err := hc.GetCurlCommand(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cmd.String(): %v\n", cmd.String())

	return
}

func minioCall() {
	creds := mcredentials.NewStaticV4(
		"9FN7VH8OSVE6KP0OWOV6",
		"FTpoOSS3uRHOPXxYZ4Nb8wUtTngVnTnKh0oWcfPM",
		"")

	minioClient, err := minio.New(
		"proxy.s3.local:7000",
		// "prod.s3.ceph.s.o3.ru:7480",
		// "ru-central2.s3.s.o3.ru:7480",
		// "dev-ru-central2.s3.s.o3.ru:8080",
		&minio.Options{
			Creds: creds,
		})
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	lbresp, err := minioClient.ListBuckets(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("lbresp: %v\n", lbresp)

	loC := minioClient.ListObjects(ctx, "non-existant-bucket", minio.ListObjectsOptions{MaxKeys: 5, UseV1: false})
	for v := range loC {
		fmt.Printf("v: %v\n", v)
	}

	poresp, err := minioClient.PutObject(ctx, "non-existant-bucket", "/pathminiogovno/mmorgoev", bytes.NewReader([]byte(`foobar\n`)), 7, minio.PutObjectOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("poresp: %v\n", poresp)

}

func awsCall() {
	creds := credentials.NewStaticCredentials(
		"8H7BG2GY8P86V6Z87D6A",
		"n0NvI2nXEeAeGrLHDewTADEHNw4lVLjeFsLaxUaQ",
		"")
	// creds := credentials.NewStaticCredentials(
	// 	"9FN7VH8OSVE6KP0OWOV6",
	// 	"FTpoOSS3uRHOPXxYZ4Nb8wUtTngVnTnKh0oWcfPM",
	// 	"")
	_, err := creds.Get()
	if err != nil {
		panic(err)
	}

	c := aws.NewConfig().
		WithRegion("us-east-1").
		WithCredentials(creds).
		// WithEndpoint("any.s3front.platform.svc.dev.k8s.o3.ru:8080").
		// WithEndpoint("any.s3front.platform.svc.prod.k8s.o3.ru:80").
		// WithEndpoint("http://prod.s3.ceph.s.o3.ru:7000").
		WithEndpoint("http://ru-central3.s3.s.o3.ru:7480").
		// WithEndpoint("http://ru-central2.s3.s.o3.ru:7480").
		// WithEndpoint("http://dev-ru-central2.s3.s.o3.ru:8080").
		// WithEndpoint("http://dev-ru-central1.s3.s.o3.ru:8080").
		// WithEndpoint("http://proxy.s3.local:7000").
		WithS3ForcePathStyle(true).
		WithDisableSSL(true).
		WithLowerCaseHeaderMaps(true).
		WithLogLevel(aws.LogDebug).
		WithLogger(aws.NewDefaultLogger())

	var sess *session.Session
	if sess, err = session.NewSession(c); err != nil {
		panic(err)
	}

	svc := s3.New(sess, c)
	_ = svc

	goresp, _ := svc.GetObject(&s3.GetObjectInput{Bucket: aws.String("webcrm"), Key: aws.String("attachments/6ccc0f51-e3c3-46d3-93da-586f5e997158/")})
	bb, _ := io.ReadAll(goresp.Body)
	fmt.Println(string(bb))
	goresp, err = svc.GetObject(&s3.GetObjectInput{Bucket: aws.String("webcrm"), Key: aws.String("attachments/0000f590-9dac-40d2-aa8f-ec59c163d5a3/")})
	if err != nil {
		fmt.Println(err)
		return
	}
	bb, _ = io.ReadAll(goresp.Body)
	fmt.Println(string(bb))
}

func getTrace() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	trace.Stop()
}
