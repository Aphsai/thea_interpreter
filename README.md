# thea_interpreter

## Why

I have an idea for something a little more ambitious, but I need to understand these things first before I jump into it. So, I'm making an interpreter (and maybe a compiler in the future) for a language I'm creating called Thea. Eventually, I will translate this knowledge into a new language called Cynthia.

I don't really have a big plan for Thea at the moment, I think it's just going to be another quirky language, right now the gimmick is that the keywords are only two characters because I hate memorizing long winded and verbose bits of code.

Though, I imagine reading something like `if (tr) rt fl; el rt tr;` (if true, return false, else, return true). It looks kind of cute to me (albeit, maybe a little hard to read).

Also, I think most keywords are just kind of useless? Like if I'm writing a recursive fibonacci function, I don't need to see the entire return keyword, my eyes just glaze over it, so it helps in filtering out noise in programming code.

```
lt fib = fn(n) {
	if (n < 2) rt n;
	el rt fib(n - 1) + fib(n - 2);	
}
```


This looks clean to me. Tight.

That's what it is right now, maybe it'll change in the future if I think of a revolutionary new idea (hah).

## Some of the things you can do with Thea

Arithmetic Expressions

```
5 + 5;
```

```
10
```

Variable Bindings

```
lt a = 5;
lt b = 10;
b * a;
```
```
50
```

Functions

```
lt a = fn(x, y, z) { x + y + z; }
a(1, 2, 3);
```

```
6
```

Higher-Order Functions

```
lt a = fn(x) { fn (y) { x + y } };
lt b = a(2);
b(3);
```

```
5
```

