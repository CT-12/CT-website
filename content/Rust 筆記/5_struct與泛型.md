---
create_at: 2024.01.22
update_at: 2025.07.09
draft: false
tags: 
 - Rust
 - Struct
 - 泛型
---

# Struct 與 泛型

泛型經常讓我搞不太明白，有可能是幾乎沒用過，很常看到後腦袋就會當機，因此打算寫一篇筆記幫助自己之後回憶。

首先在 struct 使用泛型會長得像下面這樣：

```rust
/// 只有一個泛型的情況
/// 這邊指的是 T 可以是任何型別，並且 x 和 y 會被要求是一樣的型別
struct Point<T> {
    x: T,
    y: T,
}

/// 多個泛型的情況
/// T 和 U 可以是任何的型別，但是現在 x 和 y 可以是一樣的型別，也可以是不一樣的型別
struct Point<T, U> {
    x: T,
    y: U,
}
```

編譯器會根據在實例化 struct 時實際填入的值來推斷 T 和 U 是什麼型別。到這裡都挺好瞭解的，比較不好瞭解的地方是方法。首先直接給一個極端的案例來說明這部分，下面程式碼為 struct implement 了 兩個方法：

```rust
#[derive(Debug)]
struct Point<T, U> {
    x: T,
    y: U,
}

impl<A: fmt::Debug> Point<A, A> {
    fn print1(&self) {
        println!("x: {:?}, y: {:?}", self.x, self.y);
    }
}

impl<A: fmt::Debug, B: fmt::Debug> Point<A, B> {
    fn print2(&self) {
        println!("x: {:?}, y: {:?}", self.x, self.y);
    }
}
```

可以很明顯地看到我故意在方法的泛型的地方放了 A 和 B，主要是想先說明方法的泛型符號不一定要跟 struct 的泛型符號一樣，編譯器是根據泛型參數的位置來對應的，像是 T = A、U = B，但還是推薦一樣會好看一點。

再來說明這兩個方法的不同，在 print1 的地方，泛型的部分放了兩個 A，這裡的意思是**這個方法只限定 x 和 y 型別相同的 struct 實例能夠使用**。以此類推，print2 這裡就代表**x 和 y 的型別不相同的 struct 實例可以使用這個方法**，當然，如果 x 和 y 型別相同的話也可以使用 print2。

我們也可以直接明確的定義限定的型別：

```rust
struct Point<T> {
    x: T,
    y: T,
}

// 只有型別是 f64 的 Point 實例可以使用下面這個方法
impl Point<f64> {
    fn distance_from_origin(&self) -> f64 {
        (self.x.powi(2) + self.y.powi(2)).sqrt()
    }
}

let p1 = Point {x: 1.0, y: 1.0} // 型別是 f64，這個 struct 可以使用！
let p2 = Point {x: 1, y: 1} // 型別是 i32，這個 struct 不可以使用...
```

最後我推薦看 [10.1 泛型資料型別
](https://rust-lang.tw/book-tw/ch10-01-syntax.html) 裡最下面範例 10-11 以及單型化的部分，可以更好地瞭解泛型的例子。

**補充 1**

前面說到 rust 會在實例化 struct 的時候填入對應的類別。這裡想展開說說實例方法是什麼時候被填入對應的類別的。先給範例：

```rust
struct Point<X1, Y1>{
    x: X1,
    y: Y1,
}

impl<X1, Y1> Point<X1, Y1> {
    fn mixup<X2, Y2>(self, other: Point<X2, Y2>) -> Point<X1, Y2> {
        Point {
            x: self.x,
            y: other.y,
        }
    }
}

fn main() {
    let p = Point { x: 5, y: 10.0 };       // ---- L1
    let p2 = Point { x: "Hello", y: 'c' }; // ---- L2
    let p3 = p.mixup(p2);                  // ---- L3
    println!("p3.x: {}, p3.y: {}", p3.x, p3.y);
}
```

程式執行到 `L1` 時，rust 會產生對應的類別的 struct 和方法，也就是：

```rust
struct Point<i32, f64>{
    x: i32,
    y: f64,
}

impl<i32, f64> Point<i32, f64> {
    // 現在 X2 和 Y2 還不知道會是什麼類型
    fn mixup<X2, Y2>(self, other: Point<X2, Y2>) -> Point<i32, Y2> {
        Point {
            x: self.x,
            y: other.y,
        }
    }
}
```

程式執行到 `L2` 時，rust 會產生對應的類別的 struct 和方法，也就是：

```rust
struct Point<&'static str, char>{
    x: &'static str,
    y: char,
}

impl<&'static str, char> Point<&'static str, char> {
    // 現在 X2 和 Y2 還不知道會是什麼類型
    fn mixup<X2, Y2>(self, other: Point<X2, Y2>) -> Point<&'static str, Y2> {
        Point {
            x: self.x,
            y: other.y,
        }
    }
}
```

程式執行到 `L3` 時，rust 發現 `p` 使用了 `mixup` 方法，並且傳入了 `p2`，現在他終於知道 `p（Point<i32, f64>`的 `mixup` 方法裡面的 X2 Y2 是什麼類型的，也就是：

```rust
struct Point<i32, f64>{
    x: i32,
    y: f64,
}

impl<i32, f64> Point<i32, f64> {
    fn mixup<&'static str, char>(self, other: Point<&'static str, char>) -> Point<i32, char> {
        Point {
            x: self.x, // 這裡的 self 是 Point<i32, f64>
            y: other.y,
        }
    }
}
```

泛型方法的類型推導大概就是這樣，值得注意的是，如果之後又有一行程式是 `let p4 = p.mixup(p3);`，那 rust 會在產生第二個 `mixup` 方法：

```rust
impl<i32, f64> Point<i32, f64> {
    fn mixup<i32, char>(self, other: Point<i32, char>) -> Point<i32, char> {
        Point {
            x: self.x,
            y: other.y,
        }
    }
}
```

這部分是我推測的，但感覺應該會是這樣。有錯之後會再回來改～

# References

[10.1 泛型資料型別
](https://rust-lang.tw/book-tw/ch10-01-syntax.html)

---
此筆記僅是為了紀錄學習過程，如欲完整學習 rust，請去看官方提供的文件，那裡有更加完善的學習資源。筆記內容為自己的理解因此可能存在錯誤，如有錯誤敬請見諒。