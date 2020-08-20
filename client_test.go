package go_rubbles

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var (
	responsePredictSuccessful []byte
	responsePredictError      []byte
	requestPredictSuccessful  []byte
)

func TestMain(m *testing.M) {
	var err error
	responsePredictSuccessful, err = ioutil.ReadFile("testdata/resp_predict_successful")
	if err != nil {
		log.Fatal(err)
	}
	responsePredictError, err = ioutil.ReadFile("testdata/resp_predict_error")
	if err != nil {
		log.Fatal(err)
	}
	requestPredictSuccessful, err = ioutil.ReadFile("testdata/req_predict_successful")
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestGetPredict_Successful(t *testing.T) {
	want := new(PredictResponse)
	err := json.Unmarshal(responsePredictSuccessful, want)
	assert.Nil(t, err)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		body, _ := ioutil.ReadAll(r.Body)

		assert.Equal(t, requestPredictSuccessful, body)

		_, err = w.Write(responsePredictSuccessful)
		assert.Nil(t, err)
	}))
	defer ts.Close()

	// ---
	cl, err := NewClient(ClientConfig{
		HTTPClientTimeout: 30 * time.Second,
		BaseURL:           ts.URL,
		Token:             "token",
	})
	assert.Nil(t, err)

	// ---
	products := []Product{
		{
			ProductId: "9000101411423",
			Dates: ProductDates{
				ShipmentDateFrom: "2020-09-12",
				ShipmentDateTo:   "2020-09-25",
				ShelfDateFrom:    "2020-09-20",
				ShelfDateTo:      "2020-09-25",
			},
			Parameters: ProductParameters{
				Client:      "Pyaterochka",
				ClientID:    "541",
				ClientType:  "ka",
				Type:        "Mega",
				Price:       "208.83",
				DiscountPpd: "0.47000000",
				DiscountOff: "0.413000",
				DiscountOn:  "0.08000000",
			},
		},
	}

	got, err := cl.GetPredict(context.Background(), products)
	assert.Nil(t, err)
	assert.Nil(t, got.Error)
	assert.Equal(t, want, got)
}

func TestGetPredict_Error(t *testing.T) {
	want := new(PredictResponse)
	err := json.Unmarshal(responsePredictError, want)
	assert.Nil(t, err)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		_, err = w.Write(responsePredictError)
		assert.Nil(t, err)
	}))
	defer ts.Close()

	// ---
	cl, err := NewClient(ClientConfig{
		HTTPClientTimeout: 30 * time.Second,
		BaseURL:           ts.URL,
		Token:             "token",
	})
	assert.Nil(t, err)

	// ---
	products := []Product{
		{
			ProductId: "9000101411423",
			Dates: ProductDates{
				ShipmentDateFrom: "2020-05-26",
				ShipmentDateTo:   "2020-06-01",
				ShelfDateFrom:    "2020-05-01",
				ShelfDateTo:      "2020-05-14",
			},
			Parameters: ProductParameters{
				Client:      "Pyaterochka",
				ClientID:    "214",
				Type:        "Mega",
				Price:       "208.83",
				DiscountPpd: "0.47000000",
				DiscountOff: "0.413000",
				DiscountOn:  "0.08000000",
			},
		},
	}

	got, err := cl.GetPredict(context.Background(), products)
	assert.Nil(t, err)
	assert.Nil(t, got.Result)
	assert.Equal(t, want, got)
}
