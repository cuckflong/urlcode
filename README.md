# urlcode
## Tool for URL encoding/decoding which supports pipe lining
### Well can't find a tool capable of this so guess I'll have to write it myself. Meme name BTW

URL decode will ignore invalid % encoding and continues on the rest.\
Specify how many times of encoding/decoding with -t\
Can recursively decode until not more changes with -r
```
echo "<img src=x onerror=alert()>" | urlcode -t 100 | urlcode -d -r
```
The command above will URL encode 100 times then recursively decode and finally output the original value
