package main

import (
	"fmt"
	dB "mobileapps/jobsserver/firebasedb"
	"os"
	"time"
)

func main() {
	dB.GetJobs()
}
func writeLogFile(str string) (err error) {
	os.Remove("log.txt")
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	t := time.Now()
	f.WriteString("Date Time : " + fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()) + "\n")
	if _, err = f.WriteString(str + "\n"); err != nil {
		//panic(err)
	}
	return err
}
