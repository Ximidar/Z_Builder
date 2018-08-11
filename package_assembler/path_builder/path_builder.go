/*
* @Author: Ximidar
* @Date:   2018-08-11 14:10:01
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-08-11 14:15:32
*/
package path_builder

import (
	"fmt"
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
func (pb *PathBuilder) AlterPath(path string) (string, error){

	// Get the basename

	// construct the string

	// return the string

}