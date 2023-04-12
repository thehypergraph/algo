pub fn factorial(n: u32) -> u32 {
    if n <= 1 {
        return 1;
    }
    n * factorial(n - 1)
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    pub fn factorial_of_ten() {
        let val = factorial(10);
				assert_eq!(val, 3628800);
    }
}
