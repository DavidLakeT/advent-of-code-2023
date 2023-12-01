use regex::Regex;

pub fn solve_part_1() -> u32 {
    let input = include_str!(r"inputs/day1.txt").lines();

    let mut sum = 0;

    for line in input {
        let mut digits = line.chars().filter(|c| c.is_ascii_digit());

        let first_digit = digits.next().unwrap().to_digit(10).unwrap();
        let last_digit = digits.last().unwrap().to_digit(10).unwrap();

        let number = first_digit * 10 + last_digit;

        sum += number;
    }

    sum
}

pub fn solve_part_2() -> u32 {
    let input = include_str!(r"inputs/day1.txt").lines();
    let regex = Regex::new(r"(?:one|two|three|four|five|six|seven|eight|nine|\d)").unwrap();

    let mut sum = 0;

    for line in input {
        let mut digits = regex.find_iter(line).map(|m| get_digit_value(m.as_str()));

        let first_digit = digits.next().unwrap();

        let last_digit = digits.last().unwrap_or(first_digit);

        let number = first_digit * 10 + last_digit;

        sum += number;
    }

    sum
}

fn get_digit_value(str: &str) -> u32 {
    let values = [
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];

    if let Some(index) = values.iter().position(|&x| x == str) {
        return (index + 1) as u32;
    }

    str.parse::<u32>().unwrap_or(0)
}
