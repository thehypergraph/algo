export default class Queue {
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
    return this.items.pop() ?? null
  }

  peek() {
    const slice = this.items.slice(-1)

    if (slice.length === 0) {
      return null
    }
    return slice[0]
  }

  [Symbol.iterator]() {
    let index = this.items.length - 1

    return {
      next: () => {
        if (index > -1) {
          const res = { value: this.items[index], done: false }
          index = index - 1
          return res
        } else {
          return { done: true }
        }
      },
    }
  }
}
