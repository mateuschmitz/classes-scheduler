package e2e_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func Test_Should_Create_Class_When_Parameters_Are_Correct(t *testing.T) {

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

func Test_Should_Receive_Bad_Request_When_Parameter_Is_Missing(t *testing.T) {

    class := Test_Class{
        Name: "Test Class",
        StartDate: "2021-11-28",
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

    if res.Status != BAD_REQUEST_RESPONSE {
        t.Fatalf("Expected %s got %s", BAD_REQUEST_RESPONSE, res.Status)
    }
}

func Test_Should_Receive_Bad_Request_When_Header_Is_Missing(t *testing.T) {

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
    
    res , err := client.Do(req)
    if err != nil {
        t.Fatal(err)
    }

    if res.Status != BAD_REQUEST_RESPONSE {
        t.Fatalf("Expected %s got %s", BAD_REQUEST_RESPONSE, res.Status)
    }
}
