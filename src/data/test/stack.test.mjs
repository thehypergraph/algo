import test from 'japa'

import Stack from '../Stack.mjs'

test.group('Stack', () => {
  test('push with no max returns true', (assert) => {
    const s = new Stack()
    const res = s.push(1)
    assert.isTrue(res)
  })

  test('push with no max exceeded returns false', (assert) => {
    const s = new Stack(2)
    s.push('hey')
    s.push('you')
    const res = s.push('there')
    assert.isFalse(res)
  })

  test('pop with with no items returns null', (assert) => {
    const s = new Stack()
    const res = s.pop()
    assert.isNull(res)
  })

  test('pop returns last item entered', (assert) => {
    const s = new Stack()
    s.push('a')
    s.push('b')
    s.push('c')
    const res = s.pop()
    assert.equal(res, 'c')
    assert.lengthOf(s.items, 2)
  })

  test('peek returns first item', (assert) => {
    const s = new Stack()
    s.push('a')
    s.push('b')
    s.push('c')
    const res = s.peek()
    assert.equal(res, 'c')
    assert.lengthOf(s.items, 3)
  })

  test('correctly iterates over stack items', (assert) => {
    const stack = new Stack()
    stack.push('every')
    stack.push('good')
    stack.push('boy')

    let res = ''

    for (let item of stack) {
      res += item
    }

    assert.equal(res, 'boygoodevery')
  })
})
