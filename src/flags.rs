use clap::Parser;
/// Simple program to greet a person
#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
pub struct Flags {
    /// Size of the puzzle
    #[clap(short, long, default_value_t = 3)]
    pub size: usize,

    /// If the puzzle is solvable
    #[clap(long, short)]
    pub unsolvable: bool,

    /// Number of iterations
    #[clap(short, long, default_value_t = 10000)]
    pub iterations: usize,

    /// Cost
    #[clap(short, long, default_value_t = 1)]
    pub cost: u8,

    /// Heuristic
    #[clap(short, long, default_value = "moi")]
    pub heuristic: String,
}

impl Flags {
    pub fn init() -> Flags {
        Flags::parse()
    }
}
