---
layout: layouts/post.njk
title: How to Succeed During the Interview (Part 2, Strategy & Communication)
date: 2025-06-03
author: dakolpakov
preview_image: /images/previews/prepare-for-Interviews-p2.webp
draft: true
permalink: /blog/prepare-for-Interviews-p2/
og_type: article
description: Learn how to approach algorithm interviews like a pro. Ask the right questions, communicate clearly, think out loud, and handle edge cases with confidence.

---

# How to Succeed During the Interview (Part 2: Strategy & Communication)

Before we start, let’s take a moment to talk about the interview process and what makes a good one.

You can read the first part of this article [here](/blog/prepare-for-Interviews-p1/).


---

## 🧭 Several Principles

Here’s a little spoiler:
> **A good interview is usually a two-way conversation.**

And the most important advice is this — don’t be afraid to **talk to your interviewer**.
Ask questions, share your assumptions, and let the interviewer guide you with follow-up questions or hints.

Most importantly, give them the space to make decisions whenever possible — whether it’s about the direction of the solution, level of optimization, or how deep to go.


Let’s break it down.

---

### ☑️ Align with the Interviewer

It’s always helpful to understand **what the interviewer expects** from this part of the process, especially in terms of time limits.
This can work in your favor.
For example, if you realize you're stuck halfway through a problem, you can ask:

> “Would you prefer I keep pushing on this one, or should we move to the next task?”

---

### ☑️ Read the Task Carefully

Even if you’ve seen a similar problem before, **don’t jump into coding just based on the title**.
Review the **requirements**, **examples**, and **constraints** — they often contain helpful hints that clarify the problem and narrow the scope.
Interviewers sometimes **tweak well-known problems** specifically to catch candidates who rely on memorized solutions without truly understanding them.

---

### ☑️ Ask Clarifying Questions

**Asking** clarifying questions is a **green flag** — it shows that you’re thoughtful and care about aligning with the stakeholders’ needs.

> Trust me, It's always better to solve the actual problem than something you assumed.

---

### ☑️ Don’t Rush Into Coding

Even if you're 100% confident in your skills and solution, take a moment to talk through your approach out loud with the interviewer first.
They might point out something important, help clarify the requirements, highlight edge cases, or even decide to skip the coding part to save time.

---

### ☑️ Challenge your solution and think about edge cases before coding

Sometimes, a solution might work for common inputs but completely fail on edge cases that still fall within the task requirements.
In that case, the solution becomes useless.
**Thinking through these scenarios** before writing any code is a **big green flag** for the interviewer — it shows that you think critically before jumping into implementation, which is a **valuable quality for any engineer**.

Test your plan with weird inputs:
- Empty arrays
- All identical elements
- Already sorted input
- etc

---

### ☑️ Start with a Simple Solution

Don’t aim for the most optimized version right away.
It's often better to begin with a **basic brute-force** approach that **works**, rather than **getting stuck halfway** through a more advanced one.
If you realize your solution can be improved, feel free to say something like:

> “I see that this solution has quadratic complexity complexity. Should I finish this version first, or would you prefer we jump to some optimized ideas?”

That way, you're showing awareness of **performance trade-offs** and leaving it to the interviewer to decide what direction to take.

---

### ☑️ Describe your thinking while coding

Speak your thoughts:
> “I’m using a hash map here to track counts...”

The main goal here is to **show** your **mindset** and **problem-solving process**.
Speak as you go — **explain** your assumptions, what you’re trying to achieve, and why you’re choosing a particular approach.
That doesn’t mean you need to talk non-stop: it’s totally fine to pause and think. But **try to avoid long silences**.
By thinking out loud, you give the interviewer a chance to **follow your logic**, **offer guidance**, or even **give subtle hints**.

---

### ☑️ Estimate Complexity

A key part of algorithm design is **comparing** different approaches to solving the same problem. 
In many cases, there are **countless possible solutions** — or sometimes none at all.
Of course, the most **basic requirement** is that your algorithm must return the **correct output for valid inputs**.
If it can’t do that, it’s not really an algorithm.

But once you have a few working solutions, you should be able to **compare** them.
In 99% of cases, this means evaluating their `asymptotic complexity` using `Big O notation`.
(We’ll explain those terms in a future post.) 
For now, just keep in mind: it’s not enough to find a solution — the interviewer expects you to **evaluate its performance and understand whether it’s optimal or not**.


---

### 🧠 “Can We Do Better?” — Tim Roughgarden, Stanford CS

No one expects you to come up with a unique or groundbreaking algorithm during a tech interview.
Most interviewers are looking for a well-known solution pattern from a familiar list.
But once you’ve found a working solution, **don’t stop there**.
Tim Roughgarden, in his Stanford course, says:
> As an algorithm designer, you should adopt as your Mantra the question, “**can we do better?**”

That’s a great question to ask once your code works — and it’s exactly the kind of thinking your interviewer wants to see.

---

## 🚩 Worst-Case Interview: What Not to Do

The worst-case scenario I can imagine as an interviewer looks something like this:

The candidate **starts coding immediately** after reading the task description.
They **don’t ask** any clarifying questions and **make no comments** during implementation.
At some point, they get stuck and silently try to “debug” **without explaining their thought process**.
Or maybe they write a solution that only works for a few examples but fails on edge cases.

In the best case, they might **even end up with a fully working and optimized solution** — maybe **the best one possible** — but this would still be **a bad interview**.

Why? Because the process matters more than just the final result. It’s better to not finish the solution at all, but **actively communicate**, **think out loud**, and **demonstrate how you approach the problem**.

---

## 💡 Final Thoughts

Remember:
> A technical interview isn’t just about hard skills.  
> The interviewer is also evaluating your **communication, adaptability, and mindset**.


You don’t need to be perfect.  
Just be thoughtful, ask questions, and work through problems like a teammate — not a solo coder.

