package vk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type H map[string]interface{}

type UploadResponse struct {
	Response         map[string]interface{}
	Error            string `json:"error"`
	ErrorDescription string `json:"error_descr"`
}

func Create(token, version string) (api *API) {
	api = &API{}
	api.init(token, version)

	return
}

func prepare(params map[string]interface{}) string {
	data := make([]string, 0)

	for key, value := range params {
		switch value := value.(type) {

		case string:
			data = append(data, fmt.Sprintf("%s=%s", key, value))

		case int, int32, int64:
			data = append(data, fmt.Sprintf("%s=%d", key, value))

		case float32, float64:
			data = append(data, fmt.Sprintf("%s=%.2f", key, value))

		case []int:
			dat := make([]string, 0)
			for _, v := range value {
				dat = append(dat, strconv.Itoa(v))
			}

			data = append(data, fmt.Sprintf("%s=%s", key, strings.Join(dat, ",")))

		case []float64:
			dat := make([]string, 0)
			for _, v := range value {
				dat = append(dat, fmt.Sprintf("%s=%s", key, strconv.FormatFloat(v, 'f', 6, 64)))
			}

			data = append(data, fmt.Sprintf("%s=%s", key, strings.Join(dat, ",")))

		case []string:
			data = append(data, fmt.Sprintf("%s=%s", key, strings.Join(value, ",")))

		}

	}

	return strings.Join(data, "&")
}

type API struct {
	token, version string
}

func (api *API) init(token string, version string) {
	api.token = token
	api.version = version
}

func (api *API) Upload(uploadServer, filename, fieldName string, file *os.File) UploadResponse {
	if file == nil {
		log.Println("File is empty")
		return UploadResponse{
			Error:            "no_file",
			ErrorDescription: "File is empty",
		}
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile(fieldName, filename)
	_, _ = io.Copy(part, file)
	_ = writer.Close()

	req, _ := http.NewRequest("POST", uploadServer, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return UploadResponse{
			Error:            "request_error",
			ErrorDescription: err.Error(),
		}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return UploadResponse{
			Error:            "read_error",
			ErrorDescription: err.Error(),
		}
	}
	resp.Body.Close()

	response := UploadResponse{}
	_ = json.Unmarshal(b, &response)
	_ = json.Unmarshal(b, &response.Response)

	return response
	fmt.Println("post return")

	fmt.Println(response)
	apiResponse, err := api.request("docs.save", H{
		"file": response.Response["file"],
	})
	if err != nil {
		log.Println(err)
		return UploadResponse{
			Error:            "request_error",
			ErrorDescription: err.Error(),
		}
	}
	fmt.Println(apiResponse)
	response.Response = apiResponse.Raw["response"].(map[string]interface{})

	return response
}

func (api *API) request(method string, params map[string]interface{}) (resp *Response, err error) {
	resp = &Response{}

	params["access_token"] = api.token
	params["v"] = api.version

	url := fmt.Sprintf("https://api.vk.com/method/%s?%s", method, prepare(params))
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &resp.Raw)
	if err != nil {
		return nil, err
	}
	_ = json.Unmarshal(body, resp)

	return resp, nil
}

func (api *API) Api(method string, params H) (resp *Response, err error) {
	return api.request(method, params)
}

func (api *API) ApiKey(method string, params H, key string) (resp *Response, err error) {
	resp = &Response{}

	params["access_token"] = key
	params["v"] = api.version

	url := fmt.Sprintf("https://api.vk.com/method/%s?%s", method, prepare(params))
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &resp.Raw)
	if err != nil {
		return nil, err
	}
	_ = json.Unmarshal(body, resp)

	return resp, nil
}
