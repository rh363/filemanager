package filemanager

import (
	"bufio"
	"fmt"
	"os"
)

func WriteFile(path string, txt []string) error { // WRITE A NEW FILE IF IT DOESN'T EXIST, ELSE CREATE A NEW FILE, REQUIRE FILE PATH AND CONTENT, RETURN AN ERROR

	File, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	defer File.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	writer := bufio.NewWriter(File)

	for _, data := range txt {
		_, err = writer.WriteString(data)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	writer.Flush()
	return nil
}

func ReadFile(path string) ([]string, error) { // READ A FILE CONTENT, REQUIRE FILE PATH, RETURN THE CONTENT AND AN ERROR

	File, err := os.Open(path)
	defer File.Close()
	if err != nil {
		return nil, err
	}

	reader := bufio.NewScanner(File)
	reader.Split(bufio.ScanLines)
	var txt []string

	for reader.Scan() {
		if reader.Text() != "" && reader.Text() != " " {
			fmt.Println(reader.Text())
			txt = append(txt, reader.Text())
		}
	}

	return txt, nil
}
