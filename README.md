# Go: Subtle Nil Pointer Dereference After Error Handling
This example demonstrates a common pitfall in Go error handling that can lead to a nil pointer dereference.

## The Problem
The `Calculate` function correctly handles the division-by-zero error. However, if the error isn't properly checked in the calling function (main), a nil pointer dereference could occur if one attempted to access result or other data dependent on successful execution of `Calculate`.