package handlers

import "io"

func ReadResponse(body io.Reader) ([]byte, error){
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
