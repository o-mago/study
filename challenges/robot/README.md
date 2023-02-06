## About
Algorithm to move a robot on a single axis using left ("L") and right ("R") commands and repeating previous commands ("1", "2" ...) also

## Aproach
It will check for accepted commands, first it will identify if the movement is different from "L" and "R" (the allowed commands to really move).
If it's a different char that's not a number or a number to repeat itself, negative index or a future command, it will replace the command with "E" and skip to the next command
If it's a accepted number, it will replace the number with the real command it refers to and move the robot, adding 1 or -1 to the position
If it is "L" or "R", it will change the position directly

### Time complexity:
It has a O(n) time complexity since it only iterate through the slice once.

### Space Complexity:
It has a O(1) complexity since it only create some constant size variables to suport the algorithm.
