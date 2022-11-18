interface Node<T> {
  data: T
  left: Node<T> | null
  right: Node<T> | null
}

interface MultiTree<T> {
  root: Node<T> | null
}

export default class Tree<T> {
  private tree: MultiTree<T>

  constructor() {
    this.tree = {
      root: null,
    }
  }

  public create(data: T) {
    this.tree.root = {
      data,
      left: null,
      right: null,
    }
  }

  public add(data: T) {}
}
