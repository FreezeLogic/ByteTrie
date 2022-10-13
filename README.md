# ByteTrie

Implementation of the [Trie](https://en.wikipedia.org/wiki/Trie) algorithm for bytes in Go with memory optimization.

## Install
```
go get github.com/dghubble/trie
```
## Usage
First we need to initialize Trie
```go
trie := bytetrie.NewTrie()
```
When this is done, we can insert some data using the Insert method.
```go
trie.Insert([]byte{0x00, 0xFF})	
```
And finally we can find our data using the Search method.
```go
trie.Search([]byte{0x00, 0xFF})
```
It will return bool variable. True if byte sequence presented in Trie and False otherwise.
But keep in mind what Search is length sensitive.
If you Trie contain byte sequence '0x00, 0xFF, 0xFF', Search method will return false for '0x00, 0xFF'.

Full example with comments below.
```go
package main

import (
	"fmt"
	bytetrie "github.com/FreezeLogic/ByteTrie"
)

func main() {
  // initializing Trie
	trie := bytetrie.NewTrie()
  
  // Insert data into the Trie
	trie.Insert([]byte{0x00, 0xFF})			// 1
	trie.Insert([]byte{0x00, 0xFF, 0xFF})		// 2
	trie.Insert([]byte{0x00, 0xF0, 0xFF, 0xFD})	// 3
	trie.Insert([]byte{0x01, 0xF0, 0xFF})		// 4
	trie.Insert([]byte{0x01, 0xF0, 0xFD, 0xFA})	// 5
  
  // Search same Byte sequence with given the length.
	fmt.Println(trie.Search([]byte{0x00, 0xFF}))              // True - Same like 1
	fmt.Println(trie.Search([]byte{0x00, 0xF0}))              // False - Shorter than 3
	fmt.Println(trie.Search([]byte{0x00, 0xF0, 0xFF, 0xFD}))  // True - Same like 3
	fmt.Println(trie.Search([]byte{0x01, 0xF0, 0xFD}))        // False - Not represented
	fmt.Println(trie.Search([]byte{0x0F, 0xFF}))              // False - Not represented 
```

## License
This package is distributed under the MIT license found in the LICENSE file.
