package main

import (
	"bytes"
	"cloud.google.com/go/translate"
	"context"
	"encoding/base64"
	"errors"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Handler struct {
	client *translate.Client
}

func unmarshalRequest(input *Input, r *http.Request) error {
	if r.Method != http.MethodPost {
		return errors.New("only POST allowed")
	}
	switch r.Header.Get("content-type") {
		case "application/json":
			um := jsonpb.Unmarshaler{}
			if err := um.Unmarshal(r.Body, input); err != nil {
				return errors.New("malformed json input: " + err.Error())
			}
		case "application/protobuf":
			buf := new(bytes.Buffer)
			if _, err := buf.ReadFrom(r.Body); err != nil {
				return err
			}
			err := proto.Unmarshal(buf.Bytes(), input)
			if err != nil {
				return errors.New("malformed protobuf input")
			}
		default:
			return errors.New("must use application/json or application/protobuf content type")
	}
	return nil
}

func impl(ctx context.Context, client *translate.Client, input *Input, output *Output) error {
	inputLang, err := language.Parse(input.InputLanguage)
	if err != nil {
		return err
	}
	outputLang, err := language.Parse(input.OutputLanguage)
	if err != nil {
		return err
	}

	res, err := client.Translate(ctx, []string{input.Text}, outputLang, &translate.Options{Source: inputLang})
	if err != nil {
		return err
	} else {
		output.Text = res[0].Text
	}
	return nil
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input := &Input{}
	if err := unmarshalRequest(input, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	output := &Output{}
	if err := impl(r.Context(), handler.client, input, output); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	switch r.Header.Get("content-type") {
	case "application/json":
		w.Header().Set("content-type", "application/json")
		um := jsonpb.Marshaler{}
		s, _ := um.MarshalToString(output)
		w.Write([]byte(s))
	case "application/protobuf":
		w.Header().Set("content-type", "application/protobuf")
		b, _ := proto.Marshal(output)
		w.Write(b)
	}
}

func main() {
	ctx := context.Background()
	credsBase64 := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_BASE64")
	if credsBase64 == "" {
		log.Fatal("GOOGLE_APPLICATION_CREDENTIALS_BASE64 env var required")
	}
	creds, err := base64.StdEncoding.DecodeString(credsBase64)
	if err != nil {
		log.Fatal(err)
	}
	client, err := translate.NewClient(ctx, option.WithCredentialsJSON(creds))
	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{Addr: ":9000", Handler: &Handler{client}}
	log.Print("Running on :9000")
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		s := <-c
		log.Println("Got signal:", s)
		server.Close()
	}()
	err = server.ListenAndServe()
	log.Println(err)
}
