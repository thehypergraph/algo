import test from 'japa'
import { fib } from '../Fib.mjs'

test.group('Fib', () => {
  test('fib 8 is 21', (assert) => {
    const res = fib(7)
    assert.equal(res, 21)
  })
})
