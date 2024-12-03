open data.txt | lines | each { |line|
    # Regex baybeeeeee (I hate it here)
    $line | parse --regex 'mul\x28(?P<num1>\d{1,3}),(?P<num2>\d{1,3})\x29' | each { |op|
        ($op.num1 | into int) * ($op.num2 | into int)
    } | math sum
} | math sum