fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .out_dir("./vstreamer_protos/src")
        .compile(
            &["../protos/vstreamer_protos/commander/commander.proto"],
            &["../protos"],
        )?;
    Ok(())
}
