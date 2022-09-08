# case 1

(id: 4) THINKING
(id: 4) HUNGRY
(id: 2) THINKING
4 give me my left!
(id: 1) THINKING
(id: 0) THINKING
0 giving my right
4 give me my right!
(id: 3) THINKING
3 giving my left
4 getting my right
(id: 3) HUNGRY
3 give me my left!
3 give me my right!
2 giving my left
3 getting my left
(id: 2) HUNGRY
2 give me my left!
3 getting my right
(id: 3) EATING: 1
4 getting my left
(id: 4) EATING: 1

## Explain

Might be due to the defer method in lines 113 and 115 in philo


# case 2
(id: 4) THINKING
(id: 4) HUNGRY
4 give me my left!
(id: 1) THINKING
(id: 2) THINKING
(id: 0) THINKING
0 giving my right
4 give me my right!
(id: 3) THINKING
3 giving my left
4 getting my right
(id: 3) HUNGRY
3 give me my left!
3 give me my right!
2 giving my left
3 getting my left
3 getting my right
(id: 3) EATING: 1
(id: 2) HUNGRY
2 give me my left!
4 getting my left
(id: 4) EATING: 1

# case 3
(id: 4) THINKING
(id: 4) HUNGRY
4 give me my left!
(id: 3) THINKING
(id: 1) THINKING
(id: 2) THINKING
(id: 0) THINKING
0 giving my right
4 give me my right!
3 giving my left
4 getting my right
(id: 3) HUNGRY
3 give me my left!
3 give me my right!
2 giving my left
3 getting my right
3 getting my left
(id: 3) EATING: 1
4 getting my left
(id: 4) EATING: 1
(id: 2) HUNGRY
2 give me my left!

# case 4
(id: 4) THINKING
(id: 4) HUNGRY
4 give me my left!
(id: 1) THINKING
(id: 2) THINKING
(id: 3) THINKING
(id: 0) THINKING
0 giving my right
4 give me my right!
3 giving my left
4 getting my right
(id: 3) HUNGRY
3 give me my left!
3 give me my right!
2 giving my left
-- 3 getting my left
3 getting my right
(id: 3) EATING: 1
(id: 2) HUNGRY
2 give me my left!
4 getting my left
(id: 4) EATING: 1
-- 4 giving my right to defered

# case 5
(id: 4) THINKING
(id: 4) HUNGRY
4 give me my left!
(id: 0) THINKING
0 giving my right
(id: 3) THINKING
(id: 2) THINKING
4 give me my right!
3 giving my left
(id: 1) THINKING
(id: 3) HUNGRY
3 give me my left!
4 getting my right
4 getting my left
(id: 4) EATING: 1
(id: 4) THINKING
4 giving my right
3 give me my right!
2 giving my left
3 getting my right
(id: 2) HUNGRY
2 give me my left!
2 give me my right!
1 giving my left
(id: 1) HUNGRY
1 give me my left!
2 getting my right
1 give me my right!
0 giving my left
1 getting my right
1 getting my left
(id: 1) EATING: 1
2 getting my left
(id: 2) EATING: 1
3 getting my left
(id: 3) EATING: 1
3 giving my right to defered
(id: 0) HUNGRY
0 give me my left!
