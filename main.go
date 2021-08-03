package main

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"net/http"
	"os/exec"
)

const version = "0.0.1"

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/duration", durationHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to sysinfo_server, version ", version,
		"\nAvailable endpoints: /version, /duration\n")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, version, "\n")
}

func durationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, duration(), "\n")
}

func duration() float64 {
	cmd := exec.Command("systemd-analyze", "time", "--no-pager")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(`(.*) \(firmware\) \+ (.*) \(loader\) \+ (.*) \(kernel\) \+ (.*) \(userspace\)`)
	s := out.String()

	m := re.FindStringSubmatch(s)

	k := parseDuration(m[3])
	u := parseDuration(m[4])

	return k + u
}

func parseDuration(s string) float64 {
	mre := regexp.MustCompile(`(\d+)min`)
	sre := regexp.MustCompile(`(\d+.\d+)s`)
	mm := mre.FindStringSubmatch(s)
	min := 0.0
	var err error
	if len(mm) == 2 {
		min, err = strconv.ParseFloat(mm[1], 10)
		if err != nil {
			log.Fatal("Cannot parse ", mm[1])
		}
	}
	sm := sre.FindStringSubmatch(s)
	sec, err := strconv.ParseFloat(sm[1], 10)
	if err != nil {
		log.Fatal("Cannot parse ", sm[1])
	}
	return min * 60 + sec
}
