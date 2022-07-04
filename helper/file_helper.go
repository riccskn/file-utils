package helper

import (
	"io"
	"os"
	"strings"
)

type File struct {
	file string
}

func NewFile(destination string) *File {
	return &File{
		file: destination,
	}
}

func (receiver *File) Delete() error {
	err := os.Remove(receiver.file)

	return err
}

//use the same method to move the file to another directory
func (file *File) Rename(new_name string) error {

	split := strings.Split(file.file, "/")

	mount := ""

	for i := range split {

		if i != len(split)-1 {
			if i == 0 {

				mount = mount + split[i]

			} else {

				mount = mount + "/" + split[i]

			}
		}

	}

	err := os.Rename(file.file, mount+"/"+new_name)

	return err
}

func (file *File) Copy(destination string) error {

	in, err := os.Open(file.file)

	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()

}
