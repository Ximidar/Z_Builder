/*
* @Author: Ximidar
* @Date:   2018-08-11 14:10:11
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-08-18 16:01:18
*/

package path_builder_test

import (
	"testing"
	"os"
	"github.com/ximidar/Z_Builder/package_assembler/path_builder"
)

func Test_AlterPath(t *testing.T) {
	pb := path_builder.NewPathBuilder()

	test_path := pb.AlterPath("/test/this/item.png", "5000", "1")

	if test_path != `<file = "C:\Program Files\Vanquish\$$RecentJob\ET ULTRA_(ET ULTRA)_[258.03 x 161.32]__SI 500 __[vd50_w250]__(12.03.12_16.44.52)\item.png" expose_time = 5000 add = 1>`{
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