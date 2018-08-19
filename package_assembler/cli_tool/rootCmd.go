/*
* @Author: Ximidar
* @Date:   2018-08-18 19:54:28
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-08-18 20:08:28
*/
/********** CLI Command **********/
package cli_tool

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"

	"github.com/ximidar/Z_Builder/package_assembler/path_builder"
)
var rootCmd = &cobra.Command{
  Use:   "Z_Builder",
  Short: "Build object packages for Z_Builder",
  Long: `This Program is for building packages for the Z_Builder at MWERX.`,
  Run: func(cmd *cobra.Command, args []string) {
  	if len(args) == 0 {
  		cmd.Help()
  		os.Exit(1)
  	} else {
  		fmt.Println("Z_Builder")
  		fmt.Println("Written By: Matt Pedler")
  	}
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  rootCmd.AddCommand(build)
  build.PersistentFlags().StringVar(&process_folder_path, "folder", "", "Folder with .png files")
  build.MarkFlagRequired("folder")
}

var process_folder_path string
var build = &cobra.Command{
  Use:   "build",
  Short: "This will add three text files to the supplied folder",
  Long:  `This will parse the supplied folder of pngs into a usable
  		  build folder for the Z_Builder at MWERX`,
  Run: func(cmd *cobra.Command, args []string) {

  	if process_folder_path == "" {
  		cmd.Help()
  		os.Exit(1)
  	}
  	pb := path_builder.NewPathBuilder()
  	err := pb.Open_Folder(process_folder_path)

  	if err != nil{
  		fmt.Println(err)
  		os.Exit(1)
  	}

  	fmt.Println("Finished")
  	os.Exit(0)
  },
}
