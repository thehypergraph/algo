export default class Stack<T> {
  private items: T[] = [];
  private max?: number;

  constructor(max?: number) {
    this.max = max;
  }

  public push(item: T) {
    if (this?.max !== undefined && this.items.length >= this.max) {
      return false;
    }

    this.items.unshift(item);
    return true;
  }

  public pop() {
    return this.items.shift() ?? null;
  }

  public peek() {
    const slice = this.items.slice(0, 1);

    if (slice.length === 0) {
      return null;
    }
    return slice[0];
  }

  public size() {
    return this.items.length;
  }

  [Symbol.iterator]() {
    let index = 0;

    return {
      next: () => {
        if (index < this.items.length) {
          const res = { value: this.items[index], done: false };
          index = index + 1;
          return res;
        } else {
          return { done: true };
        }
      },
    };
  }
}
