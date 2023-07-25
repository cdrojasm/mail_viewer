package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime/pprof"
	"strings"
	"sync"
	"time"
)

type doc_content_ struct {
	Message_ID                string   `json:Message_ID`
	Date                      string   `json:Date`
	From                      string   `json:From`
	To                        []string `json:To`
	Subject                   string   `json:Subject`
	Mime_version              string   `json:Mime_version`
	Content_Type              string   `json:content_Type`
	Content_Transfer_Encoding string   `json:content_Transfer_Encoding`
	Bcc                       []string `json:Bcc`
	X_from                    string   `json:X_from`
	X_to                      []string `json:X_to`
	X_cc                      []string `json:X_cc`
	X_bcc                     []string `json:X_bcc`
	X_folder                  string   `json:X_folder`
	X_origin                  string   `json:X_origin`
	X_filename                string   `json:X_filename`
	Content                   string   `json:Content`
}

type bulk_data_ struct {
	Index   string         `json:"index"`
	Records []doc_content_ `json:"records"`
}

type FileTask struct {
	filepath string
	result   string
	payload  doc_content_
}

var urlSearchZinc string = "http://localhost:4080/api/enron_mail__/_doc"
var urlSearchZinc_bulk string = "http://localhost:4080/api/_bulkv2"

func processFile(id int, wg *sync.WaitGroup, taskCh <-chan FileTask, resultCh chan<- FileTask) {
	// process file
	defer wg.Done()
	for task := range taskCh {

		fileName := task.filepath
		currentNameFile := strings.Join(strings.Split(fileName, "\\"), "/")
		fmt.Println(currentNameFile, "es un archivo")
		stringLoadedDataFile, err := loadStringDataFile_asBytes(currentNameFile)
		if err != nil {
			task.result = "file cant be loaded"
			resultCh <- task
		}
		docContent := processingDocAsBytes(stringLoadedDataFile)
		/* jsonDataBin, err := json.Marshal(docContent)
		if err != nil {
			task.result = "error during json marshall"
			resultCh <- task
		}
		errIndex := indexingDataToSearchZinc_2(urlSearchZinc, jsonDataBin, "admin", "Complexpass#123")
		if errIndex != nil {
			task.result = "error during json marshall"
			resultCh <- task
		} */
		task.result = "success"
		task.payload = docContent
		resultCh <- task
	}
}

func main() {

	f, err := os.Create("cpu_profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var nonProcessedDocs = []string{}
	var production bool = true
	if production {
		numWorkers := 8
		taskCh := make(chan FileTask)
		resultCh := make(chan FileTask)
		var wg sync.WaitGroup
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go processFile(i, &wg, taskCh, resultCh)
		}
		go func() {
			filepath.Walk("../enronData/enron_mail_20110402/maildir", func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Println(err)
					return err
				}
				// check if path is a file or directory
				// if file then write string to FilesCount
				if !info.IsDir() {
					taskCh <- FileTask{filepath: path, result: "success"}
				}
				return nil
			})
			close(taskCh)
		}()
		go func() {
			batchSize := 50
			var batch []doc_content_
			for task := range resultCh {
				if task.result != "success" {
					nonProcessedDocs = append(nonProcessedDocs, task.filepath+" "+task.result)
				} else {
					batch = append(batch, task.payload)
					if len(batch) == batchSize {
						bulkData := bulk_data_{Index: "enron_mail__", Records: batch}
						jsonDataBin, err := json.Marshal(bulkData)
						if err != nil {
							fmt.Println(err)
						}
						errIndex := indexingDataToSearchZinc_2(urlSearchZinc_bulk, jsonDataBin, "admin", "Complexpass#123")
						if errIndex != nil {
							fmt.Println(errIndex)
						}
						batch = []doc_content_{}
					}
				}
			}
		}()
		wg.Wait()
		close(resultCh)
	}
}

func byteSlice_slice_to_stringSlice(data [][]byte) []string {
	var stringSlice []string
	for _, value := range data {
		stringSlice = append(stringSlice, string(value))
	}
	return stringSlice
}

func processingDocAsBytes(data []byte) doc_content_ {
	copyData := bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))
	splitedDoc_Meta_Content := bytes.SplitN(copyData, []byte("\n\n"), 2)
	copyData = splitedDoc_Meta_Content[0]
	content := splitedDoc_Meta_Content[1]
	var contentCurrentDocument doc_content_
	contentCurrentDocument.Content = string(content)
	var contentTagData []byte
	var tagNamesString = [][]byte{[]byte("Message-ID:"), []byte("\nDate: "), []byte("\nFrom: "), []byte("\nTo: "),
		[]byte("\nSubject: "), []byte("\nCc: "), []byte("\nMime-Version: "), []byte("\nContent-Type: "), []byte("\nContent-Transfer-Encoding: "), []byte("\nX-From: "),
		[]byte("\nX-To: "), []byte("\nX-cc: "), []byte("\nX-bcc: "), []byte("\nX-Folder: "), []byte("\nX-Origin: "), []byte("\nX-FileName: ")}
	inputLayout := "2 Jan 2006 15:04:05 -0700"
	outputLayout := "2006-01-02T15:04:05Z07:00"

	tagIndex := 0
	currentTag := []byte("")
	currentSecondTag := []byte("")
	lenCurrentTag := 0
	indexCurrentTagInData := 0
	indexSecondTagInData := 0
	re := regexp.MustCompile(`(\t|\n)`)
	for tagIndex < len(tagNamesString) {
		currentTag = tagNamesString[tagIndex]
		lenCurrentTag = len(currentTag)
		secondTagIndex := tagIndex + 1
		indexCurrentTagInData = bytes.Index(copyData, currentTag)
		if indexCurrentTagInData != -1 {
			if bytes.Equal(currentTag, []byte("\nX-FileName: ")) {
				contentTagData = copyData[indexCurrentTagInData+lenCurrentTag : len(copyData)]
				copyData = []byte("")
				contentCurrentDocument.X_filename = string(contentTagData)
			} else {
				for secondTagIndex < len(tagNamesString) {
					currentSecondTag = tagNamesString[secondTagIndex]
					indexSecondTagInData = bytes.Index(copyData, currentSecondTag)
					if indexSecondTagInData != -1 {
						tagNameMatchCase := bytes.Trim(bytes.TrimSpace(bytes.Trim(currentTag, "\n")), ":")
						contentTagData = copyData[indexCurrentTagInData+lenCurrentTag : indexSecondTagInData]
						copyData = copyData[indexSecondTagInData:len(copyData)]
						if bytes.Equal(tagNameMatchCase, []byte("Message-ID")) {
							contentCurrentDocument.Message_ID = string(bytes.Trim(bytes.Trim(bytes.TrimSpace(contentTagData), ">"), "<"))
						} else if bytes.Equal(tagNameMatchCase, []byte("Date")) {
							contentTagData = bytes.TrimSpace(bytes.Split(bytes.Split(contentTagData, []byte(","))[1], []byte("("))[0])
							t, err := time.Parse(inputLayout, string(contentTagData))
							if err != nil {
								fmt.Println("Error parsing date:", err)
							} else {
								contentTagData := t.Format(outputLayout)
								contentCurrentDocument.Date = contentTagData
							}
						} else if bytes.Equal(tagNameMatchCase, []byte("From")) {
							contentCurrentDocument.From = string(contentTagData)
						} else if bytes.Equal(tagNameMatchCase, []byte("To")) {
							contentTagData = bytes.Trim(contentTagData, "\t")
							contentTagData_ := re.ReplaceAllString(string(contentTagData), "")
							contentCurrentDocument.To = strings.Split(contentTagData_, ",")
						} else if bytes.Equal(tagNameMatchCase, []byte("Subject")) {
							contentCurrentDocument.Subject = string(contentTagData)
						} else if bytes.Equal(tagNameMatchCase, []byte("Mime-Version")) {
							contentCurrentDocument.Mime_version = string(contentTagData)
						} else if bytes.Equal(tagNameMatchCase, []byte("Content-Type")) {
							contentCurrentDocument.Content_Type = string(contentTagData)
						} else if bytes.Equal(tagNameMatchCase, []byte("Content-Transfer-Encoding")) {
							contentCurrentDocument.Content_Transfer_Encoding = string(contentTagData)
						} else if bytes.Equal(tagNameMatchCase, []byte("X-From")) {
							contentCurrentDocument.X_from = string(contentTagData)
						} else if bytes.Equal(tagNameMatchCase, []byte("X-To")) {
							contentTagData = bytes.Trim(contentTagData, "\t")
							contentCurrentDocument.X_to = byteSlice_slice_to_stringSlice(bytes.Split(contentTagData, []byte(",")))
						} else if bytes.Equal(tagNameMatchCase, []byte("X-cc")) {
							contentCurrentDocument.X_cc = byteSlice_slice_to_stringSlice(bytes.Split(contentTagData, []byte(",")))
						} else if bytes.Equal(tagNameMatchCase, []byte("X-bcc")) {
							contentCurrentDocument.X_bcc = byteSlice_slice_to_stringSlice(bytes.Split(contentTagData, []byte(",")))
						} else if bytes.Equal(tagNameMatchCase, []byte("X-Folder")) {
							contentCurrentDocument.X_folder = string(contentTagData)
						} else if bytes.Equal(tagNameMatchCase, []byte("X-Origin")) {
							contentCurrentDocument.X_origin = string(contentTagData)
						} else if bytes.Equal(tagNameMatchCase, []byte("X-FileName")) {
							contentCurrentDocument.X_filename = string(contentTagData)
						}
						secondTagIndex = len(tagNamesString)
					}
					secondTagIndex += 1
				}
			}
		}
		tagIndex += 1
	}
	/* fmt.Println("doc procesado")
	fmt.Printf("%+v", contentCurrentDocument) */
	return contentCurrentDocument
}

func loadStringDataFile_asBytes(nameFile string) ([]byte, error) {
	fileReader, errReading := ioutil.ReadFile(nameFile)
	if errReading != nil {
		fmt.Println("Could not read the file due file reading error ", nameFile, errReading)
		return nil, errReading
	}
	return fileReader, errReading
}

func indexingDataToSearchZinc_2(url string, JSONData []byte, userName string, password string) error {
	fmt.Println("indexando data en searchzinc", url)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(JSONData))
	if err != nil {
		return err
	}
	username := userName
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("bad status: %s", resp.Status)
	} else {
		fmt.Println("indexando data en searchzinc")
		return nil
	}

}
