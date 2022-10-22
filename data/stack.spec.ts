import { assertEquals } from "https://deno.land/std@0.160.0/testing/asserts.ts";
import Stack from "./Stack.ts";

Deno.test("Stack push returns first item", () => {
  const s = new Stack();
  const res = s.push(1);
  assertEquals(res, true);
});

Deno.test("Stack push returns false when over max", () => {
  const s = new Stack(2);
  s.push("hey");
  s.push("you");
  const res = s.push("there");
  assertEquals(res, false);
});

Deno.test("Stack pop with with no items returns null", () => {
  const s = new Stack();
  const res = s.pop();
  assertEquals(res, null);
});

Deno.test("Stack pop returns last item entered", () => {
  const s = new Stack();
  s.push("a");
  s.push("b");
  s.push("c");
  const res = s.pop();
  assertEquals(res, "c");
  assertEquals(s.size(), 2);
});

Deno.test("Stack peek returns first item", () => {
  const s = new Stack();
  s.push("a");
  s.push("b");
  s.push("c");
  const res = s.pop();
  assertEquals(res, "c");
  assertEquals(s.size(), 2);
});

Deno.test("Stack peek returns first item", () => {
  const s = new Stack();
  s.push("a");
  s.push("b");
  s.push("c");
  const res = s.pop();
  assertEquals(res, "c");
  assertEquals(s.size(), 2);
});
