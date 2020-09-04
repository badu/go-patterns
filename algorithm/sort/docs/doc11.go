package docs

// =====================================================================================================================
// Tail Recursion
// =====================================================================================================================
//
// What is tail recursion?
// ===========================================
// A recursive function is tail recursive when recursive call is the last thing executed by the function.
// For example the following C++ function print() is tail recursive.
// // An example of tail recursive function
//void print(int n)
//{
//    if (n < 0)  return;
//    cout << " " << n;
//
//    // The last executed statement is recursive call
//    print(n-1);
//}
//
// Why do we care?
// ===========================================
// The tail recursive functions are considered better than non tail
// recursive functions as tail-recursion can be optimized by compiler.
// The idea used by compilers to optimize tail-recursive functions is simple,
// since the recursive call is the last statement, there is nothing left to do in the current function,
// so saving the current function’s stack frame is of no use
//
// Tail Call Elimination
// ===========================================
// As said above a recursive function is tail recursive if recursive call is the last thing executed by the function.
// We also discussed that a tail recursive is better than non-tail recursive as
// tail-recursion can be optimized by modern compilers.
// Modern compiler basically do tail call elimination to optimize the tail recursive code.
//
// If we take a closer look at above function, we can remove the last call with goto.
// Below are examples of tail call elimination.
//
// QuickSort is also tail recursive (Note that MergeSort is not tail recursive,
// this is also one of the reason why QuickSort performs better)
//
// /* Tail recursive function for QuickSort
// arr[] --> Array to be sorted,
//  low  --> Starting index,
//  high  --> Ending index */
//void quickSort(int arr[], int low, int high)
//{
//    if (low < high)
//    {
//        /* pi is partitioning index, arr[p] is now
//           at right place */
//        int pi = partition(arr, low, high);
//
//        // Separately sort elements before
//        // partition and after partition
//        quickSort(arr, low, pi - 1);
//        quickSort(arr, pi + 1, high);
//    }
//}
//
// The above function can be replaced by following after tail call elimination.
// /* QuickSort after tail call elimination
//  arr[] --> Array to be sorted,
//  low  --> Starting index,
//  high  --> Ending index */
//void quickSort(int arr[], int low, int high)
//{
//start:
//    if (low < high)
//    {
//        /* pi is partitioning index, arr[p] is now
//           at right place */
//        int pi = partition(arr, low, high);
//
//        // Separately sort elements before
//        // partition and after partition
//        quickSort(arr, low, pi - 1);
//
//        // Update parameters of recursive call
//        // and replace recursive call with goto
//        low = pi+1;
//        high = high;
//        goto start;
//    }
//}
// Therefore job for compilers is to identify tail recursion,
// add a label at the beginning and update parameter(s) at the end followed by adding last goto statement.
//
// Function stack frame management in Tail Call Elimination
// ===========================================
// Recursion uses stack to keep track of function calls.
// With every function call, a new frame is pushed onto the stack which contains local variables and data of that call.
// Let’s say one stack frame requires O(1) i.e,
// constant memory space, then for N recursive call memory required would be O(N).
// Tail call elimination reduces space complexity of recursion from O(N) to O(1).
// As function call is eliminated, no new stack frames are created and the function is executed in constant memory space.
// It is possible for the function to execute in constant memory space,
// because in tail recursive function, there are no statements after call statement so preserving state and
// frame of parent function is not required.
// Child function is called and finishes immediately, it doesn’t have to return control back to the parent function.
// As no computation is performed on the returned value and no statements are left for execution,
// current frame can be modified as per the requirements of current function call.
// So,
// there is no need to preserve stack frames of previous function calls and function executes in constant memory space.
// This makes tail recursion faster and memory friendly.
