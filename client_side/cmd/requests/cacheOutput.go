package requests

func CacheOutput() (string, interface{}, int64) {
	url := "http://localhost:8080/admin/cache/output"

	status, respBody, responseTime := sendRequest("GET", url, nil)

	return status, respBody, responseTime
}