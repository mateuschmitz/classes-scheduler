package e2e_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

type Test_Class struct {
    Name string `json:"name"`
    StartDate string `json:"start_date"`
    EndDate string `json:"end_date"`
    Capacity int `json:"capacity"`
}

type Test_Booking struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func Create_New_Class(t *testing.T) {
	class := Test_Class{
        Name: "Test Class",
        StartDate: "2021-11-28",
        EndDate: "2021-11-30",
        Capacity: 20,
    }
    
    jsonData, err := json.Marshal(class)
    if err != nil {
        t.Fatal(err)
    }

    client := http.Client{}
    req , err := http.NewRequest("POST", URL + "classes", bytes.NewBuffer(jsonData))

    if err != nil {
        t.Fatal(err)
    }
    
    req.Header = http.Header{
        "Content-Type": []string{"application/json"},
    }
    
    res , err := client.Do(req)
    if err != nil {
        t.Fatal(err)
    }

    if res.Status != SUCCESS_RESPONSE {
        t.Fatalf("Expected %s got %s", SUCCESS_RESPONSE, res.Status)
    }
}