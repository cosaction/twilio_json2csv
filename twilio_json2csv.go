package main

import (
   "encoding/json"
   "encoding/csv"
   "os"
   "log"
   "flag"
   "strconv"
   "strings"
   "reflect"
)

type TwilioLog struct {
   DateSent string `json:"date_sent"`
   DateUpdated string `json:"date_updated"`
   Price string `json:"price"`
   AccountSID string `json:"account_sid"`
   From string `json:"from"`
   SegmentCnt int `json:"num_segments"`
   DateCreated string `json:"date_created"`
   ErrCode int `json:"error_code"`
   Status string `json:"status"`
   Body string `json:"body"`
   SID string `json:"sid"`
   MsgSvcSID string `json:"messaging_service_sid"`
   To string `json:"to"`
   Direction string `json:"direction"`
   NumMedia int `json:"num_media"`
}

func writeOut(pFile *os.File, strOut string) {
   amt, err := pFile.WriteString(strOut)
   if (amt <= 0 || err != nil) {
      log.Fatalln("Failed to write string to file", err)
   }
}

func (structPtr *TwilioLog) GetFieldName(fieldPtr interface{}) (name string) {

    structElem := reflect.ValueOf(structPtr).Elem()
    fieldElem := reflect.ValueOf(fieldPtr).Elem()

    for i := 0; i < structElem.NumField(); i++ {
        fieldVal := structElem.Field(i)
        if fieldVal.Addr().Interface() == fieldElem.Addr().Interface() {
            return structElem.Type().Field(i).Name
        }
    }
    return
}

func main() {

   //proc cmdline args
   var inFileName, outFileName string
   flag.StringVar(&inFileName, "in", "twilio.json", "Input file -- Defaults to 'twilio.json' in the current directory")
   flag.StringVar(&outFileName, "out", "twilio.csv", "Output file -- Defaults to 'twilio.csv' in the current directory")
   var csvNoHeader = flag.Bool("noheader", false, "Do not write a CSV header -- Defaults to off (will write a CSV header)")

   flag.Parse()

   if len(inFileName) <= 0 {
      inFileName = "twilio.json"
   }

   if len(outFileName) <= 0 {
      outFileName = "twilio.csv"
   }

   //read json file
   jsonFile, fileErr := os.Open(inFileName)
   defer jsonFile.Close();
   if fileErr != nil {
      log.Fatalln("Couldn't open JSON file.", fileErr)
   }
   //jsonScan := bufio.NewScanner(jsonFile)
   decJSON := json.NewDecoder(jsonFile)

   //pull in the JSON objects with a reader
   var fullLog []TwilioLog
   for decJSON.More() {
      var logLine TwilioLog

      err := decJSON.Decode(&logLine)
      if err != nil {
			log.Fatal("Failed to decode JSON.", err)
		}

      //add to slice to write out later...
      fullLog = append(fullLog, logLine)
   }

   //write out CSV conversion
   csvFile, CSVErr := os.OpenFile(outFileName, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0644)
   defer csvFile.Close()
   if CSVErr != nil {
      log.Fatalln("Couldn't create CSV file.", CSVErr)
   }

   //now dump each line to CSV...
   csvOut := csv.NewWriter(csvFile)
   defer csvOut.Flush()

   //write out the header here, if enabled
   if (*csvNoHeader == false) {
      //we reflect to find the structure name so we don't have
      //to maintain a bunch of other strings...
      var csvHeadTxt = []string {
                        fullLog[0].GetFieldName(&fullLog[0].DateSent),
                        fullLog[0].GetFieldName(&fullLog[0].DateUpdated),
                        fullLog[0].GetFieldName(&fullLog[0].Price),
                        fullLog[0].GetFieldName(&fullLog[0].AccountSID),
                        fullLog[0].GetFieldName(&fullLog[0].From),
                        fullLog[0].GetFieldName(&fullLog[0].SegmentCnt),
                        fullLog[0].GetFieldName(&fullLog[0].DateCreated),
                        fullLog[0].GetFieldName(&fullLog[0].ErrCode),
                        fullLog[0].GetFieldName(&fullLog[0].Status),
                        fullLog[0].GetFieldName(&fullLog[0].Body),
                        fullLog[0].GetFieldName(&fullLog[0].SID),
                        fullLog[0].GetFieldName(&fullLog[0].MsgSvcSID),
                        fullLog[0].GetFieldName(&fullLog[0].To),
                        fullLog[0].GetFieldName(&fullLog[0].Direction),
                        fullLog[0].GetFieldName(&fullLog[0].NumMedia)}
      CSVErr = csvOut.Write(csvHeadTxt)
      if CSVErr != nil {
         log.Fatalln("Failed to write header to CSV", CSVErr)
      }
   }

   recCnt := len(fullLog)
   for i := 0; i < recCnt; i++ {
      var csvLine = []string {
                        fullLog[i].DateSent,
                        fullLog[i].DateUpdated,
                        fullLog[i].Price,
                        fullLog[i].AccountSID,
                        fullLog[i].From,
                        strconv.Itoa(fullLog[i].SegmentCnt),
                        fullLog[i].DateCreated,
                        strconv.Itoa(fullLog[i].ErrCode),
                        fullLog[i].Status,
                        strings.Replace(fullLog[i].Body, "\n", "\\n", -1),
                        fullLog[i].SID,
                        fullLog[i].MsgSvcSID,
                        fullLog[i].To,
                        fullLog[i].Direction,
                        strconv.Itoa(fullLog[i].NumMedia)}
      err := csvOut.Write(csvLine)
      if err != nil {
         log.Fatalln("Failed to write to CSV", err)
      }
   }
}
