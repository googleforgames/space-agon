
package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
	"flag"
	"context"
	"encoding/json"
	"os"
	"strings"
	"reflect"
	"strconv"

	"cloud.google.com/go/spanner"
)

var (
	port = flag.Int("port", 7777, "TCP Port for Listener")
	enablePrint = flag.Bool("print", false, "Enable Print")
	enableDB = flag.Bool("db", false, "Enable Database")
)

type GameEvent struct {
	EventID   string      `json:"eventid"`
	EventType string      `json:"eventtype"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

func main() {
	flag.Parse()

    ln, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
    if err != nil {
        log.Fatal(err)
    }
    defer ln.Close()

	fmt.Printf("Listening on port: %v\n", *port)

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println(err)
            continue
        }

        go handleConn(conn)
    }
}

func handleConn(conn net.Conn) {
    defer conn.Close()

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
		// Get Event
        //text := scanner.Text()
		var ge GameEvent
        err := json.Unmarshal(scanner.Bytes(), &ge)
        if err != nil {
            fmt.Println("Error decoding payload:", err)
        } else if *enablePrint {
	        fmt.Println("Received event:", ge)
		}

		// Validate Payload
		if !ge.validate() {
			fmt.Println("Invalid event received")
			return
		}

		// Process Event
		pEvent := processEvent(ge)

		// Write to Database
		if *enableDB {
			writeToDB(pEvent)
		}

    }
    if err := scanner.Err(); err != nil {
        log.Println("Error reading from connection:", err)
    }

}

func (ge GameEvent) validate() bool {
	// Validate Payload
	return true
}

func processEvent(ge GameEvent) GameEvent {
	// Process Event
	return ge
}

func writeToDB(ge GameEvent) {
	ctx := context.Background()

	key_string, value_string := formatStruct(ge)
	err := spannerWriteDML(ctx, key_string, value_string)
	if err != nil {
		fmt.Printf("Error when writing to Spanner. %v\n", err)
	}

}

func spannerWriteDML(ctx context.Context, keyString, valueString string) error {

	gcpProjectId    := os.Getenv("GCP_PROJECT_ID")
	spannerInstance := os.Getenv("SPANNER_INSTANCE")
	spannerDatabase := os.Getenv("SPANNER_DATABASE")
	spannerTable    := os.Getenv("SPANNER_TABLE_GAME_TELEMETRY")

	connectionStr := fmt.Sprintf("projects/%v/instances/%v/databases/%v", gcpProjectId, spannerInstance, spannerDatabase)

	spannerClient, err := spanner.NewClient(ctx, connectionStr)
	if err != nil {
		return err
	}
	defer spannerClient.Close()

	// Generate DML
	dml := fmt.Sprintf("INSERT %v (%v) VALUES (%v)", spannerTable, keyString, valueString)
	fmt.Printf("dml: %v\n", dml)

	_, err = spannerClient.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{
			SQL: dml,
		}
		rowCount, err := txn.Update(ctx, stmt)
		if err != nil {
			return err
		}
		log.Printf("%d record(s) inserted.\n", rowCount)
		return err
	})
	return err

}

func formatStruct(s interface{}) (string, string) {
	// Use reflection to get the fields of the struct
	st := reflect.TypeOf(s)
	sv := reflect.ValueOf(s)

	structNames := []string{}
	structValues := []string{}

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		fieldValue := sv.FieldByName(field.Name)
		fieldValueType := fieldValue.Type().String()
		structFieldValue := fieldValue.Interface()

		// Convert the interface to string
		stringValue, ok := structFieldValue.(string)
		if !ok {
			fmt.Printf("Field Type: %v\n", fieldValueType)
			if fieldValueType == "int" {
				stringValue = strconv.Itoa(structFieldValue.(int))
			} else if fieldValueType == "float64" {
				stringValue = fmt.Sprintf("%f", structFieldValue)
			} else if fieldValueType == "bool" {
				stringValue = strconv.FormatBool(structFieldValue.(bool))
			}
		} else {
			stringValue = "\"" + stringValue + "\""
		}

		// Append items to list
		structNames = append(structNames, field.Name)
		structValues = append(structValues, stringValue)

	}

	keyString := strings.Join(structNames, ", ")
	valueString := strings.Join(structValues, ", ")
	return keyString, valueString
}
