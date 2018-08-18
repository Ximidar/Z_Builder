/*
* @Author: Ximidar
* @Date:   2018-08-11 14:10:01
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-08-18 15:24:57
*/
package path_builder

import (
	"fmt"
	"path"
)

type PathBuilder struct{
	Image_Paths []string
}

func NewPathBuilder() *PathBuilder{
	pb := new(PathBuilder)

	return pb
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