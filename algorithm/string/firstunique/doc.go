package firstunique

// =====================================================================================================================
// Given a big string of characters, how to efficiently find the first unique character in it?
// =====================================================================================================================
// The efficient solution is to use character as an index in a count array.
// Traverse the given string and store index of first occurrence of every character, also store count of occurrences.
// Then traverse the count array and find the smallest index with count as 1.
