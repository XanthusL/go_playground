extern crate rand;
// 使用标准库的io
use std::io;
use rand::Rng;
use std::cmp::Ordering;

fn main() {
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
