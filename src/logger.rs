use tracing_subscriber::fmt;

pub fn init_logger() {
    let format = fmt::format().pretty().compact();

    tracing_subscriber::fmt().event_format(format).init();
}
