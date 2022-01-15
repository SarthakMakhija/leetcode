### Problem

Given a dictionary of words dict where each word follows CamelCase notation, print all words in the dictionary that match with a given pattern consisting of uppercase characters only.
CamelCase is the practice of writing compound words or phrases such that each word or abbreviation begins with a capital letter. Common examples include: “PowerPoint” and “WikiPedia”, “GeeksForGeeks”, “CodeBlocks”, etc.

Note: The order should be such that the output strings should be sorted by the lexicographic order of their abbreviations. For strings with same abbreviations the strings should be sorted by the usual lexicographic order. So, if Output Strings are Hi and HelloWorld, Hi should come first as H lexiographically is smaller than HW.

```
Example 1:

Input:
n = 8
dict[] = {Hi,Hello,HelloWorld,HiTec,HiGeek,HiTechWorld,HiTechCity,HiTechLab}
pattern = HA
Output: No match found
```

```
Example 2:

Input:
n = 3
dict[]={WelcomeGeek,WelcomeToGeeksForGeeks,GeeksForGeeks}
pattern = WTG
Output: WelcomeToGeeksForGeeks

Explanation: WTG occurs in WelcomeToGeeksForGeeks.
```

`Solution does not return in lexicographic order`

https://practice.geeksforgeeks.org/problems/camel-case04234120/1