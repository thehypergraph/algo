export default class Factorial {
  public cache: number[] = [1]

  public find(n: number): number {
    if (n <= 1) {
      return 1
    }

    return n * this.find(n - 1)
  }
}
