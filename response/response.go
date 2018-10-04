package response

import (
  "encoding/json"
)

type BasicResponse struct {
  Message string  `json:"message"`
  Code int        `json:"code"`
  Response string `json:"response"`
}

func JsonResponse(r *BasicResponse) string {
  j, _ := json.Marshal(r)
  return string(j)
}
