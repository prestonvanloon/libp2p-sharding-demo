http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.10.1/rules_go-0.10.1.tar.gz",
    sha256 = "4b14d8dd31c6dbaf3ff871adcd03f28c3274e42abc855cb8fb4d01233c0154dc",
)

http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.10.0/bazel-gazelle-0.10.0.tar.gz",
    sha256 = "6228d9618ab9536892aa69082c063207c91e777e51bd3c5544c9c060cafe1bd8",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

go_repository(
    name = "com_github_multiformats_go_multiaddr",
    importpath = "github.com/multiformats/go-multiaddr",
    commit = "123a717755e0559ec8fda308019cd24e0a37bb07",  
)

go_repository(
    name = "com_github_libp2p_go_libp2p",
    importpath = "github.com/libp2p/go-libp2p",
    commit = "6877a20c8a91eb0d53244d4c358ed8b6aa085eba", 
)

go_repository(
    name = "com_github_libp2p_go_libp2p_swarm",
    importpath = "github.com/libp2p/go-libp2p-swarm",
    commit = "c0bf6756a81190a94b5946f952cd4fc7d342b0d3",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peerstore",
    importpath = "github.com/libp2p/go-libp2p-peerstore",
    commit = "2c122745dcc9c03c80b0203ce89c445ace6aa596",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peer",
    importpath = "github.com/libp2p/go-libp2p-peer",
    commit = "296ae64e3db4af20e64270faa18e28797643480e",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_net",
    importpath = "github.com/libp2p/go-libp2p-net",
    commit = "64c31f18a9753193909544f1a606148894cbe878",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_host",
    importpath = "github.com/libp2p/go-libp2p-host",
    commit = "dfeeb00fc9decef43a6bc47b555a530cb38e0c62",  
)

go_repository(
    name = "com_github_libp2p_go_libp2p_crypto",
    importpath = "github.com/libp2p/go-libp2p-crypto",
    commit = "f2afd8b3d4e358a6092159c813cb2bdf6d984e68",  
)


go_repository(
    name = "com_github_multiformats_go_multistream",
    importpath = "github.com/multiformats/go-multistream",
    commit = "612ce31c03aebe1d5adbd3c850ee89e05a82b16d",  
)


go_repository(
    name = "com_github_whyrusleeping_go_smux_yamux",
    importpath = "github.com/whyrusleeping/go-smux-yamux",
    commit = "4b90b786fa86d970c5db5a3b61088de8c8c0c14c",  
)


go_repository(
    name = "com_github_whyrusleeping_multiaddr_filter",
    importpath = "github.com/whyrusleeping/multiaddr-filter",
    commit = "e903e4adabd70b78bc9293b6ee4f359afb3f9f59",  
)


go_repository(
    name = "com_github_whyrusleeping_go_smux_spdystream",
    importpath = "github.com/whyrusleeping/go-smux-spdystream",
    commit = "a6182ff2a058b177f3dc7513fe198e6002f7be78",  
)


go_repository(
    name = "com_github_ipfs_go_log",
    importpath = "github.com/ipfs/go-log",
    commit = "0e2a17b81af445147a78d94a45d4398c4008c219",  
)


go_repository(
    name = "com_github_libp2p_go_libp2p_protocol",
    importpath = "github.com/libp2p/go-libp2p-protocol",
    commit = "b29f3d97e3a2fb8b29c5d04290e6cb5c5018004b",  
)


go_repository(
    name = "com_github_coreos_go_semver",
    importpath = "github.com/coreos/go-semver",
    commit = "e214231b295a8ea9479f11b70b35d5acf3556d9b",  
)


go_repository(
    name = "com_github_libp2p_go_libp2p_interface_connmgr",
    importpath = "github.com/libp2p/go-libp2p-interface-connmgr",
    commit = "201197111979ad70393a51d80cb12c5dfcee9d01",  
)


go_repository(
    name = "com_github_libp2p_go_ws_transport",
    importpath = "github.com/libp2p/go-ws-transport",
    commit = "0cf90461ccc19765a5afafa9c2f58291471f1f09",  
)


go_repository(
    name = "com_github_libp2p_go_addr_util",
    importpath = "github.com/libp2p/go-addr-util",
    commit = "a9d6b939b59796c4933f56a8e628d6682d8d2a05",
)


go_repository(
    name = "com_github_whyrusleeping_go_smux_multistream",
    importpath = "github.com/whyrusleeping/go-smux-multistream",
    commit = "afa6825376c14a0462fd420a7d4b4d157c937a42",  
)


go_repository(
    name = "com_github_libp2p_go_tcp_transport",
    importpath = "github.com/libp2p/go-tcp-transport",
    commit = "578cd65fcec1d08d226d1b186e4548065e6bd65c",  
)


go_repository(
    name = "com_github_libp2p_go_stream_muxer",
    importpath = "github.com/libp2p/go-stream-muxer",
    commit = "6ebe3f58af097068454b167a89442050b023b571",
)


go_repository(
    name = "com_github_multiformats_go_multiaddr_dns",
    importpath = "github.com/multiformats/go-multiaddr-dns",
    commit = "b431cd0a670beb991d0d0dc3defd2b89125aa34e",  
)


go_repository(
    name = "com_github_libp2p_go_peerstream",
    importpath = "github.com/libp2p/go-peerstream",
    commit = "de27dedf2287a2bd7030a1ecc94a6913385e6cfe",  
)


go_repository(
    name = "com_github_jbenet_goprocess",
    importpath = "github.com/jbenet/goprocess",
    commit = "b497e2f366b8624394fb2e89c10ab607bebdde0b",  
)


go_repository(
    name = "com_github_libp2p_go_maddr_filter",
    importpath = "github.com/libp2p/go-maddr-filter",
    commit = "3c9947befbb92277cc5f85057d387097debc4139",
)


go_repository(
    name = "com_github_libp2p_go_libp2p_nat",
    importpath = "github.com/libp2p/go-libp2p-nat",
    commit = "469319376dce9de8de967a034bf579b4ae7be44d",  
)


go_repository(
    name = "com_github_libp2p_go_libp2p_transport",
    importpath = "github.com/libp2p/go-libp2p-transport",
    commit = "7f67e6aa68aa4b679c21179d32ab08d77a149cef",  
)


go_repository(
    name = "com_github_libp2p_go_libp2p_metrics",
    importpath = "github.com/libp2p/go-libp2p-metrics",
    commit = "a3a76fa99809135a2421982ff2958a544aa517ee",  
)


go_repository(
    name = "com_github_libp2p_go_libp2p_conn",
    importpath = "github.com/libp2p/go-libp2p-conn",
    commit = "8058ec7c7d8d3834119a164005482ea696f028e4",  
)


go_repository(
    name = "com_github_libp2p_go_libp2p_loggables",
    importpath = "github.com/libp2p/go-libp2p-loggables",
    commit = "7a534896f9574b69081fa4e3a192270465719ced",  
)


go_repository(
    name = "com_github_libp2p_go_libp2p_interface_pnet",
    importpath = "github.com/libp2p/go-libp2p-interface-pnet",
    commit = "6ea0626616ac1d34c6375a20fd1f48fc9a607a4d",  
)


go_repository(
    name = "com_github_libp2p_go_libp2p_interface_conn",
    importpath = "github.com/libp2p/go-libp2p-interface-conn",
    commit = "0635e022759c738c55b7769cc6307506946089b9",  
)


go_repository(
    name = "com_github_libp2p_go_libp2p_circuit",
    importpath = "github.com/libp2p/go-libp2p-circuit",
    commit = "77e780c2043662c798da30750e5af7e07b086703",  
)


go_repository(
    name = "com_github_btcsuite_btcd",
    importpath = "github.com/btcsuite/btcd",
    commit = "2be2f12b358dc57d70b8f501b00be450192efbc3",  
)


go_repository(
    name = "com_github_agl_ed25519",
    importpath = "github.com/agl/ed25519",
    commit = "5312a61534124124185d41f09206b9fef1d88403",  
)


go_repository(
    name = "com_github_multiformats_go_multiaddr",
    importpath = "github.com/multiformats/go-multiaddr",
    commit = "123a717755e0559ec8fda308019cd24e0a37bb07",  
)


go_repository(
    name = "com_github_multiformats_go_multihash",
    importpath = "github.com/multiformats/go-multihash",
    commit = "265e72146e710ff649c6982e3699d01d4e9a18bb",  
)


go_repository(
    name = "com_github_multiformats_go_multicodec_packed",
    importpath = "github.com/multiformats/go-multicodec-packed",
    commit = "9004b413b478e5a878e4a879358cce02e5df4995",  
)


go_repository(
    name = "com_github_mr_tron_base58",
    importpath = "github.com/mr-tron/base58",
    commit = "c1bdf7c52f59d6685ca597b9955a443ff95eeee6",  
)


go_repository(
    name = "com_github_opentracing_opentracing_go",
    importpath = "github.com/opentracing/opentracing-go",
    commit = "3999fca714c888484b78cb78d68cd49afc9eca7d",  
)


go_repository(
    name = "com_github_multiformats_go_multiaddr_net",
    importpath = "github.com/multiformats/go-multiaddr-net",
    commit = "97d80565f68c5df715e6ba59c2f6a03d1fc33aaf",  
)


go_repository(
    name = "com_github_whyrusleeping_go_logging",
    importpath = "github.com/whyrusleeping/go-logging",
    commit = "0457bb6b88fc1973573aaf6b5145d8d3ae972390",  
)


go_repository(
    name = "com_github_whyrusleeping_mafmt",
    importpath = "github.com/whyrusleeping/mafmt",
    commit = "ab6a47300c1df531e468771e7d08fcd6d33f032e",  
)


go_repository(
    name = "com_github_libp2p_go_flow_metrics",
    importpath = "github.com/libp2p/go-flow-metrics",
    commit = "3b3bcfcf78f2dc0e85be13ef3c3adc64cc5a9347",  
)


go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    vcs = "git",
    remote = "https://github.com/golang/crypto.git",
    commit = "21652f85b0fdddb6c2b6b77a5beca5c5a908174a",  
)


go_repository(
    name = "com_github_spaolacci_murmur3",
    importpath = "github.com/spaolacci/murmur3",
    commit = "f09979ecbc725b9e6d41a297405f65e7e8804acc",  
)


go_repository(
    name = "com_github_minio_sha256_simd",
    importpath = "github.com/minio/sha256-simd",
    commit = "ad98a36ba0da87206e3378c556abbfeaeaa98668",  
)


go_repository(
    name = "com_github_minio_blake2b_simd",
    importpath = "github.com/minio/blake2b-simd",
    commit = "3f5f724cb5b182a5c278d6d3d55b40e7f8c2efb4",  
)


go_repository(
    name = "com_github_gxed_hashland",
    importpath = "github.com/gxed/hashland",
    commit = "d9f6b97f8db22dd1e090fd0bbbe98f09cc7dd0a8",  
)


go_repository(
    name = "com_github_ipfs_go_ipfs_util",
    importpath = "github.com/ipfs/go-ipfs-util",
    commit = "9ed527918c2f20abdf0adfab0553cd87db34f656",  
)


go_repository(
    name = "com_github_whyrusleeping_go_notifier",
    importpath = "github.com/whyrusleeping/go-notifier",
    commit = "097c5d47330ff6a823f67e3515faa13566a62c6f",  
)


go_repository(
    name = "com_github_fd_go_nat",
    importpath = "github.com/fd/go-nat",
    commit = "dcaf50131e4810440bed2cbb6f7f32c4f4cc95dd",  
)


go_repository(
    name = "com_github_libp2p_go_reuseport",
    importpath = "github.com/libp2p/go-reuseport",
    commit = "99b8511da09c3f00966861e1ed210e808c568112",  
)


go_repository(
    name = "com_github_whyrusleeping_yamux",
    importpath = "github.com/whyrusleeping/yamux",
    commit = "7e8368b995874cb6ce27e42d3d9115221f4c879e",  
)


go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    commit = "eb925808374e5ca90c83401a40d711dc08c0c0f6",  
)


go_repository(
    name = "com_github_satori_go_uuid",
    importpath = "github.com/satori/go.uuid",
    commit = "36e9d2ebbde5e3f13ab2e25625fd453271d6522e",  
)


go_repository(
    name = "com_github_docker_spdystream",
    importpath = "github.com/docker/spdystream",
    commit = "bc6354cbbc295e925e4c611ffe90c1f287ee54db",  
)


go_repository(
    name = "com_github_jbenet_go_temp_err_catcher",
    importpath = "github.com/jbenet/go-temp-err-catcher",
    commit = "aac704a3f4f27190b4ccc05f303a4931fd1241ff",  
)


go_repository(
    name = "com_github_libp2p_go_libp2p_secio",
    importpath = "github.com/libp2p/go-libp2p-secio",
    commit = "437de218ca1fc65dfbf7b2b084519665bbaf7781",  
)


go_repository(
    name = "com_github_libp2p_go_msgio",
    importpath = "github.com/libp2p/go-msgio",
    commit = "d82125c9907e1365775356505f14277d47dfd4d6",  
)


go_repository(
    name = "com_github_jackpal_go_nat_pmp",
    importpath = "github.com/jackpal/go-nat-pmp",
    commit = "28a68d0c24adce1da43f8df6a57340909ecd7fdd",
)


go_repository(
    name = "com_github_jackpal_gateway",
    importpath = "github.com/jackpal/gateway",
    commit = "5795ac81146e01d3fab7bcf21c043c3d6a32b006",  
)


go_repository(
    name = "com_github_huin_goupnp",
    importpath = "github.com/huin/goupnp",
    commit = "e25a5cc217cad359cb2c52d1529d5ad7489752c2",  
)


go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    vcs = "git",
    remote = "https://github.com/golang/sys.git",
    commit = "89ac7f292d17a339edab4de8efcba5d8672ff661",
)


go_repository(
    name = "com_github_libp2p_go_sockaddr",
    importpath = "github.com/libp2p/go-sockaddr",
    commit = "09ae606455f81dc795c406a93b35993fcfc3a912",  
)


go_repository(
    name = "com_github_gxed_eventfd",
    importpath = "github.com/gxed/eventfd",
    commit = "80a92cca79a8041496ccc9dd773fcb52a57ec6f9",  
)


go_repository(
    name = "com_github_gxed_goendian",
    importpath = "github.com/gxed/GoEndian",
    commit = "0f5c6873267e5abf306ffcdfcfa4bf77517ef4a7",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    commit = "8991bc29aa16c548c550c7ff78260e27b9ab7c73",
    importpath = "github.com/davecgh/go-spew",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_multiplex",
    commit = "feac16f6d6b1dafe1963afcca468baf6e430e01d",
    importpath = "github.com/whyrusleeping/go-smux-multiplex",
)

go_repository(
    name = "com_github_whyrusleeping_go_multiplex",
    commit = "3ac9031715040479fc59cbf43a09f6a08b3ef9ae",
    importpath = "github.com/whyrusleeping/go-multiplex",
)
