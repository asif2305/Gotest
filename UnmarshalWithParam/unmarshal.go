package main

import (
	"encoding/hex"
	"fmt"
)

// Define your struct with optional fields
type MyStructWithOptionalFields struct {
	Field1 int
	Field2 *string // This field is optional
}

func main() {
	// Define your byte stream
	byteStream := "0123456789abcdef"

	// Decode byte stream from hex
	data, err := hex.DecodeString(byteStream)
	if err != nil {
		fmt.Println("Error decoding byte stream:", err)
		return
	}

	// Create an instance of your struct
	var myStruct MyStructWithOptionalFields

	// Unmarshal byte stream into struct
	err = unmarshal(data, &myStruct)
	if err != nil {
		fmt.Println("Error unmarshaling byte stream:", err)
		return
	}

	// Print the struct
	fmt.Printf("%+v\n", myStruct)
}

// Define unmarshal function
func unmarshal(data []byte, s *MyStructWithOptionalFields) error {
	// Implement your own unmarshal logic here
	// For the sake of example, let's assume we just set Field1
	s.Field1 = int(data[0])
	return nil
}
