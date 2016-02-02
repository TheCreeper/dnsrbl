# fdm

fdm is a tool to help you manage your source fastdl server and keep it up to 
date with custom maps, materials, models, and sounds.

## Usage

Simply compile using the [go toolchain](https://golang.org/dl/) then specify
the game assets directory and the fastdl directory.

Build the fastdl directory.
```
fdm -in "~/winempires/empires/custom/content" -out "~/fastdl"
```

Build the fastdl directory and cleanup cruft.
```
fdm -cleanup -in "~/winempires/empires/custom/content" -out "~/fastdl"
```
