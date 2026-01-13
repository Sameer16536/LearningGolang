package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var err error // Simulate an error variable

// The defer Keyword
// In Go, you will often open things (files, connections, locks) that MUST be closed.
//  If your function errors out in the middle, you might forget to close them, leading to memory leaks.
//  defer schedules a function call to run immediately before the function returns, no matter what happens

func process() {
	mu.Lock()
	defer mu.Unlock() // This WILL happen at the end, even if code below panics

	// ... complex logic ...
	if err != nil {
		return // Unlock happens here automatically
	}
	// More logic here
}

// 2. The "Closure Loop" Trap (The Interview Killer)
// This is the specific bug I mentioned earlier. Read this carefully.
// When you launch goroutines inside a for loop, the variable i is shared.

func closureLoop() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // By the time this runs, 'i' might already be 5!
		}()
	}
	// Output might be: 5, 5, 5, 5, 5

	// ✅ CORRECT Way (Shadowing):
	//  You must pass the variable into the function to create a copy.
	for i := 0; i < 5; i++ {
		go func(val int) { // Accept it as an argument
			fmt.Println(val)
		}(i) // Pass 'i' in here
	}
	// Output: 0, 1, 2, 3, 4 (random order)
}

// 3. The net Package
// You aren't using HTTP requests here. You are going lower level: TCP Sockets. net.Dial("tcp", "scanme.nmap.org:80") attempts a raw handshake.

// If it returns nil error -> OPEN.

// If it returns error -> CLOSED (or filtered).

func main() {
	process()
	closureLoop()
}
