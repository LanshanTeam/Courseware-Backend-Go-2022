fn main() {
    let proto_dir = "proto/demo/demo-service.proto";
    tonic_build::compile_protos(proto_dir).unwrap();
    println!("cargo:rerun-if-changed={}", proto_dir);
}
