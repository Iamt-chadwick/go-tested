package main

import
(
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"time"
)

var config appConfig
var Status  = true
var errLog *log.Logger
var GeneralLogger *log.Logger
var ErrorLogger *log.Logger

func main(){

	config = readConf("./config.json")
	router := mux.NewRouter()

	e, err := os.OpenFile("./log/error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}

	ErrorLogger = log.New(e, "", log.Ldate|log.Ltime)
	ErrorLogger.SetOutput(&lumberjack.Logger{
		Filename:   "./log/error.log",
		MaxSize:    200,
		MaxBackups: 3,
		MaxAge:     28,
	})

	GeneralLogger = log.New(e, "", log.Ldate|log.Ltime)
	GeneralLogger.SetOutput(&lumberjack.Logger{
		Filename:   "./log/general.log",
		MaxSize:    200,
		MaxBackups: 3,
		MaxAge:     28,
	})

	server := &http.Server{
		Addr: config.Host + ":" + config.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	router.HandleFunc("/", index)
	router.HandleFunc("/start", start)
	router.HandleFunc("/stop", stop)
	fmt.Println("Server is running atm")
	log.Fatal(server.ListenAndServe())

}

func index(w http.ResponseWriter, r *http.Request) {
	sendSimpleResponse(w, true, "Testing this out")
}

func CheckTime() bool {
	currentTime := time.Now()
	if currentTime.Hour() >= 20 || currentTime.Hour() < 8 {

		return false

	}
	return true
}

func start(w http.ResponseWriter, r *http.Request){
	Status = true
	fmt.Println("Processes are kicking offf and voilaaa!!!")
	for{
		if CheckTime() && Status{
			go func() {
				fmt.Println("Hi, Toni right here!")
				GeneralLogger.Println("This worked fine, logging")
			}()
			time.Sleep(time.Duration(10000) * time.Millisecond)
		}else{
			ErrorLogger.Println("Its either too late or someone asked me to shut down.")
			time.Sleep(time.Duration(10000) * time.Millisecond)
		}
	}
}

func stop(w http.ResponseWriter, r *http.Request){
	Status = false
	fmt.Println("Shutting down.......! Good bye")
}