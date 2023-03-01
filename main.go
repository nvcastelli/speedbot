package main

import (
   "fmt"
   "time"
   "os"

   st "github.com/showwin/speedtest-go/speedtest"
)

func main(){
   fmt.Println("Will run iterations over the speed bot program to get an average internet speed test while the process is running");
   
   scanLengthString := os.Args[1]

   timeLength, _ := time.ParseDuration(scanLengthString)

   //scanLengthInt, err := strconv.Atoi(scanLengthString)

   //d := time.Minute * scanLengthInt
   //scanLengthSeconds := int(d / time.Second)


   var downloadSpeeds []float64

   for stay, timeout := true, time.After(timeLength); stay; {
    fmt.Println(time.Now());
    select {
    case <-timeout:
        stay = false
	fmt.Println("+++I'm in timeout case?+++")
    default:
	 fmt.Println("---I'm in default case?---")
    


   user, _ := st.FetchUserInfo()
	// Get a list of servers near a specified location
	// user.SetLocationByCity("Tokyo")
	// user.SetLocation("Osaka", 34.6952, 135.5006)

	serverList, _ := st.FetchServers(user)
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(false)
		s.UploadTest(false)

		fmt.Printf("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
		downloadSpeeds = append(downloadSpeeds, s.DLSpeed)
	}

   fmt.Println("The list of speeds are: ")
   fmt.Println(downloadSpeeds)

      }

   }

}
