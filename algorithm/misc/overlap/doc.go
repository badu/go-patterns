package overlap

// =====================================================================================================================
// Find if two rectangles overlap
// =====================================================================================================================
// Given two rectangles, find if the given two rectangles overlap or not.
// Note that a rectangle can be represented by two coordinates, top left and bottom right. So mainly we are given following four coordinates.
// l1: Top Left coordinate of first rectangle.
// r1: Bottom Right coordinate of first rectangle.
// l2: Top Left coordinate of second rectangle.
// r2: Bottom Right coordinate of second rectangle.
// Refer image1
// We need to write a function bool doOverlap(l1, r1, l2, r2) that returns true if the two given rectangles overlap.
//
// Note : It may be assumed that the rectangles are parallel to the coordinate axis.
//
// One solution is to one by one pick all points of one rectangle and see if
// the point lies inside the other rectangle or not.
// This can be done using the algorithm discussed here:
// https://www.geeksforgeeks.org/how-to-check-if-a-given-point-lies-inside-a-polygon/.
//
// Following is a simpler approach. Two rectangles do not overlap if one of the following conditions is true.
// 1) One rectangle is above top edge of other rectangle.
// 2) One rectangle is on left side of left edge of other rectangle.
//
// Time Complexity
// ===========================================
//
//
// RUN INSTRUCTIONS:
// ===========================================
// func main() {
// l1 := overlap.Point{}
// r1 := overlap.Point{}
// l2 := overlap.Point{}
// r2 := overlap.Point{}
//
// l1.X = 0
// l1.Y = 10
//
// r1.X = 10
// r1.Y = 0
//
// l2.X = 5
// l2.Y = 5
//
// r2.X = 15
// r2.Y = 0
//
// if overlap.DoOverlap(l1, r1, l2, r2) {
//  fmt.Println("Rectangles Overlap")
// } else {
//  fmt.Println("Rectangles Don't Overlap")
// }
//}
// Output
// ===========================================
// Rectangles Overlap
