package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/fmusayev/golang-excel/config"
	"github.com/fmusayev/golang-excel/types"
)

type API struct{}

func (api API) SendRequest(id string) (types.OutputData, error) {
	var output types.OutputData

	resp, err := http.Get(config.API + id)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		return output, errors.New("not-found")
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&output)
	return output, nil
}
