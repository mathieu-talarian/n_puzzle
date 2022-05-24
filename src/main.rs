use puzzle::Puzzle;

mod flags;
mod logger;
mod puzzle;

#[macro_use]
extern crate tracing;

fn main() {
    // Log something simple. In `tracing` parlance, this creates an "event".
    logger::init_logger();
    let flags = flags::Flags::init();

    let puzzle = Puzzle::new(flags.size, !flags.unsolvable, flags.iterations);

    info!(?puzzle)
}
