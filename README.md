# thea_interpreter

## Why

I have an idea for something a little more ambitious, but I need to understand these things first before I jump into it. So, I'm making an interpreter (and maybe a compiler in the future) for a language I'm creating called Thea. Eventually, I will translate this knowledge into a new language called Cynthia.

I don't really have a big plan for Thea at the moment, I think it's just going to be another quirky language, right now the gimmick is that the keywords are only two characters because I hate memorizing long winded and verbose bits of code.

Though, I imagine reading something like `if (tr) rt fl; el rt tr;` (if true, return false, else, return true). It looks kind of cute to me (albeit, maybe a little hard to read).

Also, I think most keywords are just kind of useless? Like if I'm writing a recursive fibonacci function, I don't need to see the entire return keyword, my eyes just glaze over it, so it helps in filtering out noise in programming code.

<pre>
<b>lt</b> fib = <b>fn</b>(n) {
	<b>if</b> (n < 2) <b>rt</b> n;
	<b>el</b> <b>rt</b> fib(n - 1) + fib(n - 2);	
}
</pre>


This looks clean to me. Tight.

That's what it is right now, maybe it'll change in the future if I think of a revolutionary new idea (hah).

## Some of the things you can do with Thea

Arithmetic Expressions

<pre>
5 + 5;
</pre>

<pre>
10
</pre>

Variable Bindings

<pre>
<b>lt</b> a = 5;
<b>lt</b> b = 10;
b * a;
</pre>
<pre>
50
</pre>

Functions

<pre>
<b>lt</b> a = <b>fn</b>(x, y, z) { x + y + z; }
a(1, 2, 3);
</pre>

<pre>
6
</pre>

Higher-Order Functions

<pre>
<b>lt</b> a = <b>fn</b>(x) { <b>fn</b> (y) { x + y } };
<b>lt</b> b = a(2);
b(3);
</pre>

<pre>
5
</pre>

