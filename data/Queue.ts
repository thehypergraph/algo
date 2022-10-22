export default class Queue<T> {
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
    return this.items.pop() ?? null;
  }

  public peek() {
    const slice = this.items.slice(-1);

    if (slice.length === 0) {
      return null;
    }
    return slice[0];
  }

  public size() {
    return this.items.length;
  }

  [Symbol.iterator]() {
    let index = this.items.length - 1;

    return {
      next: () => {
        if (index > -1) {
          const res = { value: this.items[index], done: false };
          index = index - 1;
          return res;
        } else {
          return { done: true };
        }
      },
    };
  }
}
