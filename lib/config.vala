namespace Puma {
    namespace Config {
        public Toml.Element? read() throws Error {
            return new Toml.Parser.from_path("puma.toml").parse();
        }
    }
}