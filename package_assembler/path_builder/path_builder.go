/*
* @Author: Ximidar
* @Date:   2018-08-11 14:10:01
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-08-19 13:26:51
 */
package path_builder

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

/********** PathBuilder **********/

type PathBuilder struct {
	ImagePaths []string
	FolderPath string

	//bake in vars
	First_x_Layers    int
	First_x_bake_time int
	Regular_bake_time int
}

func NewPathBuilder() *PathBuilder {
	pb := new(PathBuilder)

	pb.First_x_Layers = 10
	pb.First_x_bake_time = 5000
	pb.Regular_bake_time = 1000

	return pb
}

// This function will open the supplied folder and process all the images inside of it.
func (pb *PathBuilder) Open_Folder(folderpath string) error {

	// Check if path is valid
	err := pb.check_folder_integrity(folderpath)
	if err != nil {
		return err
	}

	pb.FolderPath = folderpath
	image_counter := 0

	// process images
	err = filepath.Walk(pb.FolderPath, func(filepath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Unable to Access File: %v: %v\n", filepath, err)
			return err
		}

		if info.IsDir() {
			return nil //skip any directories. I only want the files in the directory
		}

		// Process image file
		var new_path string
		if image_counter < pb.First_x_Layers {
			new_path = pb.AlterPath(filepath, pb.First_x_bake_time)
		} else {
			new_path = pb.AlterPath(filepath, pb.Regular_bake_time)
		}

		pb.ImagePaths = append(pb.ImagePaths, new_path)
		image_counter += 1

		return nil
	})

	fmt.Printf("Processed %v images!\n", image_counter)

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", folderpath, err)
		return err
	}

	// Create the static files
	err = pb.Create_Static_Files()

	if err != nil {
		return err
	}

	return nil
}

// This function will create the static files based off of the collected paths
/*
Data_Processing_Info.txt
Jobinfo.txt
BuildList.txt
*/
func (pb *PathBuilder) Create_Static_Files() error {

	// Create Static files DPI and JI
	DPI_path := path.Clean(pb.FolderPath + "/Data_Processing_Info.txt")
	err := pb.create_file_at(DPI_path, data_processing_info)

	if err != nil {
		return err
	}

	JI_path := path.Clean(pb.FolderPath + "/Jobinfo.txt")
	err = pb.create_file_at(JI_path, job_info)

	if err != nil {
		return err
	}

	// Create Build List file
	BL_path := path.Clean(pb.FolderPath + "/BuildList.txt")
	err = pb.create_file_at(BL_path, pb.ImagePaths...)

	if err != nil {
		return err
	}

	return nil
}

/*
	<file = "C:\Program Files\Vanquish\$$RecentJob\ET ULTRA_(ET ULTRA)_[258.03 x 161.32]__SI 500 __[vd50_w250]__(12.03.12_16.44.52)\vds_0000#0.png" expose_time = 50000 add = 1>

	This function needs to return a string like the one above. Mostly this just needs to
	add the name of the png to a string
*/
func (pb *PathBuilder) AlterPath(filepath string, expose_time int) string {

	// Get the basename
	filename := path.Base(filepath)

	// construct the string
	Z_path := fmt.Sprintf(`<file = "C:\Program Files\Vanquish\$$RecentJob\ET ULTRA_(ET ULTRA)_[258.03 x 161.32]__SI 500 __[vd50_w250]__(12.03.12_16.44.52)\%s" expose_time = %v add = 1>`,
		filename, expose_time)
	Z_path += "\n"

	fmt.Printf("Altering %s to %s\n", filepath, Z_path)

	// return the string
	return Z_path

}

/*************** Helpers ***************/

func (pb PathBuilder) check_folder_integrity(folderpath string) error {
	fi, err := os.Stat(folderpath)
	if os.IsNotExist(err) {
		mess := fmt.Sprintf("Could not get Stats of %s", folderpath)
		return errors.New(mess)
	}
	mode := fi.Mode()
	if mode.IsRegular() {
		mess := fmt.Sprintf("Path is a file, not a directory: %s", folderpath)
		return errors.New(mess)
	}
	return nil
}

func (pb PathBuilder) create_file_at(dest string, file_data ...string) error {
	fmt.Printf("Creating File at %v\n", dest)
	file, err := os.Create(dest)

	if err != nil {
		return err
	}

	defer file.Close()

	buf_writer := bufio.NewWriter(file)

	for _, line := range file_data {
		expected_write := len(line)
		actual_write, err := buf_writer.WriteString(line)

		if err != nil {
			return err
		}
		if expected_write != actual_write {
			mess := fmt.Sprintf("Expected bytes: %v Actual bytes: %v", expected_write, actual_write)
			return errors.New(mess)
		}
	}

	buf_writer.Flush()
	return nil
}

// Variables for writing
const data_processing_info = `-------------- Data processing information sheet ------------

-------------------------------------------------------------

Machine settings:
 o Machine type       : ET ULTRA
 o Resolution (pixel) : 1920 x 1200
 o Platform size (mm) : 258.03 x 161.32

-------------------------------------------------------------

Material settings:
 o Material type      : SI 500 

-------------------------------------------------------------

Scene informations:
 o Scene dimensions:
    - x-Axis (mm): 169.86 [+1.00, +169.86]
    - y-Axis (mm): 124.64 [-0.64, +124.64]
    - z-Axis (mm): 172.9 [+0.00, +172.9]
 o Processed data:
    - Model   No.1: Kirstin Part.stl
    

-------------------------------------------------------------

Raster settings:
 o Voxel depth:
     -  50 µm in range [0.00 - 172.9) mm
 o Number of voxel data sets: 3458

-------------------------------------------------------------

Voxel conversion settings       :
 o Converter type               : 'DBS v2.7'
 o Buildfilter version          : '2.7 Ultra/Xede/Xtreme'
 o Plugin DLL version           : '2.7.1468.1101'
 o Buildfilter type          : 'Xede/Xtreme'
 o Number solid base plates     : 0
 o Height of burn-in range      : 4
 o Height of stardard range     : 3454
 o Perfactory Support widening (µm)       : 250
 o Perfactory Support base widening (µm)  : 1500
 o Perfactory Support base height (µm)     : 500
 o ERM module                    : Not activated
 o Number of active voxel      : 0
 o Number of billed voxel      : 187323290

-------------------------------------------------------------

`

const job_info = "VoxelDepth 0.001969"
