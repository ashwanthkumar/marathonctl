package client

import (
  "encoding/json"
  "errors"
  "github.com/parnurzeal/gorequest"
)

type Marathon struct{
  Url string
}

var httpClient = gorequest.New()

func handle(response gorequest.Response, body string, errs []error) (string, error) {
  if(response != nil) {
    if((response.StatusCode != 200 && response.StatusCode != 201) && body != "") {
      var errorResponse map[string]interface{}
      err := json.Unmarshal([]byte(body), &errorResponse)
      if err != nil {
        errs = append(errs, err)
      } else {
        errs = append(errs, errors.New(errorResponse["message"].(string)))
      }
    } else if(response.StatusCode != 200 && response.StatusCode != 201) {
      errs = append(errs, errors.New(response.Status))
    }
  }
  return body, combineErrors(errs)
}

func combineErrors(errs []error) error {
  if(len(errs) == 1) {
    return errs[0]
  } else if(len(errs) > 1) {
    msg := "Error(s):"
    for _, err := range errs {
      msg += " " + err.Error()
    }
    return errors.New(msg)
  } else {
    return nil
  }
}
