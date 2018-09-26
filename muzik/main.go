
package main

import (
  "fmt"
	"strings"
	"os"
	"io/ioutil"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"


)


func validate_directory(dir string) bool {
	 _, err := os.Stat(dir)

		if os.IsNotExist(err) == false {
			 return true }
   return false
}

func iterate_over(d string) ([]string, error){
	files, err := ioutil.ReadDir(d)

    if err != nil {
        log.Fatal(err)
    }
		var mp3list []string
    for _, file := range files {
			s := strings.Split(file.Name(),"/")
			//fmt.Println(s)

			exten := strings.Split(s[len(s)-1],".")

			//fmt.Println(len(exten))
			if len(exten) > 1{
				//fmt.Println(len(exten))
				ext := exten[len(exten)-1]
				if ext == "mp3"{
				// fmt.Println(ext)
				 mp3filename := d + file.Name()
				 mp3list = append(mp3list,mp3filename)
			 }
			}


    }
		//fmt.Println(mp3list)
 return mp3list, err
}

func main() {
	var d string

	fmt.Println("please enter your directory path that you want to play your mp3 music files from  e.g.: '/Users/xyz/' \n")
	fmt.Scan(&d )
	fmt.Println( "you have entered the directory\n",d)



	validate := validate_directory(d)
	if validate == true {
		fmt.Println("The directory is valid")
    file_list, _ := iterate_over(d)
 		if len(file_list) > 0 {
				for _,file := range file_list {
						fmt.Println("Now playing",file)
						f, err := os.Open(file)
						if err != nil {
							log.Fatal(err)
						}


							s, format, _ := mp3.Decode(f)
							speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
							playing := make(chan struct{})
							speaker.Play(beep.Seq(s, beep.Callback(func() {
								close(playing)
								})))

								<-playing
							}
						} else {
							fmt.Println(" But there are no MP3 files present in the directory")
						}


					} else{
						fmt.Println("Please enter a valid directory")
					}
}
