package e2e_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func Test_Bookings_Should_Create_Booking_When_Parameters_Are_Correct(t *testing.T) {

	Create_New_Class(t)

    booking := Test_Booking{
        Name: "User who wants to book a class",
        Date: "2021-11-28",
    }
    
    jsonData, err := json.Marshal(booking)
    if err != nil {
        t.Fatal(err)
    }

    client := http.Client{}
    req , err := http.NewRequest("POST", URL + "bookings", bytes.NewBuffer(jsonData))

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

func Test_Bookings_Should_Receive_Bad_Request_When_Parameters_Are_Missing(t *testing.T) {

	Create_New_Class(t)

    booking := Test_Booking{
        Date: "2021-11-28",
    }
    
    jsonData, err := json.Marshal(booking)
    if err != nil {
        t.Fatal(err)
    }

    client := http.Client{}
    req , err := http.NewRequest("POST", URL + "bookings", bytes.NewBuffer(jsonData))

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

func Test_Bookings_Should_Receive_Not_Acceptable_When_No_Class_Available(t *testing.T) {

	Create_New_Class(t)

    booking := Test_Booking{
        Name: "User who wants to book a class",
        Date: "2021-11-25",
    }
    
    jsonData, err := json.Marshal(booking)
    if err != nil {
        t.Fatal(err)
    }

    client := http.Client{}
    req , err := http.NewRequest("POST", URL + "bookings", bytes.NewBuffer(jsonData))

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

    if res.Status != NOT_ACCEPTABLE_RESPONSE {
        t.Fatalf("Expected %s got %s", NOT_ACCEPTABLE_RESPONSE, res.Status)
    }
}

func Test_Bookings_Should_Receive_Bad_Request_When_Header_Is_Missing(t *testing.T) {

    Create_New_Class(t)

    booking := Test_Booking{
        Name: "User who wants to book a class",
        Date: "2021-11-28",
    }
    
    jsonData, err := json.Marshal(booking)
    if err != nil {
        t.Fatal(err)
    }

    client := http.Client{}
    req , err := http.NewRequest("POST", URL + "bookings", bytes.NewBuffer(jsonData))

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