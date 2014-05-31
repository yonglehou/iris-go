// Copyright (c) 2014 Project Iris. All rights reserved.
//
// The current language binding is an official support library of the Iris
// cloud messaging framework, and as such, the same licensing terms apply.
// For details please see http://iris.karalabe.com/downloads#License

//

package iris

import "runtime"

//
type ServiceLimits struct {
	BroadcastThreads int // Broadcast handlers to execute concurrently
	BroadcastMemory  int // Memory allowance for pending broadcasts
	RequestThreads   int // Request handlers to execute concurrently
	RequestMemory    int // Memory allowance for pending requests
}

//
var defaultServiceLimits = ServiceLimits{
	BroadcastThreads: 4 * runtime.NumCPU(),
	BroadcastMemory:  64 * 1024 * 1024,
	RequestThreads:   4 * runtime.NumCPU(),
	RequestMemory:    64 * 1024 * 1024,
}
