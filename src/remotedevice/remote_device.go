package remotedevice

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/gorilla/websocket"
)

type frameStream struct {
	fps       float32
	Generator *interface{}
}

type Device struct {
	Connection  *websocket.Conn
	FrameStream frameStream
}

func CreateDefaultDevice() *Device {
	return &Device{
		Connection: nil,
		FrameStream: frameStream{
			fps:       24,
			Generator: nil,
		},
	}
}

func (d Device) disconnect() {

}

func (d Device) RemoteDeviceService() {

	fps := 24
	ticker := time.NewTicker(time.Second / time.Duration(fps))

	const dps float32 = 60 // Hue degres per second

	var curTime, lastTime time.Time
	var delta float32
	var data [3]byte

	var Hue float32 = 0

	lastTime = time.Now()
	for {
		//Prepare data

		curTime = <-ticker.C
		delta = float32(curTime.Sub(lastTime).Seconds())
		lastTime = curTime

		Hue += dps * delta
		if Hue > 360. {
			Hue -= 360
		}

		data[0], data[1], data[2] = HSV2RGB(Hue, 1, 1)

		//Send data
		err := d.Connection.WriteMessage(websocket.BinaryMessage, data[:])
		log.Print("Data:", data)
		if err != nil {
			log.Print("Unable to send:", err)
			ticker.Stop()
			break
		}
	}
}

// HSL2RGB is HSL to RGB converter
func HSL2RGB(Hue, Saturation, Lightness float32) (byte, byte, byte) {

	if Hue > 360 {
		Hue = float32(math.Mod(float64(Hue), 360))
	}
	if Saturation > 1 {
		Saturation = float32(math.Mod(float64(Saturation), 1))
	}
	if Lightness > 1 {
		Lightness = float32(math.Mod(float64(Lightness), 1))
	}

	var C float32 = ((1 - float32(math.Abs(float64(2*Lightness-1)))) * Saturation)
	var X float32 = C * (1 - float32(math.Abs(math.Mod(float64(Hue/60), 2)-1)))
	var m float32 = Lightness - C/2

	log.Print()

	var r, g, b float32
	var R, G, B byte

	if Hue < 180 {
		if Hue > 120 {
			r = 0
			g = C
			b = X
		} else if Hue > 60 {
			r = X
			g = C
			b = 0
		} else {
			r = C
			g = X
			b = 0
		}
	} else {
		if Hue < 240 {
			r = 0
			g = X
			b = C
		} else if Hue < 300 {
			r = X
			g = 0
			b = C
		} else {
			r = C
			g = 0
			b = X
		}
	}

	R = byte(math.Round(float64((r + m) * 255)))
	G = byte(math.Round(float64((g + m) * 255)))
	B = byte(math.Round(float64((b + m) * 255)))

	return R, G, B
}

// HSV2RGB is HSV to RGB converter
func HSV2RGB(Hue, Saturation, Value float32) (byte, byte, byte) {

	if Hue > 360 {
		Hue = float32(math.Mod(float64(Hue), 360))
	}
	if Saturation > 1 {
		Saturation = float32(math.Mod(float64(Saturation), 1))
	}
	if Value > 1 {
		Value = float32(math.Mod(float64(Value), 1))
	}

	var C float32 = Saturation * Value
	var X float32 = C * (1 - float32(math.Abs(math.Mod(float64(Hue/60), 2)-1)))
	var m float32 = Value - C

	var r, g, b float32
	var R, G, B byte

	if Hue < 180 {
		if Hue > 120 {
			r = 0
			g = C
			b = X
		} else if Hue > 60 {
			r = X
			g = C
			b = 0
		} else {
			r = C
			g = X
			b = 0
		}
	} else {
		if Hue < 240 {
			r = 0
			g = X
			b = C
		} else if Hue < 300 {
			r = X
			g = 0
			b = C
		} else {
			r = C
			g = 0
			b = X
		}
	}

	R = byte(math.Round(float64((r + m) * 255)))
	G = byte(math.Round(float64((g + m) * 255)))
	B = byte(math.Round(float64((b + m) * 255)))

	return R, G, B
}
