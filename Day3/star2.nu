# Separate first enabled part from the rest
let split = open data.txt | split row -n 2 "don't()"
let first = $split | get 0

# Split by 'don't()' instructions so all strings start disabled
let rest = $split | get 1 | split row "don't()" | each { |line|
    # Check if/when instructions get enabled the frist time
    # by splitting the string and only taking the part after 'do()' instruction if exists
    $line | split row -n 2 "do()" | if ($in | length) == 2 {$in | last} else {null}
}

# Combine both parts and regex parse for 'mul()' instructions
$first | append $rest | str join | parse --regex 'mul\x28(?P<num1>\d{1,3}),(?P<num2>\d{1,3})\x29' | each { |op|
    ($op.num1 | into int) * ($op.num2 | into int)
} | math sum