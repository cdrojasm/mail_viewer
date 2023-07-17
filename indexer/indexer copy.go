package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
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
	"time"
)

type doc_content struct {
	Message_ID                string   `json:Message_ID`
	Date                      string   `json:Date`
	From                      string   `json:From`
	To                        []string `json:To`
	Subject                   string   `json:Subject`
	Cc                        []string `json:cc`
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

func main() {

	f, err := os.Create("cpu_profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	urlSearchZinc := "http://localhost:4080/api/enron_mail__/_doc"
	fileList := "./listOfDocuments_test.csv"
	var currentNameFile string

	readFile, err := os.Open(fileList)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()
	for _, line := range fileLines {
		currentNameFile = strings.Join(strings.Split(line, "\\"), "/")
		infoFile, err := os.Stat(currentNameFile)
		if err != nil {
			fmt.Println(err)
		}
		if infoFile.IsDir() {
			fmt.Println(currentNameFile, "es un directorio no un archivo ")
		} else {
			stringLoadedDataFile := loadStringDataFile(currentNameFile)
			docContent := processingDoc(stringLoadedDataFile)
			jsonDataBin, err := json.Marshal(docContent)
			if err != nil {
				panic(err)
			}
			errIndex := indexingDataToSearchZinc(urlSearchZinc, jsonDataBin, "admin", "Complexpass#123")
			if errIndex != nil {
				fmt.Println(errIndex)
			}
		}
		//fmt.Println(currentNameFile)
	}
}

func processingDoc(data string) doc_content {

	copyData := strings.ReplaceAll(data, "\r\n", "\n")
	splitedDoc_Meta_Content := strings.SplitN(copyData, "\n\n", 2)
	copyData = splitedDoc_Meta_Content[0]
	content := splitedDoc_Meta_Content[1]
	var contentCurrentDocument doc_content
	contentCurrentDocument.Content = content
	var contentTagData string
	var tagNamesString = []string{"Message-ID:", "\nDate: ", "\nFrom: ", "\nTo: ",
		"\nSubject: ", "\nCc: ", "\nMime-Version: ", "\nContent-Type: ", "\nContent-Transfer-Encoding: ", "\nX-From: ",
		"\nX-To: ", "\nX-cc: ", "\nX-bcc: ", "\nX-Folder: ", "\nX-Origin: ", "\nX-FileName: "}
	//inputDateString := "Tue, 5 Dec 2000 02:45:00 -0800 (PST)"
	inputLayout := "2 Jan 2006 15:04:05 -0700"
	outputLayout := "2006-01-02T15:04:05Z07:00"

	tagIndex := 0
	currentTag := ""
	currentSecondTag := ""
	lenCurrentTag := 0
	indexCurrentTagInData := 0
	indexSecondTagInData := 0
	for tagIndex < len(tagNamesString) {
		currentTag = tagNamesString[tagIndex]
		lenCurrentTag = len(currentTag)
		secondTagIndex := tagIndex + 1
		indexCurrentTagInData = strings.Index(copyData, currentTag)
		if indexCurrentTagInData != -1 {
			if currentTag == "\nX-FileName: " {
				contentTagData = copyData[indexCurrentTagInData+lenCurrentTag : len(copyData)]
				copyData = ""
				contentCurrentDocument.X_filename = contentTagData
			} else {
				for secondTagIndex < len(tagNamesString) {
					currentSecondTag = tagNamesString[secondTagIndex]
					indexSecondTagInData = strings.Index(copyData, currentSecondTag)
					if indexSecondTagInData != -1 {
						tagNameMatchCase := strings.Trim(strings.TrimSpace(strings.Trim(currentTag, "\n")), ":")
						contentTagData = copyData[indexCurrentTagInData+lenCurrentTag : indexSecondTagInData]
						copyData = copyData[indexSecondTagInData:len(copyData)]
						switch tagNameMatchCase {
						case "Message-ID":
							contentCurrentDocument.Message_ID = strings.Trim(strings.Trim(strings.TrimSpace(contentTagData), ">"), "<")
						case "Date":
							contentTagData = strings.TrimSpace(strings.Split(strings.Split(contentTagData, ",")[1], "(")[0])
							t, err := time.Parse(inputLayout, contentTagData)
							if err != nil {
								fmt.Println("Error parsing date:", err)
								//contentCurrentDocument.Date = ""
							} else {
								contentTagData := t.Format(outputLayout)
								contentCurrentDocument.Date = contentTagData
							}
						case "From":
							contentCurrentDocument.From = contentTagData
						case "To":
							contentTagData = strings.Trim(contentTagData, "\t")
							re := regexp.MustCompile(`(\t|\n)`)
							contentTagData = re.ReplaceAllString(contentTagData, "")
							contentCurrentDocument.To = strings.Split(contentTagData, ",")
						case "Cc":
							contentCurrentDocument.From = contentTagData
						case "Subject":
							contentCurrentDocument.Subject = contentTagData
						case "Mime-Version":
							contentCurrentDocument.Mime_version = contentTagData
						case "Content-Type":
							contentCurrentDocument.Content_Type = contentTagData
						case "Content-Transfer-Encoding":
							contentCurrentDocument.Content_Transfer_Encoding = contentTagData
						case "X-From":
							contentCurrentDocument.X_from = contentTagData
						case "X-To":
							contentTagData = strings.Trim(contentTagData, "\t")
							contentCurrentDocument.X_to = strings.Split(contentTagData, ",")
						case "X-cc":
							contentCurrentDocument.X_cc = strings.Split(contentTagData, ",")
						case "X-bcc":
							contentCurrentDocument.X_bcc = strings.Split(contentTagData, ",")
						case "X-Folder":
							contentCurrentDocument.X_folder = contentTagData
						case "X-Origin":
							contentCurrentDocument.X_origin = contentTagData
						case "X-FileName":
							contentCurrentDocument.X_filename = contentTagData
						}
						secondTagIndex = len(tagNamesString)
					}
					secondTagIndex += 1
				}
			}

		}
		tagIndex += 1
	}
	/* fmt.Printf("%+v", contentCurrentDocument) */
	return contentCurrentDocument
}

func loadStringDataFile(nameFile string) string {
	fmt.Println(nameFile)
	fileReader, errReading := ioutil.ReadFile(nameFile)
	if errReading != nil {
		fmt.Println("Could not read the file due file reading error ", nameFile, errReading)
		os.Exit(0)
	}
	return string(fileReader)
}

func indexingDataToSearchZinc(url string, JSONData []byte, userName string, password string) error {
	fmt.Println("indexando data en searchzinc")
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
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// Handle the response as needed
	fmt.Println(resp)
	fmt.Println(string(responseBody))
	return nil
}

func buildDocsTree() bool {
	var filesCount int
	rootDirectory := "../enronData/enron_mail_20110402"
	statusDirectory := "./listOfDocuments.csv"
	listDocuments := []string{}
	wasSuccesfullDocTreeBuild_ := true
	var csvFile *os.File
	if _, err := os.Stat(rootDirectory); os.IsNotExist(err) {
		fmt.Println("no existe el directorio de data")
		return !wasSuccesfullDocTreeBuild_
	}
	if _, err := os.Stat(statusDirectory); os.IsNotExist(err) {
		fmt.Println("creando el archivo de status de la indexacion")

		csvFile, err = os.Create(statusDirectory)
		if err != nil {
			log.Fatal(err)
			return !wasSuccesfullDocTreeBuild_
		}
	} else {
		fmt.Println("el archivo ya ha sido construidos")
		return wasSuccesfullDocTreeBuild_
	}
	csvwriter := csv.NewWriter(csvFile)
	err := filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if path != "false" {
			//fmt.Printf("dir: %v: name: %s currenCount:%v\n", info.IsDir(), path, filesCount)
			_ = csvwriter.Write([]string{path})
			listDocuments = append(listDocuments, path)
			filesCount += 1
		}
		return nil
	})
	csvwriter.Flush()
	if err != nil {
		fmt.Println(err)
		return !wasSuccesfullDocTreeBuild_
	}
	return wasSuccesfullDocTreeBuild_
}
