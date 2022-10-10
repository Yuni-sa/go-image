package main

import (
	"log"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func connectDB() *redis.Client {
	// defines redis connection
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	// simple ping / connection check
	pong, err := rdb.Ping().Result()
	log.Println(pong, err)
	return rdb
}

func sendToRedis(image Image) {
	data := map[string]interface{}{

		"imageName": image.Name,
		"imageUrl":  image.ImageUrl,
		"header":    "a",
		"size":      image.Size,
	}
	if err := rdb.HMSet(image.Name, data).Err(); err != nil {
		panic(err)
	}
}

//func toBase64(b []byte) string {
//	return base64.StdEncoding.EncodeToString(b)
//}
//
//func toBinary() []byte {
//	var size int64 = fileInfo.Size()
//	bytes := make([]byte, size)
//
//	// read file into bytes
//	buffer := bufio.NewReader(file)
//	_, err = buffer.Read(bytes)
//}
//
//func encodeAndStore(image Image) {
//
//	// Read the entire file into a byte slice
//	bytes, err := os.ReadFile(fmt.Sprintf("./images/%s", image.Name))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var base64Encoding string
//
//	// Determine the content type of the image file
//	mimeType := http.DetectContentType(bytes)
//
//	// Prepend the appropriate URI scheme header depending
//	// on the MIME type
//	switch mimeType {
//	case "image/jpeg":
//		base64Encoding += "data:image/jpeg;base64,"
//	case "image/png":
//		base64Encoding += "data:image/png;base64,"
//	}
//
//	// Append the base64 encoded output
//	base64Encoding += toBase64(bytes)
//
//	image.ImageUrl = base64Encoding
//
//	// Print the full base64 representation of the image
//	sendToRedis(image)
//}
