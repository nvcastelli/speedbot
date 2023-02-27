package main

import (
   "fmt"
   "time"
   
   st "github.com/showwin/speedtest-go/speedtest"
)

func main(){
   fmt.Println("Will run iterations over the speed bot program to get an average internet speed test while the process is running");
   
   for stay, timeout := true, time.After(time.Second); stay; {
    fmt.Println(time.Now());
    select {
    case <-timeout:
        stay = false
    default:
    }
}

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
	}
}
