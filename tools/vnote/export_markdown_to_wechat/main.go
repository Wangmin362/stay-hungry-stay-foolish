package main

// 用于导出markdown到微信公众号要求的，格式。主要做了以下几个事情：
// 1、把阿里云的图片链接转为微信的，如果这个图片没有上传到微信，那么上传
// 2、把外链改为引用的方式，以明文URL贴在Markdown底部
// 3、把Markdown的[TOC]标记去除掉
// 4、尝试看看能不能把markdown直接通过在线工具，譬如https://markdown.com.cn/转为适合微信公众号的markdown样式
func main() {

}
