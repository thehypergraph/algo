const std = @import("std");
const fib = @import("series/fib.zig").fib;

pub fn main() !void {
    const stdout_file = std.io.getStdOut().writer();
    var bw = std.io.bufferedWriter(stdout_file);
    const stdout = bw.writer();

    const res = fib(20);

    try stdout.print("Res {}\n", .{res});

    try bw.flush();
}
