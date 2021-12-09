import test from 'japa'
import Factorial from '../Factorial.mjs'

test.group('Factorial', () => {
  test('find fac 5', (assert) => {
    const fac = new Factorial()
    const res = fac.find(5)
    assert.equal(res, 120)
  })
})
