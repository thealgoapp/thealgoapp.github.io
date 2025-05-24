---
layout: layouts/post.njk
title: Big O!
date: 2025-06-05
author: dakolpakov
preview_image: /images/big-o/preview.webp
permalink: /blog/big-o/
draft: true
og_type: article
description: "A practical guide to Big-O notation and algorithm complexity for software engineers. Includes real examples, common pitfalls, and interview tips."

---

# Big-O!

Today, I want to dive into one of the most important theoretical topics: **how to estimate and compare different algorithms that solve the same problem**.

I’ll try to be especially careful not to overwhelm you, because this topic can easily lean into deep academic theory with a lot of math (and to be fair, I’m not particularly good at it myself). But let me be honest with you: as I mentioned in [How to Prepare for Algorithm Interviews](/blog/prepare-for-Interviews-p1/), **algorithms and data structures are both a science and an art — but our goal is to pass a developer interview**.

To be very clear: for that purpose, in 99% of cases, you **don’t need a deep academic background in CS or math**.
You just need to remember **a few basic formulas**, understand **what they mean**, and know **how to figure out which one describes your algorithm**.

---

## How to compare

First of all, let’s think about how we can compare different algorithms and what requirements they should meet. You can imagine dozens of possible criteria:
Code size: How much code is required to implement the algorithm? 

- **Readability**: How easy is it to understand the algorithm just by reading the code? 
- **Maintainability** and **flexibility**: How easy is it to modify or extend the algorithm in the future? 
- **Fault tolerance**: How well does the algorithm handle invalid input, edge cases, or runtime errors? 
- **Scalability**: How well does the algorithm perform under heavy usage or repeated calls?

All of these are important, especially in real-world tasks.
But when we talk about **algorithms and data structures**, we usually focus on **two key criteria**:
👉 **Time complexity** and **space (memory) complexity**.

---

## Asymptotic complexity

To estimate **time and space complexity** in the context of algorithms and data structures, we typically use the concept of **asymptotic complexity**.
Let’s be honest — it’s **a very rough estimate**, but that doesn’t mean we can’t use it to effectively **compare different algorithms**.

We can’t predict the exact (or even approximate) execution time of an algorithm — there are simply too many influencing factors.
These range from software to hardware, from ambient temperature to (jokingly) the phase of the moon.

You might run your code on a personal laptop with a general-purpose OS, full of background processes and managed by a standard task scheduler.
Or you might run the same program on a high-performance machine with a specialized OS that limits background activity, gives your code maximum priority, and operates in a temperature-controlled environment to avoid CPU throttling.

The outcomes could be **completely different**. That’s why, **at the design stage**, we avoid relying on specific timing assumptions.

Of course, we can use **benchmarking tools** later to get more precise performance data — but that requires **deep domain knowledge**.
You need to understand your expected input workload, the hardware used in production, the programming language, and the entire tech stack.

And even with all of that, there’s still **no guarantee** your time estimates will be accurate.

Another useful approach is **to count the number of actions** the algorithm needs to perform.
This is actually the right mindset — and it’s where most computer science courses begin their lectures on **asymptotic complexity**.

But in my experience, teachers often forget to emphasize one key point early on:
**👉 It’s just an approximation**.

And that often leads to two problems:
- Students start paying **too much attention** to counting every single step. 
- Later, they get confused when we suddenly **start ignoring exact counts**.


Many variables can affect the actual result.
In 99.99% of real-world cases, you’ll be writing code in a **high-level programming language**, which gets compiled down to machine code.
And depending on the language, compiler, and target hardware, the number of actual CPU instructions may vary significantly.

So whether your function contains **one line of code**, **three**, or even **a hundred** — it doesn’t really matter that much.

What **does** matter is:
**How does the number of operations grow as the input size increases?**
That’s what asymptotic complexity is all about — it tells us **how fast the execution time grows relative to the input size**.

---

## What is “Big-O”?

There are several different ways to describe **asymptotic complexity**.
The most common notations are:

- **Big O** — `O(f(n))`: describes the **upper bound** of growth.
- **Theta** — `Θ(f(n))`: represents the **exact bound** (tight bound).
- **Big Omega** — `Ω(f(n))`: describes the **lower bound** of growth.
- **Small o** — `o(f(n))`: **a strict upper bound** — the algorithm grows strictly slower than f(n).
- **Small omega** — `ω(f(n))`: a strict lower bound — the algorithm grows strictly faster than f(n).
- ...and more.

But to be honest — unless you're studying computer science **academically**, you can safely skip everything except **Big O notation**.

Before we dive deeper into Big O, let’s define a few **key principles** to keep in mind.

### Key principles

#### ✅ We focus on large inputs
This makes sense — if you're sorting just five numbers, it doesn’t matter which algorithm you use.

Even randomly shuffling them a few times might eventually work. But once the input size starts to grow, algorithm efficiency becomes critical.

#### ✅ We take a pessimistic view — worst-case analysis

We usually estimate how the algorithm behaves in **the worst-case scenario**, so we know it will perform at least that well in any other situation.

Sometimes we make exceptions — especially when the worst-case is **extremely unlikely**.

For example, looking up a value by key in a hash map has **constant time complexity O(1)** on average.
Technically, there is a **worst-case scenario** where the complexity could **degrade to O(n)** (if all keys hash to the same bucket, and we have to search through them one by one).

But in **real-world applications**, this rarely happens, especially if you're using **standard built-in hash maps** with a **well-designed hash function that distributes keys evenly**.
In that case, the **average-case complexity** is far more relevant than the worst-case.

---

### ⚖️ A word about balance

**Big O notation** is designed to describe the **upper bound of an algorithm’s growth**.
But that doesn’t mean it’s enough to just pick the fastest-growing function you can think of — we need to strike a **balance between a theoretical approximation and real-world usefulness**.

For example, imagine an extremely fast-growing function like `f(n) = ((n!)^n!)!`.
This function explodes to infinity even for small inputs.
If `n = 3`, we already get `f(n) = 46656!` — a number so huge it’s almost unimaginable.

Technically, this would still qualify as a **valid Big O upper bound** for any algorithm.
But in practice, that kind of estimate is **useless** — we wouldn’t be able to compare different algorithms meaningfully.

So, how do we strike that balance?
We focus on the **fastest-growing part** of the code — that’s what determines the overall complexity.

Let’s look at an example:

```go
package main

import "fmt"

func main() {
	input := []int{1, 7, 2, 4, 6, 3, 5, 8, 9}
	increased := make([]int, 0, len(input))
	summary := 0

	for i := 0; i < len(input); i++ {
		for j := i+1; j < len(input); j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}

	for i := 0; i < len(input); i++ {
		increased = append(increased, input[i]+10)
	}

	summary = 123 + 456

	fmt.Println(input)
	fmt.Println(increased)
	fmt.Println(summary)
}
```

Now let’s estimate the complexity.

First of all, we can split this code **into three parts**.

---

#### First part – simple in-place sorting.
We compare all elements **with each other**, and the number of operations depends on the input length.
As I said before, the **exact count** depends on many factors, including the programming language.
I wrote this code in Go, which allows you to swap elements in a single line:
```go
input[i], input[j] = input[j], input[i]
```

Other languages usually require **creating a temporary variable** for this operation.
**Eventually**, the compiler will **translate** this code into some **CPU instructions**, and we can’t know how many instructions will be generated in the end.
It depends on the **compiler** and **hardware**.

So it doesn’t matter how many lines are inside the loop.
The important thing is that these actions are **inside a loop**. And this loop is **inside the outer loop**. And both of them **depend on the length of the input**.

> 📝**Here’s a little trick** — some built-in functions aren’t atomic. For example, **sorting**.
> One of the most common interview mistakes is when the candidate **forgets to account** for such functions, assuming that “one line of code = one action.”

It means that we perform n actions in the inner loop and repeat the outer loop n times — in total, **n × n = n² operations**.

You could argue that the inner loop performs fewer actions because it starts from the next element after the outer loop index.
The exact complexity is the sum of the first k natural numbers, which means:
`f(n) = n(n - 1)/2 ≈ (n²)/2`

But in **asymptotic complexity**, we can **ignore constant** multipliers.
It’s the same idea as with step-counting — if we don’t care about the exact number of operations, we **can safely ignore any constant** multipliers as well.

#### Second part - incrementing each original element.
This part is simple: we perform n actions.

#### Third part - just summing two integers.
This doesn’t depend on the input size, so it has constant complexity.

#### Finally
We get something like this:
`f(n) = n² + n + 1`.

But let’s remind ourselves of a key principle: "**We focus on large inputs**".
If we’re talking about, say, an input size of **1,000,000**, then everything **except n²** becomes **negligible**.

Here’s the **key idea**:
We care only about **the fastest-growing part** of the algorithm.
Everything else — like constant multipliers, lower-degree terms of n, or constant-time operations — **doesn’t matter** at this stage.

> 📝Of course, they can matter in real-world applications where input sizes are small or performance margins are tight — but not when analyzing asymptotic growth.

---

## Typical time complexities

Does it mean we should know by heart every math/statistics formula? Of course, no. Here are some typical complexities that we usually face.

### O(1) - Constant
The **fastest algorithms** are those whose performance **doesn’t depend on input size**.
In practice, we usually can’t achieve this complexity for **the entire algorithm**, but we can often improve specific parts by using data structures that allow certain operations to run in constant time.

A typical example is a **hash map**, which provides **constant-time access** to elements by key.
So, for instance, when we need to find a pair of elements in an array, we can use a temporary hash map instead of a nested loop, reducing the complexity from **O(n²)** to **O(n)**.

### O(log n) - Logarithmic

This is what we call a **pretty fast class of algorithms**.
In these cases, **the base of the logarithm doesn’t matter**, so we simply write `log(n)` without specifying it.

We typically achieve this complexity when we **repeatedly split the input into smaller parts**.
Examples include **binary search**, **tree-based search**, **input splitting** in merge sort, and similar techniques.

Just like with constant-time complexity, we usually **can’t apply logarithmic time complexity to the entire algorithm**, because it often requires the input **to be preprocessed or structured in a specific way** — for example, sorted, as in binary search.

However, logarithmic complexity can still be **a part of a larger algorithm** and help reduce its **overall complexity**.

> **Spoiler alert**: One of the most common time complexities you’ll encounter in real-world algorithms is **O(nlog(n))**.


### O(n) - Linear
Still a very **fast class of algorithms**.

This complexity means you can solve the problem by **scanning the input once or a few times**, possibly with the help of simple temporary structures like a hash map to store intermediate data.

The key is that you **don’t** have **nested loops**, no **comparing each element with every other element**, and no **deep recursion** — just a **straightforward pass** over the data.

You’ll encounter this complexity both in **real-world applications** and in **interview problems**.
And in most interviews, if you manage to reduce your solution to **linear time** — you’ve likely found the **optimal answer**.


### O(n log n) - Linearithmic
Usually, this is the **last “fast” class of algorithms**, at least when it comes to **interview** problems (**not necessarily in the real world**).

As I mentioned earlier, we can often **reduce quadratic complexity to linearithmic** by using auxiliary structures (like **trees**) or by applying the **divide-and-conquer** principle.
Some data structures — particularly **trees and graphs** — naturally provide this complexity for useful operations like search.

The most common context where you’ll see this complexity is **sorting** — optimal sorting algorithms have `O(n*log(n))` time complexity.

And here’s a **strong recommendation**: **remember this fact**, and try to **dive deeper** into at least one such algorithm — it’s a **very popular interview topic**.

> A small life hack: **pay close attention to built-in sorting functions**.
> Every programming language provides a built-in sorting function, and it’s usually just a **one-liner**.
> A very **common mistake** in interviews is when candidates forget to **account** for this "one-line" function when trying to estimate the overall complexity of their code.

Just like with linear-time complexity, if you manage to reduce your algorithm to **linearithmic complexity**, that’s a good sign — you’ve probably found an **optimal solution**.

### O(n²) - Quadratic

This is usually the first **“slow” class of algorithms** in **interviews**.
In interviews, O(n²) is often the complexity of the **first working** solution — and a good moment to ask yourself:

**👉 Can we do better?**

In most cases, the answer will be **yes**.

> It can still be a reasonable estimation for more specific problems — for example, many matrix-related problems have slower complexity.
> But matrix operations aren’t a very common interview topic in general.
> They may show up in domain-specific roles, like graphics programming or scientific computing.

You typically reach this complexity when you use a **nested loop**, **compare each element with every other**, and so on.

If you see **one nested loop** — that’s **quadratic**. **Two nested loops** — that’s **cubic**. And so on.


### Other estimations.
Of course, you can imagine/face other estimations, like **O(2ⁿ) - Exponential** or **O(n!) - Factorial**.
But usually this means you are in the **wrong way**, because it’s typical estimations for different **brute-force** solutions.
