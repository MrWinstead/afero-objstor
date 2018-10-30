# Known Issues

## File Open, File Delete Race Condition

Consider two nodes each using this libary. Node A calls `Open` on key `/myfile`
and at the same time, node B calls `Remove` on the same file. Because it takes
two calls to object storage to open a file (fetch attributes, load into local
storage), it is possible to only half of node A's object storage read operations
to succeed.

The library handles this in that `Open` must be able to fetch both file
attributes and file contents to succeed