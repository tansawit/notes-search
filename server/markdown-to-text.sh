find /Users/sawit/Dropbox/projects/second-brain/content/docs/topic -type f -not -name "_index.md" -exec sh -c ' s=${0##*/} && s=${s%.md} && pandoc "${0}" -t plain -o notes/${s}.txt' {} \;
