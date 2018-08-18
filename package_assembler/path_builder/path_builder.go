/*
* @Author: Ximidar
* @Date:   2018-08-11 14:10:01
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-08-18 15:58:33
*/
package path_builder

import (
	"fmt"
	"path"
	"os"
	"errors"
)

type PathBuilder struct{
	Image_Paths []string
}

func NewPathBuilder() *PathBuilder{
	pb := new(PathBuilder)

	return pb
}

// This function will open the supplied folder and process all the images inside of it.
func (pb *PathBuilder) Open_Folder(folderpath string) (error){

	// Check if path is valid
	err := pb.check_folder_integrity(folderpath)
	if err != nil{
		return err
	}

	// process images

	return nil
}

// This function will create the static files based off of the collected paths
func (pb *PathBuilder) Create_Static_Files() (error){
	return nil
}


/*
	<file = "C:\Program Files\Vanquish\$$RecentJob\ET ULTRA_(ET ULTRA)_[258.03 x 161.32]__SI 500 __[vd50_w250]__(12.03.12_16.44.52)\vds_0000#0.png" expose_time = 50000 add = 1>

	This function needs to return a string like the one above. Mostly this just needs to 
	add the name of the png to a string
*/
func (pb *PathBuilder) AlterPath(filepath string, expose_time string, add_number string) (string){

	// Get the basename
	filename := path.Base(filepath)

	// construct the string
	Z_path := fmt.Sprintf(`<file = "C:\Program Files\Vanquish\$$RecentJob\ET ULTRA_(ET ULTRA)_[258.03 x 161.32]__SI 500 __[vd50_w250]__(12.03.12_16.44.52)\%s" expose_time = %s add = %s>`,
		filename, expose_time, add_number)

	// return the string
	return Z_path

}


/*************** Helpers ***************/

func (pb PathBuilder) check_folder_integrity(folderpath string) (error){
	fi, err := os.Stat(folderpath)
    if os.IsNotExist(err) {
        mess := fmt.Sprintf("Could not get Stats of %s", folderpath)
        return errors.New(mess)
    }
    mode := fi.Mode()
    if mode.IsRegular(){
	    mess := fmt.Sprintf("Path is a file, not a directory: %s", folderpath)
	    return errors.New(mess)
    }
    return nil
}