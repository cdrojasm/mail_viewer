package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	os.Setenv("IP", "192.168.200.107")
	os.Setenv("Port", "4080")
	os.Setenv("name_collecion", "enron_mail__")
	os.Setenv("password", "Complexpass#123")
	os.Setenv("user", "admin")
	fmt.Println("starting API")
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	r.Route("/documents", func(r chi.Router) {
		r.Use(DocumentCtx)      // load the document according to request url parameters
		r.Get("/", GetDocument) // GET /documents/123
	})

	// Passing -routes to the program will generate docs for the above
	// router definition. See the `routes.json` file in this folder for
	// the output.
	if *routes {
		fmt.Println(docgen.JSONRoutesDoc(r))
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/go-chi/chi/v5",
			Intro:       "Welcome to the chi/_examples/rest generated docs.",
		}))
		return
	}
	http.ListenAndServe(":3000", r)

}

// DocumnetCtx middleware is used to load the document from zincsearch
func DocumentCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		var default_page int64 = 0
		var default_amount_docs int64 = 20
		queryparameters := r.URL.Query()
		var documentID = queryparameters.Get("documentID")
		var page_str = queryparameters.Get("page")
		var amountDocs_str = queryparameters.Get("amountDocs")
		page, err := strconv.ParseInt(page_str, 10, 64)
		if err != nil {
			/* add an error handler with msg */
			render.Render(w, r, ErrNotFound)
			return
		}
		amountDocs, err := strconv.ParseInt(amountDocs_str, 10, 64)
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}
		/* make refactory to pass parameters in a clear way */
		if page < -1 {
			page = default_page
		}
		if amountDocs < -1 {
			amountDocs = default_amount_docs
		}

		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "documentID", documentID)
		ctx = context.WithValue(ctx, "page", page)
		ctx = context.WithValue(ctx, "amountDocs", amountDocs)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetDocument(w http.ResponseWriter, r *http.Request) {
	var document_id_string_query string
	var page int64
	var amount_docs int64
	document_id_string_query_value := r.Context().Value("documentID")
	if document_id_string_query_value != nil {
		document_id_string_query = document_id_string_query_value.(string)
	} else {
		render.Render(w, r, ErrNotFound)
		return
	}
	page_value := r.Context().Value("page")
	if page_value != nil {
		page = page_value.(int64)
	} else {
		render.Render(w, r, ErrNotFound)
		return
	}
	amount_docs_value := r.Context().Value("amountDocs")
	if amount_docs_value != nil {
		amount_docs = amount_docs_value.(int64)
	} else {
		render.Render(w, r, ErrNotFound)
		return
	}

	zincSearchResponse, err := zincSearchGetDocument(document_id_string_query, page, amount_docs)
	if err != nil {
		render.Render(w, r, ErrRender(err))
	}

	if err := render.Render(w, r, NewDocumentResponse(zincSearchResponse)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

}

func NewDocumentResponse(query_zinc_search_response *responseZincSearch) *payloadResponseZincSearch {
	document_list_hits := query_zinc_search_response.Hits.Hits
	var documentList []Document
	index := ""
	countDocs := query_zinc_search_response.Hits.Total.Value
	for _, documentHit_element := range document_list_hits {
		index = documentHit_element.Index_
		highlight := documentHit_element.Highlight.Content
		document_i := Document{
			Message_ID:                documentHit_element.Source_.Message_ID,
			Date:                      documentHit_element.Source_.Date,
			From:                      documentHit_element.Source_.From,
			To:                        documentHit_element.Source_.To,
			Subject:                   documentHit_element.Source_.Subject,
			Cc:                        documentHit_element.Source_.Cc,
			Mime_version:              documentHit_element.Source_.Mime_version,
			Content_Type:              documentHit_element.Source_.Content_Type,
			Content_Transfer_Encoding: documentHit_element.Source_.Content_Transfer_Encoding,
			Bcc:                       documentHit_element.Source_.Bcc,
			X_from:                    documentHit_element.Source_.X_from,
			X_to:                      documentHit_element.Source_.X_to,
			X_cc:                      documentHit_element.Source_.X_cc,
			X_bcc:                     documentHit_element.Source_.X_bcc,
			X_folder:                  documentHit_element.Source_.X_folder,
			X_origin:                  documentHit_element.Source_.X_origin,
			X_filename:                documentHit_element.Source_.X_filename,
			Content:                   documentHit_element.Source_.Content,
			Highlight:                 highlight,
		}
		documentList = append(documentList, document_i)
	}
	resp := &payloadResponseZincSearch{Documents: documentList, Index: index, Count: countDocs}
	return resp
}

//--
// Error response payloads & renderers
//--

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

// --
// Data model objects and persistence mocks:
// --
type Document struct {
	Message_ID                string   `json:"Message_ID"`
	Date                      string   `json:"Date"`
	From                      string   `json:"From"`
	To                        []string `json:"To"`
	Subject                   string   `json:"Subject"`
	Cc                        []string `json:"cc"`
	Mime_version              string   `json:"Mime_version"`
	Content_Type              string   `json:"content_Type"`
	Content_Transfer_Encoding string   `json:"content_Transfer_Encoding"`
	Bcc                       []string `json:"Bcc"`
	X_from                    string   `json:"X_from"`
	X_to                      []string `json:"X_to"`
	X_cc                      []string `json:"X_cc"`
	X_bcc                     []string `json:"X_bcc"`
	X_folder                  string   `json:"X_folder"`
	X_origin                  string   `json:"X_origin"`
	X_filename                string   `json:"X_filename"`
	Content                   string   `json:"Content"`
	Highlight                 []string `json:"Highlight"`
}

type Highlight struct {
	Content []string `json:"Content"`
}

type document_query_data struct {
	Index_    string    `json:"_index"`
	Type_     string    `json:"_type"`
	ID_       string    `json:"_id"`
	Score_    float64   `json:"_score"`
	TimeStamp string    `json:"@timestamp"`
	Source_   Document  `json:"_source"`
	Highlight Highlight `json:"highlight"`
}

type total_s struct {
	Value int `json:"value"`
}
type Hits struct {
	Total     total_s               `json:"total"`
	Max_score float64               `json:"max_score"`
	Hits      []document_query_data `json:"hits"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type responseZincSearch struct {
	Took      int    `json:"took"`
	Timed_out bool   `json:"timed_out"`
	Shards    Shards `json:"_shards"`
	Hits      Hits   `json:"hits"`
}

type payloadResponseZincSearch struct {
	Documents []Document `json:"documents"`
	Index     string     `json:"index"`
	Count     int        `json:"count"`
}

func (pRZS *payloadResponseZincSearch) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (pRZS *payloadResponseZincSearch) Bind(r *http.Request) error {
	return nil
}

func zincSearchGetDocument(id string, page int64, amountDocs int64) (*responseZincSearch, error) {
	var queryResponse *responseZincSearch = new(responseZincSearch)
	zincSearchServerIP := os.Getenv("IP")
	zincSearchServerPort := os.Getenv("Port")
	collectionName := os.Getenv("name_collecion")
	password := os.Getenv("password")
	user := os.Getenv("user")
	queryType := "match"
	if id == "" {
		queryType = "alldocuments"
	}
	queryPayLoad_match := fmt.Sprintf(`{
		"search_type":"match",
		"query":{
			"field":"_all",
			"start_time": "2020-12-25T15:08:48.777Z",
        	"end_time": "2024-12-28T16:08:48.777Z",
			"term":"%v"
		},
		"sort_fields":["@Date"],
		"from":%v,
		"max_results":%v,
		"highlight":{
            "pre_tags": ["<b>"],
            "post_tags": ["</b>"],
            "fields":{
                "Content":{},
                "Subject":{}
            }
        }
	}`, id, page, amountDocs)
	queryPayLoad_allDocs := fmt.Sprintf(`{
		"search_type":"alldocuments",
		"query":{
			"field":"_all",
			"start_time": "2020-12-25T15:08:48.777Z",
        	"end_time": "2023-12-28T16:08:48.777Z"
		},
		"sort_fields":["@Date"],
		"from":%v,
		"max_results":%v
	}`, page, amountDocs)
	var queryPayLoad string
	if queryType == "match" {
		queryPayLoad = queryPayLoad_match
	} else {
		queryPayLoad = queryPayLoad_allDocs
	}
	fmt.Println(queryPayLoad)
	urlBase := fmt.Sprintf("http://%v:%v/api/%v/_search", zincSearchServerIP, zincSearchServerPort, collectionName)
	req, err := http.NewRequest("POST", urlBase, strings.NewReader(queryPayLoad))
	if err != nil {
		return queryResponse, errors.New("It was not possible to stablish a connection to ZincSearchServer")
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return queryResponse, errors.New("Connection was stablished but data can not be retrieved")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return queryResponse, errors.New("JSON response can not be parsed from response request body")
	}
	responseString := string(body)
	errJson := json.Unmarshal([]byte(responseString), queryResponse)
	if errJson != nil {
		fmt.Println("json unmarshall error:")
		fmt.Println(errJson)
		return queryResponse, errors.New("JSON response can not be parse to response zinc search structure")
	}
	return queryResponse, nil
}
