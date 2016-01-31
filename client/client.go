package client

import (
  "errors"
  "github.com/parnurzeal/gorequest"
)

type Marathon struct{
  Url string
}

var request = gorequest.New()

func handle(response gorequest.Response, body string, errs []error) (string, error) {
  return string(body), combineErrors(errs)
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
