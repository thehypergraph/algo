pub struct Data {
    pub items: Vec<i32>,
}

impl Data {
    pub fn new() -> Data {
        Data { items: vec![] }
    }

    pub fn add(&mut self, item: i32) {
        self.items.push(item);
    }

    pub fn search(&self, item: i32) -> Option<usize> {
        for (i, v) in self.items.iter().enumerate() {
            if *v == item {
                return Some(i);
            }
        }
        None
    }
}
#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    pub fn search_existing() {
        let mut data = Data::new();
        for x in 0..10 {
            data.add(x);
        }
        assert_eq!(data.search(5), Some(5));
    }

    #[test]
    pub fn search_missing() {
        let mut data = Data::new();
        for x in 0..10 {
            data.add(x);
        }
        assert_eq!(data.search(50), None);
    }
}
