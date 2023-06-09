# Golang TODO Server

This is companion code for my GFY talk.

It's not a pinnacle of good golang code, it's a kitchen sink.

## Bugs

* [ ] -h runs the server still
* [ ] POST works on any url because we aren't http path matching
* [ ] setting -url on todo client new includes the flag in the todo name
* [ ] no mutex on the memory store :yolo:
* [ ] "/" is 200-ing any not found page.
* [ ] Todo Vue UI doesn't hit api (using local storage ATM)
