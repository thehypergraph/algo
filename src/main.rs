pub mod data;
use data::queue::Queue;

fn main() {
    let mut queue = Queue::new(10);
    queue.push(1);
    queue.push(2);
    queue.push(3);
    queue.push(4);

    println!("Hello, world!");

    while queue.size() > 0 {
        println!("{}", queue.pop().unwrap());
    }

}
