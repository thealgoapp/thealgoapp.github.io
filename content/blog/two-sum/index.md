---
layout: layouts/post.njk
title: Two Sum
date: 2025-06-02
original_url: https://leetcode.com/problems/two-sum/
author: dakolpakov
preview_image: /images/previews/two-sum/preview.webp
complexity: easy
draft: true
permalink: /blog/two-sum/
og_type: article
description: 
tags:
  - easy
  - array
  - hash table
---

# Two Sum!

Let’s move on to the first problem — “**Two Sum**”, one of the most well-known LeetCode questions.
Despite its popularity, I’ve actually encountered this task in real interviews.
It’s often used as a warm-up, but many engineers only manage to reach the basic brute-force solution.

> 🔁 Don’t forget the principles we discussed in the previous [post](/blog/prepare-for-Interviews-p2/)!

## 📄Description

Given an array of integers nums and an integer target, return the indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.


### Constraints

- `2 <= nums.length <= 104`
- `-109 <= nums[i] <= 109`
- `-109 <= target <= 109`
- `Only one valid answer exists.`


### Examples

{% tabs %}

{% tab "Example 1" %}

**Inputs**

```
nums = [2,7,11,15]

target = 9
```

**Output**: `[0,1]` or `[1, 0]`

**Explanation**: Because `nums[0]` + `nums[1]` == `9`, we return `[0, 1]`.

{% endtab %}

{% tab "Example 2" %}

**Inputs**

```
nums = [3,2,4]

target = 6
```

**Output**: `[1, 2]` or `[2, 1]`

{% endtab %}

{% tab "Example 3" %}

**Inputs**

```
nums = [3,3]

target = 6
```

**Output**: `[0, 1]` or `[1, 0]`

{% endtab %}

{% endtabs %}

---

## 🐢 Brute Force

As we discussed in the previous post, let’s not try to jump straight to the most optimized solution.
It’s better to start with a working, well-known implementation and improve it step by step.

Also, pay close attention to the **problem definition** and **constraints**.
One important detail is that there’s a guarantee: **every input has exactly one valid answer**.

This means we don’t need to handle edge cases or check for multiple solutions — we can safely return the first valid result we find.

In this case, the brute-force solution is straightforward — we need to find a pair of values in a single array that add up to the target.
So, we can use two loops: an outer loop (let’s call it `i`) to pick the first element, and a nested loop (`j`) to look for a matching pair.

A small trick: we can start the nested loop from the next index (`i + 1`).
This way, we avoid unnecessary comparisons and ensure we don’t return the same element twice.

If the sum of `nums[i]` and `nums[j]` equals the `target`, we’ve found our answer — the pair `[i, j]`.

### 💻Implementation

{% renderFile "_includes/components/solution.njk", taskName = "two-sum", fileName="bruteforce" %}

### 📊 Complexity

When we talk about complexity, we usually try to estimate two things:

- **Time complexity** — The **approximate** number of operations the algorithm performs — this is called **asymptotic complexity**. We’ll cover it in more detail in a future post, but for now, remember that we typically describe it using **Big O notation**. This helps us estimate how the algorithm’s execution time grows as the input size increases.
- **Space complexity** — how much additional space our algorithm requires to work.

> A little spoiler: in most cases, there's a trade-off — we can use more memory to improve speed. 
> And that often makes sense in the real world, since memory is generally cheaper than CPU time.

In our brute-force solution, we compare all possible pairs of values using an outer and a nested loop.

In the **worst-case scenario**, the first element is compared with the second, third, … all the way to the n‑th element.
Then, the second element is compared with the third, fourth, … up to the end, and so on — where n is the length of the input array.

You might notice this forms an arithmetic progression: `(n - 1) + (n - 2) + ... + 1`

Using the formula for the sum of the first k natural numbers, we get: `S = n(n - 1)/2 ≈ (n²)/2`

As mentioned earlier, we aim for an **approximate estimation** — often called **the order of growth**.

To simplify, we focus only on **the dominant term** (the one that grows the fastest as n increases) and ignore constants and lower-order terms.

That gives us a final time complexity of: **O(n²)** or **quadratic complexity**.

> 🔍 Rule of thumb:  
> One nested loop → at least O(n²)  
> Two nested loops → at least O(n³)  
> And so on.

And since we don’t use any additional data structures, our brute-force algorithm has **constant space complexity** — **O(1)**.

---

## ⚡ Optimized Solution (with Hash Map)

Do you remember the Algorithm Designer’s Mantra from **Tim Roughgarden**?

> Can we do better?

Let’s try!

We know that:

```
nums[i] + nums[j] == target  →  nums[j] = target - nums[i]
```

Now, imagine we could quickly check whether a specific number exists in the array. Then we wouldn’t need a nested loop at all. Instead, for each i, we’d compute:
```
complement = target - nums[i]
```
...and simply check if `complement` exists in the array.

But there’s a catch: arrays don’t give us fast lookups.
If we want to check whether a value exists in an array, we have to scan it element by element — that’s still **O(n)** per lookup, which brings us back to **O(n²)** in total.

Fortunately, we have a better tool for this: a **hash map**.
This data structure allows us to **check for the existence of a value in constant time** — **O(1)**.

> To be fair, in rare worst-case scenarios — like when many keys collide due to a poor hash function — the performance can degrade to O(n).
> But if we use a built-in hash map from a modern language without tampering, it will generally provide a uniform distribution and maintain O(1) complexity.

So the idea is simple:

While iterating through the array, we store each number and its index in a hash map — the **number as the key** and its **index as the value**.
Then, we iterate through the array again and, for each element, check whether its **complement** exists in the map.
If it does — we’ve found the pair.

Now we still have two loops, but they’re **not nested** — so instead of **O(n²)**, we get **O(2n)**, which simplifies to **O(n)**.
In other words, our algorithm now has **linear time complexity**.

We can do even better by storing values in the hash map and checking for complements **in a single loop**.

For each element, we first calculate its **complement** and check whether it already exists in the hash map. 
- If it **does**, we’ve found the pair. 
- If it **doesn’t**, we store the current element and its index in the map.

This way, when we encounter the **first** element of the result pair, its complement has yet to be seen, so we store it.

When we reach the **second** element, we find that its complement (the first one) is already in the map, and we return the pair!

### 💻Implementation

{% renderFile "_includes/components/solution.njk", taskName = "two-sum", fileName="solution" %}

### 📊 Complexity

As we mentioned earlier, the **time complexity** is now **O(n)** — **linear time**.

What about **space complexity**?

In this solution, we use an additional data structure — **a hash map**.
In the worst-case scenario, the map could store up to **n - 1 elements** (if the matching pair is found only at the very end).

Just like with time complexity, we focus on **the order of growth**, so we simplify this to **O(n)** space complexity.

This is a classic **trade-off**: we use more memory, but significantly improve performance.
In this case, it’s a **clear win**.

## ✅ Takeaways

- **Start simple** — even brute-force has value
- Think about **complexity**: time and space
- Use the right data structure: **hash maps** are great for **quick checks**
- Ask: **Can we do better?**
