# Sorting algorithms

Algorithms are based on [algorithms with go](https://algorithmswithgo.com/).

Sorting algorithms implemented: 

- Bubble Sort: Time complexity of O(N²);
<p style="text-align:center;"><img src="https://upload.wikimedia.org/wikipedia/commons/3/37/Bubble_sort_animation.gif"></p>    
<p style="text-align:center;">Source: <a href="https://pt.wikipedia.org/wiki/Bubble_sort">Wikipedia</a></p>

- Insertion Sort: Time complexity of O(N²);
<p style="text-align:center;"><img src="https://upload.wikimedia.org/wikipedia/commons/2/25/Insertion_sort_animation.gif"></p>    
<p style="text-align:center;">Source: <a href="https://pt.wikipedia.org/wiki/Insertion_sort">Wikipedia</a></p>

- Shell Sort: Time complexity of O(N (log N)²).
<p style="text-align:center;"><img src="https://upload.wikimedia.org/wikipedia/commons/d/d8/Sorting_shellsort_anim.gif"></p>    
<p style="text-align:center;">Source: <a href="https://en.wikipedia.org/wiki/Shellsort">Wikipedia</a></p>

- Merge Sort: Time complexity of O(N log N)
<p style="text-align:center;"><img src="https://upload.wikimedia.org/wikipedia/commons/c/c5/Merge_sort_animation2.gif"></p>    
<p style="text-align:center;">Source: <a href="https://pt.wikipedia.org/wiki/Merge_sort">Wikipedia</a></p>

- Quick Sort: Time comlexity of O(N²)
<p style="text-align:center;"><img src="https://upload.wikimedia.org/wikipedia/commons/6/6a/Sorting_quicksort_anim.gif"></p>    
<p style="text-align:center;">Source: <a href="https://pt.wikipedia.org/wiki/Quicksort">Wikipedia</a></p>

## Benchmarks

- 1) Bubble Sort:

```
pkg: github.com/nahumsa/sorting-algorithms/bubble
BenchmarkBubbleSortInt/100-4         	   45069	     26526 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSortInt/200-4         	   14342	     79352 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSortInt/400-4         	    4713	    254083 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSortInt/800-4         	    1260	    906926 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSortInt/1600-4        	     350	   3576257 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSortInt/3200-4        	      82	  14179024 ns/op	       0 B/op	       0 allocs/op
```

- 2) Insertion Sort:

```
pkg: github.com/nahumsa/sorting-algorithms/insertion
BenchmarkInsertionSortInt/100-4         	   32619	     36145 ns/op	   23378 B/op	     102 allocs/op
BenchmarkInsertionSortInt/200-4         	   12968	     91908 ns/op	   89208 B/op	     203 allocs/op
BenchmarkInsertionSortInt/400-4         	    3894	    347731 ns/op	  349298 B/op	     403 allocs/op
BenchmarkInsertionSortInt/800-4         	    1238	   1238189 ns/op	 1386824 B/op	     803 allocs/op
BenchmarkInsertionSortInt/1600-4        	     277	   3954662 ns/op	 5489901 B/op	    1604 allocs/op
BenchmarkInsertionSortInt/3200-4        	      73	  18373725 ns/op	21718151 B/op	    3205 allocs/op
```

- 3) Merge Sort:

```
pkg: github.com/nahumsa/sorting-algorithms/merge
BenchmarkInsertionSortInt/100-4         	   59292	     21293 ns/op	    5824 B/op	      99 allocs/op
BenchmarkInsertionSortInt/200-4         	   31364	     42571 ns/op	   13440 B/op	     199 allocs/op
BenchmarkInsertionSortInt/400-4         	   14608	     87831 ns/op	   30080 B/op	     399 allocs/op
BenchmarkInsertionSortInt/800-4         	    8118	    214011 ns/op	   66688 B/op	     799 allocs/op
BenchmarkInsertionSortInt/1600-4        	    3315	    352300 ns/op	  146944 B/op	    1599 allocs/op
BenchmarkInsertionSortInt/3200-4        	    1534	   1128712 ns/op	  321152 B/op	    3199 allocs/op
```

- 4) Quick Sort:

```
pkg: github.com/nahumsa/sorting-algorithms/quick
BenchmarkInsertionSortInt/100-4         	  202778	      6099 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSortInt/200-4         	   95415	     12674 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSortInt/400-4         	   45226	     26556 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSortInt/800-4         	   19652	     62347 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSortInt/1600-4        	    9517	    124597 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertionSortInt/3200-4        	    4659	    263247 ns/op	       0 B/op	       0 allocs/op
```

- 5) Shell Sort:

``` 
pkg: github.com/nahumsa/sorting-algorithms/shell
BenchmarkInsertionSortInt/100-4         	   21268	     55858 ns/op	    4144 B/op	     225 allocs/op
BenchmarkInsertionSortInt/200-4         	    8606	    135254 ns/op	    8608 B/op	     467 allocs/op
BenchmarkInsertionSortInt/400-4         	    2769	    392936 ns/op	   17888 B/op	     965 allocs/op
BenchmarkInsertionSortInt/800-4         	     978	   1182319 ns/op	   38320 B/op	    1991 allocs/op
BenchmarkInsertionSortInt/1600-4        	     294	   4058921 ns/op	   79536 B/op	    4081 allocs/op
BenchmarkInsertionSortInt/3200-4        	      69	  15145407 ns/op	  165216 B/op	    8338 allocs/op
```
