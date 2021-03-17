// 视频的转码 和 封面 截取等处理
package tools

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

var VideoNameChan chan string = make(chan string, 100)

func TransCodingVideo() { // 从 VideoNameChan 拿到视频路径, 判断是否为 h264 编码, 不是则转为 h264
	for {
		videoPath := <-VideoNameChan
		if isH264(videoPath) {
			continue
		} else {
			tarVideo := videoPath + "_.mp4"
			cmdStr := "ffmpeg -i " + videoPath + " -vcodec h264 " + tarVideo
			cmd := exec.Command("bash", "-c", cmdStr)
			log.Println("开始进行视频: " + videoPath + " 的转码")
			cmd.CombinedOutput()
			log.Println("视频: " + videoPath + " 转码成功")
			os.Remove(videoPath)
			log.Println("移除原来视频")
			os.Rename(tarVideo, videoPath)
			log.Println("重命名转码后的视频.")
		}
	}
}

func VideoCoverGet(videoPath string) { // 拿到一个视频的路径, 抽取第一帧当封面
	tarImg := strings.Replace(videoPath, ".mp4", ".jpg", -1)
	cmdStr := "ffmpeg -i " + videoPath + " -ss 1 -f image2 " + tarImg

	cmd := exec.Command("bash", "-c", cmdStr)
	log.Println("开始进行视频: " + videoPath + " 的封面截取")
	cmd.CombinedOutput()
	log.Println("视频: " + videoPath + " 封面截取成功")
}

func isH264(path string) bool { // 接收一个视频, 判断是否为 H264 格式
	name := path
	cmdStr := "ffprobe -i " + name + " -v quiet -print_format json -show_streams"
	cmd := exec.Command("bash", "-c", cmdStr)
	res, err := cmd.CombinedOutput()
	if err != nil {
		log.Panicln("视频编码查询错误.")
	}
	if strings.Contains(string(res), "h264") {
		log.Println("接收的视频格式是 h264, 不用进行转码")
		return true
	} else {
		log.Println("接收的视频格式错误, 开始进行转码")
		return false
	}
}
