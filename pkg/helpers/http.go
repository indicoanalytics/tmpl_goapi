package helpers

import (
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	Header        interface{}       `json:"header"`
	Proto         string            `json:"proto"`
	ContentLength int64             `json:"content_length"`
	Host          string            `json:"host"`
	RemoteAddr    string            `json:"remote_addr"`
	RemoteIP      string            `json:"ipaddress"`
	Method        string            `json:"method"`
	URL           string            `json:"route"`
	Body          interface{}       `json:"body"`
	QueryParams   url.Values        `json:"query_params"`
	Params        map[string]string `json:"params"`
}

func FromHTTPRequest(context *fiber.Ctx) *Request {
	queryParams, _ := url.ParseQuery(string(context.Request().URI().QueryString()))

	req := &Request{
		Header:        context.GetReqHeaders(),
		Proto:         context.Protocol(),
		ContentLength: int64(context.Request().Header.ContentLength()),
		Host:          context.Hostname(),
		RemoteAddr:    context.IP(),
		RemoteIP:      context.IP(),
		Method:        context.Method(),
		URL:           context.OriginalURL(),
		QueryParams:   queryParams,
		Params:        context.AllParams(),
		Body:          parseBody(context),
	}

	return req
}

func CreateResponse(context *fiber.Ctx, payload interface{}, status ...int) error {
	returnStatus := http.StatusOK
	if len(status) > 0 {
		returnStatus = status[0]
	}

	return context.Status(returnStatus).JSON(payload)
}

func parseBody(context *fiber.Ctx) interface{} {
	contentType := context.Get("Content-Type")

	if contentType == fiber.MIMEApplicationJSON {
		response := map[string]interface{}{}
		_ = Unmarshal(context.Body(), &response)

		return response
	}

	if contentType == fiber.MIMEMultipartForm {
		response, _ := context.MultipartForm()

		return response
	}

	if contentType == fiber.MIMEApplicationForm {
		response, _ := url.ParseQuery(string(context.Body()))

		return response
	}

	return string(context.Body())
}
