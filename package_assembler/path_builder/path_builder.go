/*
* @Author: Ximidar
* @Date:   2018-08-11 14:10:01
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-08-18 16:34:17
*/
package path_builder

import (
	"fmt"
	"path"
	"os"
	"errors"
	"path/filepath"
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
	err = filepath.Walk(folderpath, func(filepath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Unable to Access File: %v: %v\n", filepath, err)
			return err
		}
		
		if info.IsDir() {
			return nil //skip any directories. I only want the files in the directory
		}

		// Process image file
		new_path := pb.AlterPath(filepath)
		pb.Image_Paths = append(pb.Image_Paths, new_path)		
		
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", folderpath, err)
		return err
	}

	// Create the static files
	err = pb.Create_Static_Files()

	if err != nil{
		return err
	}

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
func (pb *PathBuilder) AlterPath(filepath string) (string){

	// Get the basename
	filename := path.Base(filepath)

	// Get expose time
	expose_time := "5000" // TODO add function to calculate expose time

	// construct the string
	Z_path := fmt.Sprintf(`<file = "C:\Program Files\Vanquish\$$RecentJob\ET ULTRA_(ET ULTRA)_[258.03 x 161.32]__SI 500 __[vd50_w250]__(12.03.12_16.44.52)\%s" expose_time = %s add = %s>`,
		filename, expose_time, "1")

	fmt.Printf("Altering %s to %s\n", filepath, Z_path)

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