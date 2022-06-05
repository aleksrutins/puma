name = "puma-test"
authors = ["Aleks Rutins <aleks@rutins.com>"]

library "stuff" {
    src = "src/stuff"
    main = "src/stuff/mod.ts"
}

bundle "puma-test" {
    src = "src/puma-test"
    dependencies = [
        "stuff:stuff"
    ]
}