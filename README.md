# verbose-engine
really basic music file standardiser in bash

### using
assumes you have `/artist/album/*.mp3` layout

give run permission with
``` chmod +x vergine ```

place in the artist folder

run with ``` ./vergine "artist name" "ext" ```

"artist name" should be the name you want infront of song name i.e `Rick Astley - Never Gonna Give You Up`

"ext" should be the file extension you want i.e `mp3`,`flac` or `opus`


### todo
* improve options
  * add -d option for removing
  * add -v option for verbose output
* improve media type detection
* add folder cleanup
* increase directory depth (potentially dangerous)
