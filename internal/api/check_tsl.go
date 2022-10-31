package api

import (
	"crypto/tls"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type DayStruct struct {
	Day string
}

func CheckTLS(url string) DayStruct {

	log.SetFlags(log.Lshortfile)

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	elm := strings.Split(url, "://")

	if elm[0] == "https" {
		e := strings.Split(elm[1], ":")
		if len(e) == 1 {
			url = e[0] + ":443"
		} else {
			url = e[0] + ":" + e[1]
		}
	} else {
		e := strings.Split(elm[0], ":")
		if len(e) == 1 {
			url = e[0] + ":443"
		} else {
			url = e[0] + ":" + e[1]
		}
	}

	conn, err := tls.Dial("tcp", url, conf)
	if err != nil {
		log.Println(err)
	}
	defer func(conn *tls.Conn) {
		err = conn.Close()
		if err != nil {

		}
	}(conn)
	//fmt.Println(conn.ConnectionState())

	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	issuer := conn.ConnectionState().PeerCertificates[0].Issuer
	fmt.Println(issuer, int64(expiry.Sub(time.Now()).Hours()/24), conn.ConnectionState().PeerCertificates[0].IssuingCertificateURL)
	res := strconv.Itoa(int(expiry.Sub(time.Now()).Hours() / 24))
	days := DayStruct{
		Day: res,
	}
	return days
}
