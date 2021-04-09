/*
log oluşturmak için kullanılır. Örnek olarak bir hata logu oluşturuldu.
Unix sistemlerde /var/log/ altında görülebilir.

Panic kullanımıda görsterilmiştir.
*/

package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {

	sysLog, err := syslog.New(syslog.LOG_ALERT, "some program")

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Fatal(sysLog)

	//log.Panic(sysLog)

	fmt.Println("Will you see this?")
}
