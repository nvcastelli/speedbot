package main

import (
   "fmt"
   "time"
   "os"

   st "github.com/showwin/speedtest-go/speedtest"
)

type pingSet struct {
   iteration float64
   download float64
   upload float64
   latency time.Duration
   currentTime time.Time
}

func newPingSet(iteration float64, download float64, upload float64, latency time.Duration) pingSet {
   p := pingSet{iteration: iteration}
   p.download = download
   p.upload = upload
   p.latency = latency
   p.currentTime = time.Now()
   return p
}

func getCurrentTime (currentPingSet pingSet) time.Time {
   return currentPingSet.currentTime
}

func main(){
   fmt.Println("Will run iterations over the speed bot program to get an average internet speed test while the process is running");
   
   scanLengthString := os.Args[1]

   timeLength, _ := time.ParseDuration(scanLengthString)

   var count float64  = 0
   var sum float64 = 0
   var downloadSpeeds []pingSet


   for stay, timeout := true, time.After(timeLength); stay; {
      fmt.Println(time.Now());
      select {
         case <-timeout:
            stay = false
         default:
            user, _ := st.FetchUserInfo()

         	serverList, _ := st.FetchServers(user)
         	targets, _ := serverList.FindServer([]int{})

         	for _, s := range targets {
         		s.PingTest()
         		s.DownloadTest()
         		s.UploadTest()
               count++

         		fmt.Printf("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
         		downloadSpeeds = append(downloadSpeeds, newPingSet(count, s.DLSpeed, s.ULSpeed, s.Latency))
         	}

      }

   }

   for _, value := range downloadSpeeds {
      sum += value.download
   }

   fmt.Println("The list of speeds are: ")
   fmt.Println(downloadSpeeds)

   fmt.Println("The number of PINGs ran are:")
   fmt.Println(count)

   averageSpeeds := sum/count
   fmt.Println("The download speed average is:")
   fmt.Println(averageSpeeds)

}
