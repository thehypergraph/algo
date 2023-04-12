#[derive(Clone, Debug)]
pub struct Queue<T> {
    items: Vec<T>,
    max: usize,
}

impl Queue<usize> {
    pub fn new(max: usize) -> Queue<usize> {
        Queue {
            items: Vec::new(),
            max,
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

    pub fn items(&self) -> Vec<usize> {
        self.items.clone()
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
        let mut queue = Queue::new(10);
        queue.push(1);
        queue.push(2);
        queue.push(3);
        queue.push(4);

        let mut count = 0;

        for x in queue.iter() {
            count += x;
        }
        assert_eq!(count, 10);
    }
}
