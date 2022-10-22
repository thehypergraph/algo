import { assertEquals } from "https://deno.land/std@0.160.0/testing/asserts.ts";
import Queue from "./Queue.ts";

Deno.test("Queue push with no max returns true", () => {
  const q = new Queue<number>();
  const res = q.push(1);
  assertEquals(res, true);
});

Deno.test("Queue push with no max exceeded returns false", () => {
  const q = new Queue<string>(2);
  q.push("hey");
  q.push("you");
  const res = q.push("there");
  assertEquals(res, false);
});

Deno.test("Queue pop with with no items returns null", () => {
  const q = new Queue();
  const res = q.pop();
  assertEquals(res, null);
});

Deno.test("Queue pop returns first item", () => {
  const q = new Queue<string>();
  q.push("a");
  q.push("b");
  q.push("c");
  const res = q.pop();
  assertEquals(res, "a");
  assertEquals(q.size(), 2);
});

Deno.test("Queue peek returns first item", () => {
  const q = new Queue<string>();
  q.push("a");
  q.push("b");
  q.push("c");
  const res = q.peek();
  assertEquals(res, "a");
  assertEquals(q.size(), 3);
});

// Deno.test("Queue", () => {
//   const q = new Queue<string>();
//   q.push("every");
//   q.push("good");
//   q.push("boy");

//   let res = "";

//   for (const item of q) {
//     res += item;
//   }

//   assertEquals(res, "boygoodevery");
// });
