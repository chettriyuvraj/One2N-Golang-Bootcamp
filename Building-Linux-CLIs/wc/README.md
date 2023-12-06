# Readme
- The only difference from wc output is that the 'base-error' is printed in lowercase, i.e 'is a directory' instead of 'Is a directory'
- Keeping multi-flag output same as wc output so that test case solutions can be quickly obtained (commented out code for extra space)
- Story 4: No test cases added in the interest of time
- Story 5: Test cases added only for "-lwc" instead of combinations of "-l -w -c" in the interest of time

- Story 8:
    - TODO: Implement this
    - The program will currently process 1 file at a time, we can use goroutines for parallelism, not exactly to the order of 1 goroutine per file (very large number of files might land us into trouble), but maybe create 'n' number of goroutines and then use these 
    - Since it is using buffered I/O, it should be able to handle large files in memory (to be tested)
    - Yes, already reading file line-by-line
    - Load testing to be done