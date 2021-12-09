export default class Stack {
  items = []
  max = undefined

  constructor(max = undefined) {
    this.max = max
  }

  push(item) {
    if (this?.max !== undefined && this.items.length >= this.max) {
      return false
    }

    this.items.unshift(item)
    return true
  }

  pop() {
    return this.items.shift() ?? null
  }

  peek() {
    const slice = this.items.slice(0, 1)

    if (slice.length === 0) {
      return null
    }
    return slice[0]
  }

  [Symbol.iterator]() {
    let index = 0

    return {
      next: () => {
        if (index < this.items.length) {
          const res = { value: this.items[index], done: false }
          index = index + 1
          return res
        } else {
          return { done: true }
        }
      },
    }
  }
}
