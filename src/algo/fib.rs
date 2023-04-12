pub fn fib(n: u32) -> u32 {
    match n {
        0 => 0,
        1 => 1,
        _ => fib(n - 1) + fib(n - 2),
    }
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    pub fn fib_of_ten() {
        let val = fib(10);
        assert_eq!(val, 55);
    }
}
