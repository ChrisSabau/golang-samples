// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START documentai_quickstart]
// [START documentai_process_document]

// processDocument sends a file at a given filePath for online processing.
package snippets

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	documentai "cloud.google.com/go/documentai/apiv1"
	"cloud.google.com/go/documentai/apiv1/documentaipb"
	"google.golang.org/api/option"
)

func processDocument(w io.Writer, projectID, locationID, processorID, filePath, mimeType string) error {
	// projectID := "PROJECT_ID"
	// locationID := "us"
	// processorID := "aaaaaaaa"
	// filePath := "invoice.pdf"
	// mimeType := "application/pdf"

	ctx := context.Background()

	endpoint := fmt.Sprintf("%s-documentai.googleapis.com:443", locationID)
	client, err := documentai.NewDocumentProcessorClient(ctx, option.WithEndpoint(endpoint))
	if err != nil {
		log.Fatalf("error creating Document AI client: %v", err)
		return err
	}
	defer client.Close()

	// Open local file.
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("ioutil.ReadFile: %v", err)
		return err
	}

	req := &documentaipb.ProcessRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/processors/%s", projectID, locationID, processorID),
		Source: &documentaipb.ProcessRequest_RawDocument{
			RawDocument: &documentaipb.RawDocument{
				Content:  data,
				MimeType: mimeType,
			},
		},
	}
	resp, err := client.ProcessDocument(ctx, req)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Handle the results.
	document := resp.GetDocument()
	fmt.Fprintf(w, "Document Text: %s", document.GetText())
	return nil
}

// [END documentai_quickstart]
// [END documentai_process_document]