package toh

// =====================================================================================================================
// Tower of Hanoi
// =====================================================================================================================
//
// Problem
// ===========================================
// Tower of Hanoi is a mathematical puzzle where we have three rods and n disks.
// The objective of the puzzle is to move the entire stack to another rod, obeying the following simple rules:
// 1. Only one disk can be moved at a time.
// 2. Each move consists of taking the upper disk from one of the stacks and placing it on top of another stack i.e. a disk can only be moved if it is the uppermost disk on a stack.
// 3. No disk may be placed on top of a smaller disk.
//
// Examples
// ===========================================
// Input : 2
// Output : Disk 1 moved from A to B
//          Disk 2 moved from A to C
//          Disk 1 moved from B to C
//
// Input : 3
// Output : Disk 1 moved from A to C
//          Disk 2 moved from A to B
//          Disk 1 moved from C to B
//          Disk 3 moved from A to C
//          Disk 1 moved from B to A
//          Disk 2 moved from B to C
//          Disk 1 moved from A to C
//
//
// For n disks, total 2^n -1 moves are required.
// eg: For 4 disks 2^4  -1 = 15 moves are required.
//
// For n disks, total 2^(n-1) function calls are made.
//
// eg: For 4 disks 2^(4-1) = 8 function calls are made.
//
// Time Complexity
// ===========================================
//
//
// RUN INSTRUCTIONS:
// ===========================================
// func main() {
// toh.TowerOfHanoi(3, "A", "C", "B")
//}
// Output
// ===========================================
// Move disk 1 from A to C
// Move disk 2 from A to B
// Move disk 1 from C to B
// Move disk 3 from A to C
// Move disk 1 from B to A
// Move disk 2 from B to C
// Move disk 1 from A to C
