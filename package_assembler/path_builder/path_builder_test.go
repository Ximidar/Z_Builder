/*
* @Author: Ximidar
* @Date:   2018-08-11 14:10:11
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-08-18 19:12:28
*/

package path_builder_test

import (
	"testing"
	"os"
	"github.com/ximidar/Z_Builder/package_assembler/path_builder"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"io"

)

func Test_AlterPath(t *testing.T) {
	pb := path_builder.NewPathBuilder()

	test_path := pb.AlterPath("/test/this/item.png")

	if test_path != `<file = "C:\Program Files\Vanquish\$$RecentJob\ET ULTRA_(ET ULTRA)_[258.03 x 161.32]__SI 500 __[vd50_w250]__(12.03.12_16.44.52)\item.png" expose_time = 5000 add = 1>` + "\n"{
		t.Fatal("AlterPath did not assemble the string correctly")
	}
}


func Test_valid_paths(t *testing.T){
	pb := path_builder.NewPathBuilder()

	err := pb.Open_Folder("/fake/path/")

	if err == nil{
		t.Fatal("Entered a fake path and it did not return with an error")
	}

	file, err := os.Create("/tmp/tempfile.txt")
	if err != nil{
		t.Fatal(err)
	}

	defer file.Close()

	err = pb.Open_Folder("/tmp/tempfile.txt")

	if err == nil{
		t.Fatal("Entered a file path and it did not return with an error")
	}
}

func Test_Image_Processing(t *testing.T){
	test_path, err := Make_Temp_Test()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(test_path)

	// Create path builder
	pb := path_builder.NewPathBuilder()

	// open test dir
	test_dir := test_path + "/vds_.slice/" //open the directory with pictures
	err = pb.Open_Folder(test_dir)

	if err != nil {
		t.Fatal("Valid Path failed to open")
	}
}

/********** Helpers **********/

// This function will take from the example_builds and make a test in /tmp/
func Make_Temp_Test() (string, error){

	path_to_builds := "/home/ximidar/workspace/Z_Builder/example_builds/End_Cap" // Change this path for your test env
	temp_test_path := "/tmp/temp_build/"

	//delete test path

	if _, err := os.Stat(temp_test_path); os.IsNotExist(err) {
		// make temp dir
		mkdirerr := os.MkdirAll(temp_test_path, os.ModePerm)
		if mkdirerr != nil{
			return "", err
		}
	} else {
		// delete path
		rmerr := os.RemoveAll(temp_test_path)
		if rmerr != nil {
			return "", err
		}
	}
	

	//copy src dir to test dir
	err := CopyDir(path_to_builds, temp_test_path)
	if err != nil{
		return "",err
	}

	// Delete build info



	file_to_delete := []string{temp_test_path + "/vds_.slice/Data_Processing_Info.txt",
							   temp_test_path + "/vds_.slice/Jobinfo.txt",
							   temp_test_path + "/vds_.slice/BuildList.txt",
							}

	for _, file := range file_to_delete {
		fmt.Printf("Removing file %v\n", file)
		os.Remove(file)
	}

	return temp_test_path, nil

}

// Open source copying tool
/* MIT License
 *
 * Copyright (c) 2017 Roland Singer [roland.singer@desertbit.com]
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

// CopyFile copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file. The file mode will be copied from the source and
// the copied data is synced/flushed to stable storage.
func CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

// CopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
// Symlinks are ignored and skipped.
func CopyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("destination already exists")
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}