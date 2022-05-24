use std::fmt;
#[derive(Debug, Clone)]
pub struct Puzzle {
    pub size: usize,
    pub board: Vec<usize>,
    pub tiles: Vec<usize>,
}

impl fmt::Display for Puzzle {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(
            f,
            "size = {}, board= {:?} , tiles = {:?}",
            self.size, self.board, self.tiles
        )
    }
}

impl Default for Puzzle {
    fn default() -> Self {
        let size: usize = 3;
        Self {
            size: 3,
            board: vec![0; size.pow(2)],
            tiles: vec![0; size.pow(2)],
        }
    }
}

impl Puzzle {
    pub fn new(size: usize, solvable: bool, iterations: usize) -> Self {
        let mut puzzle = Self {
            size,
            board: vec![0; size * size],
            tiles: vec![0; size * size],
        };
        puzzle.init();
        puzzle
    }
    pub fn init(&mut self) {
        for i in 0..self.board.len() {
            self.board[i] = i
        }
    }
}
