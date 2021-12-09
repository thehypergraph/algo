export default class Factorial {
  cache = [1]

  find(n) {
    if (n <= 1) {
      return 1
    }

    return n * this.find(n - 1)
  }
}
