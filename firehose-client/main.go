package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

func main() {
	mySession := session.Must(session.NewSession())
	svc := firehose.New(mySession)
	dsn := os.Getenv("DELIVERY_STREAM_NAME")
	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		data := req.URL.Query().Get("data")
		b := []byte(data)
		// svc
		_, err := svc.PutRecord(&firehose.PutRecordInput{
			DeliveryStreamName: &dsn,
			Record: &firehose.Record{
				Data: b,
			},
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("failed. data=%v dsn=%s", data, dsn)))
			return
		}

		w.Write(b)
	})

	http.ListenAndServe(":"+port, nil)
}
