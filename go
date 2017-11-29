#!/bin/bash

a=$1 # name of the band
b=$2 # file ext to encode to

if [[ $3 ]]; then
  echo "[FATAL] please encase band name in quotes"
  exit
fi

for i in */; do
  cd "$i" || return

  for j in *; do
    e="${j##*.}" # file extension

    if [[ $e = "mp3" || $e = "flac" || $e = "opus" ]]; then
      echo "[INFO] file /$j"

      d=${j%.*} # filename with no extension
      f=$(echo ${j%%-*} | xargs) # chars up to first "-" with whitespace trimmed

      if [[ $a && "$f" != "$a" ]]; then # if var empty and name not correct
        d="$a -${j#*-}"
        d="${d%.*}"
        echo "[INFO] bad name found at '$i$j', renaming to '$a -${j#*-}'"
      elif [[ $e = "$b" ]]; then
        echo "[INFO] potential duplicate found at '$i$j' adding z"
        d="z$d"
      fi

      echo "[INFO] encoding $d.$b"
      ffmpeg -hide_banner -y -v 24 -i "$j" -c:a libmp3lame -q:a 5 -vn "$d.$b"

      if [[ ${d:0:1} = "z" ]]; then
        echo "[INFO] renaming '$i$d.$b' to '${d#*z}.$b'"
        mv "$i$d.$b" "${d#*z}.$b"
      fi

      # a=${j:1}
      # mv "$j" "$a"
    else
      echo "[WARN] skipping '$i$j'"
    fi
  done
  cd ..
done
#
# sleep 1
# picard .
