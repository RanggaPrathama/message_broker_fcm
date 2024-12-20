package response

type GlobalResponse struct{
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}