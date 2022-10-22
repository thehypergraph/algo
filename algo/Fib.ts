export default class Fib {
  public cache: number[] = [1];

  public fib(n: number): number {
    if (this.cache[n]) {
      return this.cache[n];
    }
    this.cache[n] = this.fib(n - 1) + this.fib(n - 2);
    return this.cache[n];
  }
}
