void main() {
    new ValaConsole.Console("puma").log(Puma.Config.read()["hello"].as<string>());
}