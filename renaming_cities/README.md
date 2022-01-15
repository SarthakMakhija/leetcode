### Problem

Some cities are going to be renamed and accordingly name of their railway stations will also change. Changing the name of railway station should also result in changed station code. Railways have an idea that station code should be the shortest prefix out of all railway stations renamed prior to this. If some city has same name, then prefix will be the name with suffix as the count of occurence of that city prior to this and including this, seperated with spaces.

Given N renamed cities consisting of lowercase alphabets only. The task is to generate a station ID for all the stations.

```
Example 1:

Input:
N = 6
Cities[] = {shimla,safari,jammu,delhi,jammu,dehradun}
Output:
s
sa
j
d
jammu 2
deh

Explanation: Till shimla, no stations are there.
So, it's first character will be the unique
smallest prefix. 

For safari, first character of shimla matches, so unique smallest prefix is sa

Similarly, j is smallest unique prefix for jammu and d is for delhi. 

For last city jammu, we have countered jammu before, and therefor no smallest
prefix is possible. So, we can generate its code as jammu with suffix equal to the count of jammu,
i.e, 2. Smallest unique prefix is jammu2. 

Now, delhi can be renamed as d, and dehradun can be renamed as deh, since deh is the smallest non
matching prefix. 
```

https://practice.geeksforgeeks.org/problems/renaming-cities28581833/1/