rule rule1 {
    meta:
        description = "just a dummy rule"
        author = "Luca Corbatto"
        date = "2020-01-01"

    strings:
        $str1 = "hello world"

    condition:
        $str1
}