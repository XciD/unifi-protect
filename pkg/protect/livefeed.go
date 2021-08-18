package protect

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type liveFrameType = byte

const (
	keyFrameLiveFrameType liveFrameType = iota + 247
	codecLiveFrameType
	initSegmentFrameType
	moovFrameType
	moofFrameType
	videoFrameType
	audioFrameType
	mdatFrameType
	endSegmentFrameType
)

type LiveFeed struct {
	nvr        *NVR
	Events     chan []byte
	camera     string
	connection *websocket.Conn
	channel    int
}

func NewLiveFeed(nvr *NVR, camera string, channel int) *LiveFeed {
	return &LiveFeed{
		nvr:     nvr,
		camera:  camera,
		channel: channel,
		Events:  make(chan []byte),
	}
}

func (l *LiveFeed) PumpData() {
	getWsURL := fmt.Sprintf("/proxy/protect/api/ws/livestream?"+
		"allowPartialGOP&"+
		"camera=%s&"+
		fmt.Sprintf("channel=%d&", l.channel)+
		"chunkSize=1024&"+
		"extendedVideoMetadata&"+
		"fragmentDurationMillis=100&"+
		"progressive&"+
		"requestId=yx263bv5k&"+
		"type=fmp4", l.camera)

	urlResp := URL{}
	if err := l.nvr.Call(http.MethodGet, getWsURL, nil, &urlResp); err != nil {
		panic(err)
	}

	parsedURL, err := url.Parse(urlResp.URL)

	if err != nil {
		panic(err)
	}

	parsedURL.Scheme = "wss"
	logrus.Infof("connecting to %s", parsedURL.String())

	c, _, err := websocket.DefaultDialer.Dial(parsedURL.String(), nil)
	if err != nil {
		logrus.Fatalf("dial: %s", err.Error())
		return
	}

	l.connection = c

	var leftOver []byte
	var currentSegment RawSegment
	var moov []byte

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			// Send an empty message to warn the close
			l.Events <- []byte{}
			return
		}

		if leftOver != nil {
			message = append(leftOver, message...)
			leftOver = nil
		}

		offset := 0
		messageLength := len(message)
		for {
			if messageLength-offset < 4 {
				leftOver = message[offset:]
				break
			}

			offsetWithHeader := offset + 4
			header := message[offset:offsetWithHeader]

			if header[0] > endSegmentFrameType || header[0] < initSegmentFrameType {
				logrus.Errorf("Error, invalid header %d", header[0])
				break
			}

			length := ((int(header[1])<<8 | int(header[2])) << 8) | int(header[3])

			if messageLength < offset+length+4 {
				leftOver = message[offset:]
				break
			}

			body := message[offsetWithHeader : offsetWithHeader+length]

			switch header[0] {
			case keyFrameLiveFrameType:
				currentSegment.keyframes = body
			case codecLiveFrameType:
				logrus.Infof("Codec %s", string(body))
			case initSegmentFrameType:
				currentSegment = RawSegment{
					moof:      make([][]byte, 0),
					mdat:      make([][]byte, 0),
					video:     make([][]byte, 0),
					audio:     make([][]byte, 0),
					keyframes: nil,
				}
			case moovFrameType:
				moov = body
			case moofFrameType:
				currentSegment.moof = append(currentSegment.moof, body)
			case videoFrameType:
				currentSegment.video = append(currentSegment.video, body)
			case audioFrameType:
				currentSegment.audio = append(currentSegment.audio, body)
			case mdatFrameType:
				currentSegment.mdat = append(currentSegment.mdat, body)
			case endSegmentFrameType:
				l.Events <- currentSegment.Concat(moov)
				moov = nil
			}
			offset = offsetWithHeader + length
		}
	}
}

func (l *LiveFeed) Close() error {
	return l.connection.Close()
}

type RawSegment struct {
	moof      [][]byte
	mdat      [][]byte
	video     [][]byte
	audio     [][]byte
	keyframes []byte
}

func (r RawSegment) Concat(moov []byte) []byte {
	buffer := bytes.NewBuffer(moov)
	for _, i := range r.moof {
		_, _ = buffer.Write(i)
	}
	for _, i := range r.mdat {
		_, _ = buffer.Write(i)
	}
	for _, i := range r.video {
		_, _ = buffer.Write(i)
	}
	for _, i := range r.audio {
		_, _ = buffer.Write(i)
	}
	return buffer.Bytes()
}
