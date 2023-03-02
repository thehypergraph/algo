#[derive(Debug)]
pub struct Node<T>  {
		data: Option<T>,
		left: Option<Box<Node<T>>>,
		right: Option<Box<Node<T>>>,
}

impl Node<usize> {
	pub fn new() -> Node<usize> {
		Node { data: None, left: None, right: None }
	}

	// Insert a new node into the tree
	pub fn insert(&mut self, value: usize) {
		match self.data {
			Some(data) => {
				if value < data {
					match self.left {
						Some(ref mut left) => left.insert(value),
						None => self.left = Some(Box::new(Node { data: Some(value), left: None, right: None })),
					}
				} else {
					match self.right {
						Some(ref mut right) => right.insert(value),
						None => self.right = Some(Box::new(Node { data: Some(value), left: None, right: None })),
					}
				}
			},
			None => self.data = Some(value),
		}
	}

	// Print the tree in order
	pub fn print(&self) {
		match self.left {
			Some(ref left) => left.print(),
			None => (),
		}
		match self.data {
			Some(data) => print!("{} ", data),
			None => (),
		}
		match self.right {
			Some(ref right) => right.print(),
			None => (),
		}
	}

	// Find the node with the given value
	pub fn find(&self, value: usize) -> bool {
		match self.data {
			Some(data) => {
				if value == data {
					true
				} else if value < data {
					match self.left {
						Some(ref left) => left.find(value),
						None => false,
					}
				} else {
					match self.right {
						Some(ref right) => right.find(value),
						None => false,
					}
				}
			},
			None => false,
		}
	}

	// Find the minimum value in the tree
	pub fn min(&self) -> Option<usize> {
		match self.left {
			Some(ref left) => left.min(),
			None => self.data,
		}
	}

	// Find the maximum value in the tree
	pub fn max(&self) -> Option<usize> {
		match self.right {
			Some(ref right) => right.max(),
			None => self.data,
		}
	}

	// Remove the node with the given value
	pub fn remove(&mut self, value: usize) -> bool {
		match self.data {
			Some(data) => {
				if value == data {
					match self.left {
						Some(ref mut left) => {
							match left.right {
								Some(ref mut right) => {
									match right.max() {
										Some(max) => {
											self.data = Some(max);
											right.remove(max);
										},
										None => (),
									}
								},
								None => (),
							}
						},
						None => (),
					}
					match self.right {
						Some(ref mut right) => {
							match right.left {
								Some(ref mut left) => {
									match left.min() {
										Some(min) => {
											self.data = Some(min);
											left.remove(min);
										},
										None => (),
									}
								},
								None => (),
							}
						},
						None => (),
					}
					true
				} else if value < data {
					match self.left {
						Some(ref mut left) => left.remove(value),
						None => false,
					}
				} else {
					match self.right {
						Some(ref mut right) => right.remove(value),
						None => false,
					}
				}
			},
			None => false,
		}
	}

	// Return the height of the tree
	pub fn height(&self) -> usize {
		match self.left {
			Some(ref left) => {
				match self.right {
					Some(ref right) => {
						let left_height = left.height();
						let right_height = right.height();
						if left_height > right_height {
							left_height + 1
						} else {
							right_height + 1
						}
					},
					None => left.height() + 1,
				}
			},
			None => {
				match self.right {
					Some(ref right) => right.height() + 1,
					None => 0,
				}
			},
		}
	}

	// Return the number of nodes in the tree
	pub fn size(&self) -> usize {
		match self.left {
			Some(ref left) => {
				match self.right {
					Some(ref right) => left.size() + right.size() + 1,
					None => left.size() + 1,
				}
			},
			None => {
				match self.right {
					Some(ref right) => right.size() + 1,
					None => 1,
				}
			},
		}
	}

	// Return the number of leaves in the tree
	pub fn leaves(&self) -> usize {
		match self.left {
			Some(ref left) => {
				match self.right {
					Some(ref right) => left.leaves() + right.leaves(),
					None => left.leaves(),
				}
			},
			None => {
				match self.right {
					Some(ref right) => right.leaves(),
					None => 1,
				}
			},
		}
	}

	// Return the number of full nodes in the tree
	pub fn full_nodes(&self) -> usize {
		match self.left {
			Some(ref left) => {
				match self.right {
					Some(ref right) => left.full_nodes() + right.full_nodes() + 1,
					None => left.full_nodes(),
				}
			},
			None => {
				match self.right {
					Some(ref right) => right.full_nodes(),
					None => 0,
				}
			},
		}
	}
}
