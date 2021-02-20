## Maximum Subarray Problem
<p>
Suppose that you been offered the opportunity to invest in the Volatile Chemical Corporation. Like the chemicals the company produces, the stock price of the Volatile Chemical Corporation is rather volatile. You are allowed to buy one unit of stock only one time and then sell it at a later date, buying and selling after the close of trading for the day. To compensate for this restriction, you are allowed to learn what the price of the stock will be in the future. Your goal is to maximize your profit. Figure 4.1 shows the price of the stock over a 17-day period. You may buy the stock at any one time, starting after day 0, when the price is $100 per share. Of course, you would want to “buy low, sell high”, buy at the lowest possible price and later on sell at the highest possible price, to maximize your profit. Unfortunately, you might not be able to buy at the lowest price and then sell at the highest price within a given period. In Figure 4.1, the lowest price occurs after day 7, which occurs after the highest price, after day 1.</p>

![Example1](example1.png?raw=true "Example1")

<p>
You might think that you can always maximize profit by either buying at the lowest price or selling at the highest price. For example, in Figure 4.1, we would maximize profit by buying at the lowest price, after day 7. If this strategy always worked, then it would be easy to determine how to maximize profit: find the highest and lowest prices, and then work left from the highest price to find the lowest prior price, work right from the lowest price to find the highest later price, and take the pair with the greater difference. Figure 4.2 shows a simple counterexample, demonstrating that the maximum profit sometimes comes neither by buying at the lowest price nor by selling at the highest price.</p>

![Example2](example2.png?raw=true "Example2")

Write an algorithm that behaves as follow:
- **Input**: array of size *n* such that position *i* holds the price of one unit of stock at the end of day *i*
- **Output**: indexes *buy*, *sell* and the value *profit*, meaning that we optimize our income by buying one unit of stock at the end of day *buy-1*, selling at the end of day *sell*. *profit* rapresents our net margin.

### How to compile and run
```
go build
./maximum-subarray-problem 100 113 110 85 105 102 86 63 81 101 94 106 101 79 94 90 97
```
To execute using rondom data:
```
for i in {1..5000}; do PRICES="$((1 + $RANDOM)) $PRICES"; done
./maximum-subarray-problem "$(echo "$PRICES")"
PRICES=
```
