import { assertEquals } from 'https://deno.land/std@0.160.0/testing/asserts.ts'
import Factorial from './Factorial.ts'

Deno.test('Factorial', () => {
  const fac = new Factorial()
  const res = fac.find(5)
  assertEquals(res, 120)
})
