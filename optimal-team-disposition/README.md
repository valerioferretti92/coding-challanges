## Optimal team disposition, code challange by Code Gladiators

The program reads lines from stdin. The input files it reads is made of the following blocks:
- Number of test cases N
- List of N test cases
Each test case is a block of lines structured this way:
- Number of members of a team M
- Line listing M numbers separated by spaces corresponding to the strength of each memeber of our team
- Line listing M numbers separated by spaces corresponding to the strength of each memeber of enemy team

For each test case, the program must determine the maximum number of wins that our team can obatin when confronting with enemy team. It needs to find the best disposition of members of our team to maximize wins member wise. If the two member confronting each other has the same strength, enemy memeber wins. For example, the input provided is the following:  
 
1  
10  
3 6 7 5 3 5 6 2 9 1  
2 7 0 9 3 6 0 6 2 6  

the output must be:  

7

because, the optimal permutation of member of our team allow 7 wins member wise.

### How to compile and execute
`g++ team-win-counter.cpp -o team-win-counter`  
`cat original-test-cases/data1.txt | ./team-win-counter`

### How to generate new test files
If needed, new test files containing any number of test cases can be generated, through test-case-generator.cpp. The following shows how to compile it and execute it:  
 
`g++ test-case-generator.cpp -o test-case-generator`  
`./test-case generator #TEST_CASE [TEAM_SIZE] [STRENGTH_UPPER_LIMIT]`

Authomatically genberated test file is called agdata.txt and it is place in here. Originally provided test files are placed in the data folder. 
