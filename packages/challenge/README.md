# Challenge description

Design a library that uses multiple packages, that has
the following structure:

```
semver
    pkg A
        import pkg AA1, AA2
        pkg AA1
            declare some vars + init
        pkg AA2
            declare some vars + init
    pkg B
        import pkg BB1
        pkg BB1
            import pkg C
            declare some vars + init
    pkg C
        declare some vars + init
    
main
    import semver library
    main
        use some exported symbols
```

Good luck!
