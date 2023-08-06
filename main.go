package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	for {
		doTomatoSet()
	}
}

func doTomatoSet() {
	doWork()
	doShortRest()
	doWork()
	doShortRest()
	doWork()
	doShortRest()
	doWork()
	doLongRest()
}

func doWork() {
	fmt.Println("work")

	work_time := 20 * time.Minute
	time.Sleep(work_time)

	play_alarm()
}

func doShortRest() {
	fmt.Println("short rest")

	short_rest_time := 5 * time.Minute
	time.Sleep(short_rest_time)

	play_alarm()
}

func doLongRest() {
	fmt.Println("long rest")

	long_rest_time := 15 * time.Minute
	time.Sleep(long_rest_time)

	play_alarm()
}

func getAlarmFile() string {
	return `C:\Windows\Media\Alarm01.wav`
}

func play(audio_file string) {
	name := "/mnt/c/Program Files/VideoLAN/VLC/vlc.exe"
	cmd := exec.Command(name, "--play-and-exit", "-I", "dummy", audio_file)
	cmd.Start()
}

func play_alarm() {
	var alarm_file = getAlarmFile()
	play(alarm_file)
}
