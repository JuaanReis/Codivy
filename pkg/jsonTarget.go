package pkg

import (
	"encoding/json"
	"bytes"
)

func ToJSON(data any) (*bytes.Buffer, error) {
    b, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }
    return bytes.NewBuffer(b), nil
}