# BAD BAD VERY NOT GOOD SOLUTION
open data.txt | lines | reduce --fold 0 { |line, acc|
    let num_arr = $line | split words | each {into int}

    mut safe = 1
    mut prev_val = $num_arr | first
    mut prev_sign: string = ""

    for $num in ($num_arr | skip 1) {
        let sign = $"(($prev_val - $num) > 0)"
        if ($prev_sign | is-empty) {
            $prev_sign = $sign
        } else if (not ($prev_sign | str contains $sign)) {
            $safe = 0
            break
        }
        if (($prev_val - $num | math abs) in 1..3) {
            $prev_val = $num
        } else {
            $safe = 0
            break
        }
    }
    if ($safe == 0) {
        mut looping = true
        # I hate this tbh
        for $index in 0..(($num_arr | length) - 1) {
            let inner_arr = $num_arr | drop nth $index

            $safe = 1
            mut prev_val = $inner_arr | first
            mut prev_sign: string = ""

            for $num in ($inner_arr | skip 1) {
                let sign = $"(($prev_val - $num) > 0)"
                if ($prev_sign | is-empty) {
                    $prev_sign = $sign
                } else if (not ($prev_sign | str contains $sign)) {
                    $safe = 0
                    break
                }
                if (($prev_val - $num | math abs) in 1..3) {
                    $prev_val = $num
                } else {
                    $safe = 0
                    break
                }
            }
            if ($safe == 1) {break}
        }
    }
    $safe + $acc
}