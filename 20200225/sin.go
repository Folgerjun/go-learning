package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {

	fmt.Println("开始绘制正弦函数（Sin）图像")
	sinPic()
	fmt.Println("绘制正弦函数（Sin）图像成功")
}

// 绘制正弦函数（Sin）图像
func sinPic() {

	// 图片大小
	const size = 300
	// 根据给定大小创建灰度图
	// image.Rect 描述一个方形的两个定位点 (x1,y1) 和 (x2,y2)
	pic := image.NewGray(image.Rect(0, 0, size, size))

	// 遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// 填充为白色 像素的灰度设为 255 就是白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	// 从 0 到最大像素生成 x 坐标
	for x := 0; x < size; x++ {

		// 让 sin 的值的范围在 0~2Pi 之间
		s := float64(x) * 2 * math.Pi / size

		// sin 的幅度为一半的像素 向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2

		// 用黑色绘制 sin 轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}

	// 创建文件
	file, err := os.Create("sin.png")

	if err != nil {
		log.Fatal(err)
	}
	// 使用 png 格式将数据写入文件
	png.Encode(file, pic) //将image信息写入文件中

	// 关闭文件
	file.Close()
}
