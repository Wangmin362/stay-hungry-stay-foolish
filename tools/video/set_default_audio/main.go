package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Stream struct {
	Streams []struct {
		Index              int    `json:"index"`
		CodecName          string `json:"codec_name"`
		CodecLongName      string `json:"codec_long_name"`
		Profile            string `json:"profile,omitempty"`
		CodecType          string `json:"codec_type"`
		CodecTagString     string `json:"codec_tag_string"`
		CodecTag           string `json:"codec_tag"`
		Width              int    `json:"width,omitempty"`
		Height             int    `json:"height,omitempty"`
		CodedWidth         int    `json:"coded_width,omitempty"`
		CodedHeight        int    `json:"coded_height,omitempty"`
		ClosedCaptions     int    `json:"closed_captions,omitempty"`
		FilmGrain          int    `json:"film_grain,omitempty"`
		HasBFrames         int    `json:"has_b_frames,omitempty"`
		SampleAspectRatio  string `json:"sample_aspect_ratio,omitempty"`
		DisplayAspectRatio string `json:"display_aspect_ratio,omitempty"`
		PixFmt             string `json:"pix_fmt,omitempty"`
		Level              int    `json:"level,omitempty"`
		ColorRange         string `json:"color_range,omitempty"`
		ColorPrimaries     string `json:"color_primaries,omitempty"`
		ChromaLocation     string `json:"chroma_location,omitempty"`
		Refs               int    `json:"refs,omitempty"`
		RFrameRate         string `json:"r_frame_rate"`
		AvgFrameRate       string `json:"avg_frame_rate"`
		TimeBase           string `json:"time_base"`
		StartPts           int    `json:"start_pts"`
		StartTime          string `json:"start_time"`
		ExtradataSize      int    `json:"extradata_size,omitempty"`
		Disposition        struct {
			Default         int `json:"default"`
			Dub             int `json:"dub"`
			Original        int `json:"original"`
			Comment         int `json:"comment"`
			Lyrics          int `json:"lyrics"`
			Karaoke         int `json:"karaoke"`
			Forced          int `json:"forced"`
			HearingImpaired int `json:"hearing_impaired"`
			VisualImpaired  int `json:"visual_impaired"`
			CleanEffects    int `json:"clean_effects"`
			AttachedPic     int `json:"attached_pic"`
			TimedThumbnails int `json:"timed_thumbnails"`
			NonDiegetic     int `json:"non_diegetic"`
			Captions        int `json:"captions"`
			Descriptions    int `json:"descriptions"`
			Metadata        int `json:"metadata"`
			Dependent       int `json:"dependent"`
			StillImage      int `json:"still_image"`
		} `json:"disposition"`
		Tags struct {
			Language                    string `json:"language"`
			BPS                         string `json:"BPS"`
			BPSEng                      string `json:"BPS-eng"`
			DURATION                    string `json:"DURATION"`
			DURATIONEng                 string `json:"DURATION-eng"`
			NUMBEROFFRAMES              string `json:"NUMBER_OF_FRAMES"`
			NUMBEROFFRAMESEng           string `json:"NUMBER_OF_FRAMES-eng"`
			NUMBEROFBYTES               string `json:"NUMBER_OF_BYTES"`
			NUMBEROFBYTESEng            string `json:"NUMBER_OF_BYTES-eng"`
			STATISTICSWRITINGAPP        string `json:"_STATISTICS_WRITING_APP"`
			STATISTICSWRITINGAPPEng     string `json:"_STATISTICS_WRITING_APP-eng"`
			STATISTICSWRITINGDATEUTC    string `json:"_STATISTICS_WRITING_DATE_UTC"`
			STATISTICSWRITINGDATEUTCEng string `json:"_STATISTICS_WRITING_DATE_UTC-eng"`
			STATISTICSTAGS              string `json:"_STATISTICS_TAGS"`
			STATISTICSTAGSEng           string `json:"_STATISTICS_TAGS-eng"`
			Title                       string `json:"title,omitempty"`
		} `json:"tags"`
		SampleFmt      string `json:"sample_fmt,omitempty"`
		SampleRate     string `json:"sample_rate,omitempty"`
		Channels       int    `json:"channels,omitempty"`
		ChannelLayout  string `json:"channel_layout,omitempty"`
		BitsPerSample  int    `json:"bits_per_sample,omitempty"`
		InitialPadding int    `json:"initial_padding,omitempty"`
		BitRate        string `json:"bit_rate,omitempty"`
		DurationTs     int    `json:"duration_ts,omitempty"`
		Duration       string `json:"duration,omitempty"`
	} `json:"streams"`
}

func main() {
	for {
		setDefaultAudio("E:\\")
		time.Sleep(time.Minute)
	}
}

// 探测视频文件的音轨
// ffprobe -v quiet -print_format json -show_streams input_video.mp4
// 删除多余的音轨：只保留国语音轨
// ffmpeg -i input.mkv -map 0:0 -map 0:2 -c copy -disposition:a:0 default -y output4.mp4
func setDefaultAudio(vDir string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	filepath.Walk(vDir, func(p string, info fs.FileInfo, err error) error {
		if vDir == p {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(info.Name())
		switch ext {
		case ".part", ".downloading", ".xltd", ".zip":
			//fmt.Println(p, ", 正在下载！！！")
			return nil
		}

		cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_streams", p)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return nil
		}

		var stream Stream
		err = json.Unmarshal(output, &stream)
		if err != nil {
			return nil
		}
		if len(stream.Streams) <= 2 {
			return nil // 说明只有一个音轨，无法切换
		}
		stm := stream.Streams
		var chineseIdx int
		for idx, s := range stm {
			if s.CodecType == "audio" && (s.Tags.Title == "国语" || s.Tags.Title == "mandarin") {
				chineseIdx = idx
				break
			}
		}
		if chineseIdx != 0 { // 0肯定是视频
			// ffmpeg -i input.mkv -map 0:0 -map 0:2 -c copy -disposition:a:0 default -y output4.mp4
			base := filepath.Base(p)
			dir := filepath.Dir(p)
			newPath := fmt.Sprintf("%s\\new-%s", dir, base)
			cmd = exec.Command("ffmpeg", "-i", p, "-map", "0:0", "-map", fmt.Sprintf("0:%d", chineseIdx),
				"-c", "copy", "-disposition:a:0", "default", "-y", newPath)

			start := time.Now().Unix()
			fmt.Println(fmt.Sprintf("【%s】开始处理视频：%s", time.Now().Format("2006-01-02 15:04:05"), p))
			output, err = cmd.CombinedOutput()
			if err != nil {
				return nil
			}
			if err = os.Remove(p); err != nil {
				return nil
			}
			if err = os.Rename(newPath, p); err != nil {
				return nil
			}
			total := (time.Now().Unix() - start) / 60
			fmt.Println(fmt.Sprintf("【%s】耗时%d分钟，处理完成视频：%s", time.Now().Format("2006-01-02 15:04:05"), total, p))
			fmt.Println()
		}

		return nil
	})
}
