import test from 'japa'

import Queue from '../Queue.mjs'

test.group('Queue', () => {
  test('push with no max returns true', (assert) => {
    const q = new Queue()
    const res = q.push(1)
    assert.isTrue(res)
  })

  test('push with no max exceeded returns false', (assert) => {
    const q = new Queue(2)
    q.push('hey')
    q.push('you')
    const res = q.push('there')
    assert.isFalse(res)
  })

  test('pop with with no items returns null', (assert) => {
    const q = new Queue()
    const res = q.pop()
    assert.isNull(res)
  })

  test('pop returns first item', (assert) => {
    const q = new Queue()
    q.push('a')
    q.push('b')
    q.push('c')
    const res = q.pop()
    assert.equal(res, 'a')
    assert.lengthOf(q.items, 2)
  })

  test('peek returns first item', (assert) => {
    const q = new Queue()
    q.push('a')
    q.push('b')
    q.push('c')
    const res = q.peek()
    assert.equal(res, 'a')
    assert.lengthOf(q.items, 3)
  })

  test('correctly iterates over queue items', (assert) => {
    const q = new Queue()
    q.push('every')
    q.push('good')
    q.push('boy')

    let res = ''

    for (let item of q) {
      res += item
    }

    assert.equal(res, 'everygoodboy')
  })
})
