fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .out_dir("./dist/src")
        .compile(
            &["../protos/vstreamer_protos/commander/commander.proto"],
            &["../protos"],
        )?;
    Ok(())
}
