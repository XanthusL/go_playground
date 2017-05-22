extern crate rand;
// 使用标准库的io
use std::io;
use rand::Rng;
use std::cmp::Ordering;

fn main() { lifetimes(); }
// 0x00000000
fn guessNumber() {
    let rand_number = rand::thread_rng().gen_range(1, 101);
    loop {
        // " ! " 代表调用了一个宏而不是一个普通函数
        println!("Input a number:");
        // let foo = bar;默认不可变 immutable
        // mut 使绑定可变
        let mut guess = String::new();
        io::stdin().read_line(&mut guess)
            .expect("Failed to read line");
        println!("You guessed: {}", guess);
        println!("Random number is: {}", rand_number);

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        match guess.cmp(&rand_number) {
            Ordering::Less => println!("Too small"),
            Ordering::Greater => println!("Too big"),
            Ordering::Equal => {
                println!("Bingo");
                break;
            }
        }
    }
}
// 0x00000001
fn ownership() {
    let v = vec![1, 2, 3];
    let v2 = v;
    // Compile fails, 'use of moved value'
    // println!("v[0] is: {}",v[0]);

    let n = 1;
    let n2 = n;
    // Compile success, i32 implements Copy
    println!("n is: {}", n);
}
// 0x00000002
fn borrow() {
    let mut v1 = vec![1, 2, 3];
    func4borrow(&mut v1);
    /*
    第一，任何借用必须位于比拥有者更小的作用域。
    第二，对于同一个资源（resource）的借用，以下情况不能同时出现在同一个作用域下：
            1 个或多个不可变引用（&T）
            唯一 1 个可变引用（&mut T）
    译者注：即同一个作用域下，要么只有一个对资源 A 的可变引用（&mut T），
    要么有 N 个不可变引用（&T），但不能同时存在可变和不可变的引用
    */
}

fn func4borrow(v: &mut Vec<i32>) {
    v.push(4)
}
// 0x00000003
fn lifetimes() {
    //    let r;              // Introduce reference: `r`.
    //    {
    //        let i = 1;      // Introduce scoped value: `i`.
    //        r = &i;         // Store reference of `i` in `r`.
    //    }                   // `i` goes out of scope and is dropped.
    //    // 悬垂指针（dangling pointer）
    //    println!("{}", r);  // `r` still refers to `i`.
    let line = "lang:en=hello";
    let lang = "en";
    let v;
    {
        let p = format!("lang:{}", &lang);
        v = func4lifetimes(line, p.as_str());
    }
    println!("{}", v);
}

fn func4lifetimes<'a, 'b>(line: &'a str, prefix: &'b str) -> &'a str {
    //format!("{},{}", line, prefix)
    return "asdfasdf";
}