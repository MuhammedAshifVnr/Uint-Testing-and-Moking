package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetuser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/user/:id", GetUser)

	testcases := []struct {
		Name         string
		Id           string
		expectedCode int
		expectedBody string
	}{
		{
			expectedCode: 200,
			expectedBody: `{"Name":"Ashif","Id":1}`,
			Name:         "valied test",
			Id:           "1",
		},
		{
			expectedCode: 200,
			expectedBody: `{"Name":"Nuhman","Id":2}`,
			Name:         "valid test 2",
			Id:           "2",
		},{
			expectedCode: 404,
			expectedBody: `"user not found"`,
			Name: "wrong id",
			Id: "6",
		},{
			expectedCode: 400,
			expectedBody: `"invalid user id"`,
			Name: "abc testing",
			Id: "abc",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {

			req, err := http.NewRequest(http.MethodGet, "/user/"+tc.Id, nil)
			if err != nil {
				t.Fatal("can't create request :", err)
			}
			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)
			assert.Equal(t, tc.expectedCode, rr.Code)
			assert.JSONEq(t, tc.expectedBody, rr.Body.String())
		})
	}
}
