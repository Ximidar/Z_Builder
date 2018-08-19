# Package Assembler

This will take the png folder and turn it into a printable file

# Go Packages

- cobra cli
- more to come

# How to install and build
I use a program called lndir to link my source files to my gopath

so to link this program
```
mkdir {{ Full path to gopath}}/src/github.com/ximidar/Z_Builder
lndir {{ full path to Z_Builder }} {{ Full path to gopath}}/src/github.com/ximidar/Z_Builder
```

Then to build the executable
```
cd Z_Builder
make
```

The built binary can be found in Z_Builder/bin/