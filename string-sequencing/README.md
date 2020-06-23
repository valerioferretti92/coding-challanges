## String sequencing

You are working on a project far Amazon Video and you need to cut films into scenes. To help streamline the creation of the final films, the team needs to develop an automated way of breaking up individual shots (short sequence from a particular camera angle) in a film into scenes (a sequence of shots). There is already an algorithm that breaks the film up into shots and labels them with a letter. Identical shots are labelled with the same letter. Write a function which will partition a sequence of shot labels into scenes so that no shot labels appear in more than one scene. The output should be the length of each scene.

### Input:
The input to the function/method consists of an argument inputList, a list of characters representing the sequence of shots.

### Output
Return a list of integers representing the length of each scene, in the order in which it appears in the given sequence of shots.

### Example
Input: inputList= ababcbacadefegdehijhklij]\
Output: 978\
Explanation: The first scene consists of the shots a, b, and c. The second scene consists of d, e, f, and g. Finally, the last scene consists of h, i, j, and k. The answer is 3, 7, 8 because a, b, and c only appear in the first 3 characters, then d, e, f, and g appear in the next 7 The final 8 characters consist entirely of h, i, j, and k.

### How to run
Go src files can either be run through just in time compilation with the following command:
```
go run string-sequencing.go ababcbacadefegdehijhklij
```
or they can be compile and run as in the following:
```
go build string-sequencing.go
./string-sequencing ababcbacadefegdehijhklij
```

