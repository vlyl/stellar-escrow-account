package util

import "os"

func AppendToFile(outputFile, content string) error {
	f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	buf := []byte(content)
	_, err = f.Write(buf)
	if err != nil {
		return err
	}
	_ = f.Close()
	return err
}
