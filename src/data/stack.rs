pub struct Stack<T> {
    items: Vec<T>,
    max: usize,
}

impl Stack<usize> {
    pub fn new(max: usize) -> Stack<usize> {
        Stack {
            items: Vec::new(),
            max: max,
        }
    }

    pub fn push(&mut self, item: usize) -> bool {
        if self.max != 0 && self.items.len() >= self.max {
            return false;
        }
        self.items.insert(0, item);
        true
    }

    pub fn pop(&mut self) -> Option<usize> {
        self.items.pop()
    }

    pub fn peek(&self) -> Option<&usize> {
        self.items.last()
    }

    pub fn size(&self) -> usize {
        self.items.len()
    }

    pub fn iter(&self) -> std::slice::Iter<usize> {
        self.items.iter()
    }
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    pub fn iterate() {
        let mut stack = Stack::new(10);
        stack.push(1);
        stack.push(2);
        stack.push(3);
        stack.push(4);

        let mut count = 0;

        for x in stack.iter() {
            count += x;
        }
        assert_eq!(count, 10);
    }
}
