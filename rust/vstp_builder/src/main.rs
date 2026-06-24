fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_prost_build::configure()
        .out_dir("./vstreamer_protos/src")
        .compile_protos(
            &["../protos/vstreamer_protos/commander/commander.proto"],
            &["../protos"],
        )?;
    Ok(())
}
