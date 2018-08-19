/*
* @Author: Ximidar
* @Date:   2018-08-18 19:54:28
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-08-19 14:00:48
 */
/********** CLI Command **********/
package cli_tool

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/ximidar/Z_Builder/package_assembler/path_builder"
)

var rootCmd = &cobra.Command{
	Use:   "Z_Builder",
	Short: "Build object packages for Z_Builder",
	Long:  `This Program is for building packages for the Z_Builder at MWERX.`,
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
	build.PersistentFlags().StringVar(&process_folder_path,
		"folder",
		"",
		"Folder with .png files")
	build.PersistentFlags().IntSliceVarP(&bake_in,
		"bake_in",
		"b",
		[]int{10, 5000},
		"2 integers that define the bake time for the first few layers. [layers, bake_time]")
	build.PersistentFlags().IntVarP(&layer_burn,
		"layer_burn",
		"l",
		1000,
		"Use this variable to define the burn time for each layer")
}

var process_folder_path string
var bake_in []int
var layer_burn int
var build = &cobra.Command{
	Use:   "build",
	Short: "This will add three text files to the supplied folder",
	Long: `This will parse the supplied folder of pngs into a usable
  		  build folder for the Z_Builder at MWERX`,
	Run: func(cmd *cobra.Command, args []string) {

		if process_folder_path == "" {
			cmd.Help()
			os.Exit(1)
		}

		if len(bake_in) > 2 || len(bake_in) < 2 {
			mess := fmt.Sprintf("Bake in command is not the correct length, Make sure to only use 2 ints. Example: \"-b 10,5000\"")
			fmt.Println(mess)
			os.Exit(1)
		}

		pb := path_builder.NewPathBuilder()
		pb.First_x_Layers = bake_in[0]
		pb.First_x_bake_time = bake_in[1]
		pb.Regular_bake_time = layer_burn
		err := pb.Open_Folder(process_folder_path)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Finished")
		os.Exit(0)
		fmt.Println(bake_in)
		fmt.Println(layer_burn)
		fmt.Println(process_folder_path)
		os.Exit(0)
	},
}
