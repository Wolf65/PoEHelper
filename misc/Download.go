package misc

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(filepath, filename, url string) (err error) {
	// Create the file
	out, err := os.Create(fmt.Sprintf("%s\\%s", filepath, filename))
	if err != nil {
		fmt.Println("create ", err)
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get ", err)
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("copy ", err)
		return err
	}

	return nil
}
